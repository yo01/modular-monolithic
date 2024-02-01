package user

import (
	"modular-monolithic/app"
	userService "modular-monolithic/module/v1/user/service"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"github.com/gorilla/mux"
)

type HandlerConfig struct {
	R           *mux.Router
	Carrier     *mcarrier.Carrier
	UserService userService.IUserService
}

// Inject Dependencies
func Inject(appConfig app.AppConfig) {
	// init service
	userSvc := userService.NewUserService(appConfig.Carrier)

	// init handler
	InitRoutes(HandlerConfig{
		Carrier:     appConfig.Carrier,
		R:           appConfig.Router,
		UserService: userSvc,
	})
}
