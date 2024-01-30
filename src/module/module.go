package module

import (
	"modular-monolithic/app"
	v1 "modular-monolithic/module/v1"

	"git.motiolabs.com/library/motiolibs/mhttp"
)

func Inject(appConfig app.AppConfig) {
	//health check
	appConfig.Router.HandleFunc("/health", mhttp.HealthCheck)

	v1.Inject(appConfig)
}
