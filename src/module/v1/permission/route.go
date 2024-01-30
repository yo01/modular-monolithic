package permission

import (
	"modular-monolithic/module/v1/permission/handler"
	"net/http"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	PermissionHandler := handler.PermissionHandler{
		Carrier:           c.Carrier,
		PermissionService: c.PermissionService,
	}

	//User Register
	permissionRoutes := c.R.PathPrefix("/permissions").Subrouter()

	permissionRoutes.HandleFunc("/", PermissionHandler.List).Methods(http.MethodGet).Name("permission.list")
	permissionRoutes.HandleFunc("/{id}", PermissionHandler.Detail).Methods(http.MethodGet).Name("permission.detail")
	permissionRoutes.HandleFunc("/", PermissionHandler.Create).Methods(http.MethodPost).Name("permission.save")
	permissionRoutes.HandleFunc("/{id}", PermissionHandler.Edit).Methods(http.MethodPut).Name("permission.edit")
	permissionRoutes.HandleFunc("/{id}", PermissionHandler.Delete).Methods(http.MethodDelete).Name("permission.delete")
}
