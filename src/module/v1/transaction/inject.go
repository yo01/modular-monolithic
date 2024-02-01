package transaction

import (
	"modular-monolithic/app"
	transactionService "modular-monolithic/module/v1/transaction/service"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"github.com/gorilla/mux"
)

type HandlerConfig struct {
	R                  *mux.Router
	Carrier            *mcarrier.Carrier
	TransactionService transactionService.ITransactionService
}

// Inject Dependencies
func Inject(appConfig app.AppConfig) {
	// init service
	transactionSvc := transactionService.NewTransactionService(appConfig.Carrier)

	// init handler
	InitRoutes(HandlerConfig{
		Carrier:            appConfig.Carrier,
		R:                  appConfig.Router,
		TransactionService: transactionSvc,
	})
}
