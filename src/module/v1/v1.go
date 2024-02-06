package v1

import (
	"modular-monolithic/app"
	"modular-monolithic/context"
	"modular-monolithic/module/v1/auth"
	"modular-monolithic/module/v1/cart"
	"modular-monolithic/module/v1/menu"
	"modular-monolithic/module/v1/permission"
	"modular-monolithic/module/v1/product"
	"modular-monolithic/module/v1/role"
	"modular-monolithic/module/v1/transaction"
	"modular-monolithic/module/v1/user"
)

func Inject(appConfig app.AppConfig) {
	// grouping api/v1
	appConfig.Router = appConfig.Router.PathPrefix("/api/v1").Subrouter()

	// IMPLEMENTATION PAGE REQUEST
	appConfig.Router.Use(context.PageRequestCtx)

	// user module
	user.Inject(appConfig)

	// role module
	role.Inject(appConfig)

	// permission module
	permission.Inject(appConfig)

	// menu module
	menu.Inject(appConfig)

	// product module
	product.Inject(appConfig)

	// transaction module
	transaction.Inject(appConfig)

	// cart module
	cart.Inject(appConfig)

	// auth module
	auth.Inject(appConfig)
}
