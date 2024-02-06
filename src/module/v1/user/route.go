package user

import (
	"modular-monolithic/module/v1/user/handler"
	"modular-monolithic/security/middleware"

	"net/http"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	UserHandler := handler.UserHandler{
		Carrier:     c.Carrier,
		UserService: c.UserService,
	}

	// USER ROUTES WITH MIDDLEWARE
	userRoutesWithMiddleware := c.R.PathPrefix("/users").Subrouter()
	userRoutesWithMiddleware.Use(middleware.JWT)

	userRoutesWithMiddleware.HandleFunc("/{id}", UserHandler.Edit).Methods(http.MethodPut).Name("user.edit")
	userRoutesWithMiddleware.HandleFunc("/{id}", UserHandler.Delete).Methods(http.MethodDelete).Name("user.delete")

	// USER ROUTES WITHOUT MIDDLEWARE
	userRoutesWithoutMiddleware := c.R.PathPrefix("/users").Subrouter()

	userRoutesWithMiddleware.HandleFunc("", UserHandler.List).Methods(http.MethodGet).Name("user.list")
	userRoutesWithoutMiddleware.HandleFunc("/{id}", UserHandler.Detail).Methods(http.MethodGet).Name("user.detail")
	userRoutesWithoutMiddleware.HandleFunc("", UserHandler.Create).Methods(http.MethodPost).Name("user.save")
}
