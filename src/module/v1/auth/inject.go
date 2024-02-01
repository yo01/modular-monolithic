package auth

import (
	"modular-monolithic/app"
	authService "modular-monolithic/module/v1/auth/service"
	userService "modular-monolithic/module/v1/user/service"

	"git.motiolabs.com/library/motiolibs/mcarrier"

	"github.com/gorilla/mux"
)

type HandlerConfig struct {
	R           *mux.Router
	Carrier     *mcarrier.Carrier
	AuthService authService.IAuthService
	UserService userService.IUserService
}

// Inject Dependencies
func Inject(appConfig app.AppConfig) {
	// init service
	authSvc := authService.NewAuthService(appConfig.Carrier)
	userSvc := userService.NewUserService(appConfig.Carrier)

	// init handler
	InitRoutes(HandlerConfig{
		Carrier:     appConfig.Carrier,
		R:           appConfig.Router,
		AuthService: authSvc,
		UserService: userSvc,
	})
}
