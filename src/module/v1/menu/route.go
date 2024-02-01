package menu

import (
	"net/http"

	"modular-monolithic/module/v1/menu/handler"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	MenuHandler := handler.MenuHandler{
		Carrier:     c.Carrier,
		MenuService: c.MenuService,
	}

	//User Register
	menuRoutes := c.R.PathPrefix("/menus").Subrouter()

	menuRoutes.HandleFunc("/", MenuHandler.List).Methods(http.MethodGet).Name("menu.list")
	menuRoutes.HandleFunc("/{id}", MenuHandler.Detail).Methods(http.MethodGet).Name("menu.detail")
	menuRoutes.HandleFunc("/", MenuHandler.Create).Methods(http.MethodPost).Name("menu.save")
	menuRoutes.HandleFunc("/{id}", MenuHandler.Edit).Methods(http.MethodPut).Name("menu.edit")
	menuRoutes.HandleFunc("/{id}", MenuHandler.Delete).Methods(http.MethodDelete).Name("menu.delete")
}
