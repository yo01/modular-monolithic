package product

import (
	"modular-monolithic/module/v1/product/handler"
	"net/http"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	ProductHandler := handler.ProductHandler{
		Carrier:        c.Carrier,
		ProductService: c.ProductService,
	}

	//User Register
	productRoutes := c.R.PathPrefix("/products").Subrouter()

	productRoutes.HandleFunc("/", ProductHandler.List).Methods(http.MethodGet).Name("product.list")
	productRoutes.HandleFunc("/{id}", ProductHandler.Detail).Methods(http.MethodGet).Name("product.detail")
	productRoutes.HandleFunc("/", ProductHandler.Create).Methods(http.MethodPost).Name("product.save")
	productRoutes.HandleFunc("/{id}", ProductHandler.Edit).Methods(http.MethodPut).Name("product.edit")
	productRoutes.HandleFunc("/{id}", ProductHandler.Delete).Methods(http.MethodDelete).Name("product.delete")
}
