package menu

import (
	"modular-monolithic/app"

	menuService "modular-monolithic/module/v1/menu/service"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"github.com/gorilla/mux"
)

type HandlerConfig struct {
	R           *mux.Router
	Carrier     *mcarrier.Carrier
	MenuService menuService.IMenuService
}

// Inject Dependencies
func Inject(appConfig app.AppConfig) {
	// init service
	menuSvc := menuService.NewMenuService(appConfig.Carrier)

	// init handler
	InitRoutes(HandlerConfig{
		Carrier:     appConfig.Carrier,
		R:           appConfig.Router,
		MenuService: menuSvc,
	})
}
