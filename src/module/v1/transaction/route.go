package transaction

import (
	"net/http"

	"modular-monolithic/module/v1/transaction/handler"
	"modular-monolithic/security/middleware"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	TransactionHandler := handler.TransactionHandler{
		Carrier:            c.Carrier,
		TransactionService: c.TransactionService,
	}

	// TRANSACTION ROUTES WITH MIDDLEWARE
	transactionRoutesWithMiddleware := c.R.PathPrefix("/transactions").Subrouter()
	transactionRoutesWithMiddleware.Use(middleware.JWT)

	transactionRoutesWithMiddleware.HandleFunc("", TransactionHandler.List).Methods(http.MethodGet).Name("transaction.list")
	transactionRoutesWithMiddleware.HandleFunc("/{id}", TransactionHandler.Detail).Methods(http.MethodGet).Name("transaction.detail")
	transactionRoutesWithMiddleware.HandleFunc("", TransactionHandler.Create).Methods(http.MethodPost).Name("transaction.save")
	// transactionRoutesWithMiddleware.HandleFunc("/{id}", TransactionHandler.Edit).Methods(http.MethodPut).Name("transaction.edit")
	transactionRoutesWithMiddleware.HandleFunc("/{id}", TransactionHandler.Delete).Methods(http.MethodDelete).Name("transaction.delete")
	// transactionRoutesWithMiddleware.HandleFunc("/{id}/payment", TransactionHandler.Payment).Methods(http.MethodPut).Name("transaction.payment")

	// TRANSACTION ROUTES WITHOUT MIDDLEWARE
}
