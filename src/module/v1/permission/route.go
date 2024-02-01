package permission

import (
	"net/http"

	"modular-monolithic/module/v1/permission/handler"
	"modular-monolithic/security/middleware"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	PermissionHandler := handler.PermissionHandler{
		Carrier:           c.Carrier,
		PermissionService: c.PermissionService,
	}

	// PERMISSION ROUTES WITH MIDDLEWARE
	permissionRoutesWithMiddleware := c.R.PathPrefix("/permissions").Subrouter()
	permissionRoutesWithMiddleware.Use(middleware.JWT)

	permissionRoutesWithMiddleware.HandleFunc("/", PermissionHandler.List).Methods(http.MethodGet).Name("permission.list")
	permissionRoutesWithMiddleware.HandleFunc("/{id}", PermissionHandler.Detail).Methods(http.MethodGet).Name("permission.detail")
	permissionRoutesWithMiddleware.HandleFunc("/", PermissionHandler.Create).Methods(http.MethodPost).Name("permission.save")
	permissionRoutesWithMiddleware.HandleFunc("/{id}", PermissionHandler.Edit).Methods(http.MethodPut).Name("permission.edit")
	permissionRoutesWithMiddleware.HandleFunc("/{id}", PermissionHandler.Delete).Methods(http.MethodDelete).Name("permission.delete")

	// PERMISSION ROUTES WITHOUT MIDDLEWARE
}
