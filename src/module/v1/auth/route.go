package auth

import (
	"modular-monolithic/module/v1/auth/handler"
	"net/http"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	AuthHandler := handler.AuthHandler{
		Carrier:     c.Carrier,
		AuthService: c.AuthService,
		UserService: c.UserService,
	}

	//User Register
	authRoutes := c.R.PathPrefix("/auth").Subrouter()

	authRoutes.HandleFunc("/login", AuthHandler.Login).Methods(http.MethodPost).Name("auth.login")
}
