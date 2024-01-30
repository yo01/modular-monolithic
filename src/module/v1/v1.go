package v1

import (
	"modular-monolithic/app"
	"modular-monolithic/module/v1/role"
	"modular-monolithic/module/v1/user"
)

func Inject(appConfig app.AppConfig) {
	// grouping api/v1
	appConfig.Router = appConfig.Router.PathPrefix("/api/v1").Subrouter()

	// //middleware x-api-key
	// appConfig.Router.Use(func(next http.Handler) http.Handler {
	// 	return mmiddleware.ValidateAPIKey(appConfig.Config.AppApiKey, next)
	// })

	// user module
	user.Inject(appConfig)

	// role module
	role.Inject(appConfig)
}
