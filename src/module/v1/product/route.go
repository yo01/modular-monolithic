package product

import (
	"net/http"

	"modular-monolithic/module/v1/product/handler"
	"modular-monolithic/security/middleware"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	ProductHandler := handler.ProductHandler{
		Carrier:        c.Carrier,
		ProductService: c.ProductService,
	}

	// PRODUCT ROUTES WITH MIDDLEWARE
	productRoutesWithMiddleware := c.R.PathPrefix("/products").Subrouter()
	productRoutesWithMiddleware.Use(middleware.JWT)

	productRoutesWithMiddleware.HandleFunc("", ProductHandler.Create).Methods(http.MethodPost).Name("product.save")
	productRoutesWithMiddleware.HandleFunc("/{id}", ProductHandler.Edit).Methods(http.MethodPut).Name("product.edit")
	productRoutesWithMiddleware.HandleFunc("/{id}", ProductHandler.Delete).Methods(http.MethodDelete).Name("product.delete")

	// PRODUCT ROUTES WITHOUT MIDDLEWARE
	productRoutesWithoutMiddleware := c.R.PathPrefix("/products").Subrouter()

	productRoutesWithoutMiddleware.HandleFunc("", ProductHandler.List).Methods(http.MethodGet).Name("product.list")
	productRoutesWithoutMiddleware.HandleFunc("/{id}", ProductHandler.Detail).Methods(http.MethodGet).Name("product.detail")
}
