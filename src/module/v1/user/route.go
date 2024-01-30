package user

import (
	"modular-monolithic/module/v1/user/handler"
	"net/http"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	UserHandler := handler.UserHandler{
		Carrier:     c.Carrier,
		UserService: c.UserService,
	}

	//User Register
	userRoutes := c.R.PathPrefix("/users").Subrouter()

	userRoutes.HandleFunc("/", UserHandler.List).Methods(http.MethodGet).Name("user.list")
	userRoutes.HandleFunc("/{id}", UserHandler.Detail).Methods(http.MethodGet).Name("user.detail")
	userRoutes.HandleFunc("/", UserHandler.Create).Methods(http.MethodPost).Name("user.save")
	userRoutes.HandleFunc("/{id}", UserHandler.Edit).Methods(http.MethodPut).Name("user.edit")
	userRoutes.HandleFunc("/{id}", UserHandler.Delete).Methods(http.MethodDelete).Name("user.delete")
}
