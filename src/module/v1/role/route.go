package role

import (
	"modular-monolithic/module/v1/role/handler"
	"modular-monolithic/security/middleware"

	"net/http"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	RoleHandler := handler.RoleHandler{
		Carrier:     c.Carrier,
		RoleService: c.RoleService,
	}

	// ROLE ROUTES WITH MIDDLEWARE
	roleRoutesWithMiddleware := c.R.PathPrefix("/roles").Subrouter()
	roleRoutesWithMiddleware.Use(middleware.JWT)

	roleRoutesWithMiddleware.HandleFunc("/{id}", RoleHandler.Edit).Methods(http.MethodPut).Name("role.edit")
	roleRoutesWithMiddleware.HandleFunc("/{id}", RoleHandler.Delete).Methods(http.MethodDelete).Name("role.delete")

	// ROLE ROUTES WITHOUT MIDDLEWARE
	roleRoutesWithoutMiddleware := c.R.PathPrefix("/roles").Subrouter()

	roleRoutesWithoutMiddleware.HandleFunc("/", RoleHandler.Create).Methods(http.MethodPost).Name("role.save")
	roleRoutesWithoutMiddleware.HandleFunc("/", RoleHandler.List).Methods(http.MethodGet).Name("role.list")
	roleRoutesWithoutMiddleware.HandleFunc("/{id}", RoleHandler.Detail).Methods(http.MethodGet).Name("role.detail")
}
