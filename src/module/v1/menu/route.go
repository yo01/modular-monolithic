package menu

import (
	"net/http"

	"modular-monolithic/module/v1/menu/handler"
	"modular-monolithic/security/middleware"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	MenuHandler := handler.MenuHandler{
		Carrier:     c.Carrier,
		MenuService: c.MenuService,
	}

	// MENU ROUTES WITH MIDDLEWARE
	menuRoutesWithMiddleware := c.R.PathPrefix("/menus").Subrouter()
	menuRoutesWithMiddleware.Use(middleware.JWT)

	menuRoutesWithMiddleware.HandleFunc("", MenuHandler.Create).Methods(http.MethodPost).Name("menu.save")
	menuRoutesWithMiddleware.HandleFunc("/{id}", MenuHandler.Edit).Methods(http.MethodPut).Name("menu.edit")
	menuRoutesWithMiddleware.HandleFunc("/{id}", MenuHandler.Delete).Methods(http.MethodDelete).Name("menu.delete")

	// MENU ROUTES WITHOUT MIDDLEWARE
	menuRoutesWitouthMiddleware := c.R.PathPrefix("/menus").Subrouter()

	menuRoutesWitouthMiddleware.HandleFunc("", MenuHandler.List).Methods(http.MethodGet).Name("menu.list")
	menuRoutesWitouthMiddleware.HandleFunc("/{id}", MenuHandler.Detail).Methods(http.MethodGet).Name("menu.detail")
}
