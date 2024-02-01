package product

import (
	"modular-monolithic/app"
	productService "modular-monolithic/module/v1/product/service"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"github.com/gorilla/mux"
)

type HandlerConfig struct {
	R              *mux.Router
	Carrier        *mcarrier.Carrier
	ProductService productService.IProductService
}

// Inject Dependencies
func Inject(appConfig app.AppConfig) {
	// init service
	productSvc := productService.NewProductService(appConfig.Carrier)

	// init handler
	InitRoutes(HandlerConfig{
		Carrier:        appConfig.Carrier,
		R:              appConfig.Router,
		ProductService: productSvc,
	})
}
