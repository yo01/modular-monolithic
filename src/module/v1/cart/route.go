package cart

import (
	"net/http"

	"modular-monolithic/module/v1/cart/handler"
	"modular-monolithic/security/middleware"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	CartHandler := handler.CartHandler{
		Carrier:         c.Carrier,
		CartService:     c.CartService,
		CartItemService: c.CartItemService,
	}

	// CART ROUTES WITH MIDDLEWARE
	cartRoutesWithMiddleware := c.R.PathPrefix("/carts").Subrouter()
	cartRoutesWithMiddleware.Use(middleware.JWT)

	cartRoutesWithMiddleware.HandleFunc("", CartHandler.List).Methods(http.MethodGet).Name("cart.list")
	cartRoutesWithMiddleware.HandleFunc("/{id}", CartHandler.Detail).Methods(http.MethodGet).Name("cart.detail")
	cartRoutesWithMiddleware.HandleFunc("", CartHandler.Create).Methods(http.MethodPost).Name("cart.save")
	cartRoutesWithMiddleware.HandleFunc("/{id}", CartHandler.Edit).Methods(http.MethodPut).Name("cart.edit")
	cartRoutesWithMiddleware.HandleFunc("/{id}", CartHandler.Delete).Methods(http.MethodDelete).Name("cart.delete")

	// CART ROUTES WITHOUT MIDDLEWARE
}
