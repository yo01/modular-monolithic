package cart

import (
	"net/http"

	"modular-monolithic/module/v1/cart/handler"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	CartHandler := handler.CartHandler{
		Carrier:     c.Carrier,
		CartService: c.CartService,
	}

	//User Register
	cartRoutes := c.R.PathPrefix("/carts").Subrouter()

	cartRoutes.HandleFunc("/", CartHandler.List).Methods(http.MethodGet).Name("cart.list")
	cartRoutes.HandleFunc("/{id}", CartHandler.Detail).Methods(http.MethodGet).Name("cart.detail")
	cartRoutes.HandleFunc("/", CartHandler.Create).Methods(http.MethodPost).Name("cart.save")
	cartRoutes.HandleFunc("/{id}", CartHandler.Edit).Methods(http.MethodPut).Name("cart.edit")
	cartRoutes.HandleFunc("/{id}", CartHandler.Delete).Methods(http.MethodDelete).Name("cart.delete")
}
