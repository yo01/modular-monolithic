package cart

import (
	"modular-monolithic/app"
	cartService "modular-monolithic/module/v1/cart/service"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"github.com/gorilla/mux"
)

type HandlerConfig struct {
	R           *mux.Router
	Carrier     *mcarrier.Carrier
	CartService cartService.ICartService
}

// Inject Dependencies
func Inject(appConfig app.AppConfig) {
	// init service
	cartSvc := cartService.NewCartService(appConfig.Carrier)

	// init handler
	InitRoutes(HandlerConfig{
		Carrier:     appConfig.Carrier,
		R:           appConfig.Router,
		CartService: cartSvc,
	})
}
