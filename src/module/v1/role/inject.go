package role

import (
	"modular-monolithic/app"

	roleService "modular-monolithic/module/v1/role/service"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"github.com/gorilla/mux"
)

type HandlerConfig struct {
	R           *mux.Router
	Carrier     *mcarrier.Carrier
	RoleService roleService.IRoleService
}

// Inject Dependencies
func Inject(appConfig app.AppConfig) {
	// init service
	roleSvc := roleService.NewRoleService(appConfig.Carrier)

	// init handler
	InitRoutes(HandlerConfig{
		Carrier:     appConfig.Carrier,
		R:           appConfig.Router,
		RoleService: roleSvc,
	})
}
