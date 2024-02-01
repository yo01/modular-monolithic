package transaction

import (
	"net/http"

	"modular-monolithic/module/v1/transaction/handler"
)

// InitRoutes for the module
func InitRoutes(c HandlerConfig) {
	TransactionHandler := handler.TransactionHandler{
		Carrier:            c.Carrier,
		TransactionService: c.TransactionService,
	}

	//User Register
	transactionRoutes := c.R.PathPrefix("/transactions").Subrouter()

	transactionRoutes.HandleFunc("/", TransactionHandler.List).Methods(http.MethodGet).Name("transaction.list")
	transactionRoutes.HandleFunc("/{id}", TransactionHandler.Detail).Methods(http.MethodGet).Name("transaction.detail")
	transactionRoutes.HandleFunc("/", TransactionHandler.Create).Methods(http.MethodPost).Name("transaction.save")
	transactionRoutes.HandleFunc("/{id}", TransactionHandler.Edit).Methods(http.MethodPut).Name("transaction.edit")
	transactionRoutes.HandleFunc("/{id}", TransactionHandler.Delete).Methods(http.MethodDelete).Name("transaction.delete")

	// ADDITIONAL
	transactionRoutes.HandleFunc("/{id}/payment", TransactionHandler.Payment).Methods(http.MethodPut).Name("transaction.payment")
}
