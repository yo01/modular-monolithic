package permission

import (
	"modular-monolithic/app"
	permissionService "modular-monolithic/module/v1/permission/service"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"github.com/gorilla/mux"
)

type HandlerConfig struct {
	R                 *mux.Router
	Carrier           *mcarrier.Carrier
	PermissionService permissionService.IPermissionService
}

// Inject Dependencies
func Inject(appConfig app.AppConfig) {
	// init service
	permissionSvc := permissionService.NewRoleService(appConfig.Carrier)

	// init handler
	InitRoutes(HandlerConfig{
		Carrier:           appConfig.Carrier,
		R:                 appConfig.Router,
		PermissionService: permissionSvc,
	})
}
