package role

import (
	"modular-monolithic/module/v1/role/handler"

	"net/http"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	RoleHandler := handler.RoleHandler{
		Carrier:     c.Carrier,
		RoleService: c.RoleService,
	}

	//User Register
	roleRoutes := c.R.PathPrefix("/roles").Subrouter()

	roleRoutes.HandleFunc("/", RoleHandler.List).Methods(http.MethodGet).Name("role.list")
	roleRoutes.HandleFunc("/{id}", RoleHandler.Detail).Methods(http.MethodGet).Name("role.detail")
	roleRoutes.HandleFunc("/", RoleHandler.Create).Methods(http.MethodPost).Name("role.save")
	roleRoutes.HandleFunc("/{id}", RoleHandler.Edit).Methods(http.MethodPut).Name("role.edit")
	roleRoutes.HandleFunc("/{id}", RoleHandler.Delete).Methods(http.MethodDelete).Name("role.delete")
}
