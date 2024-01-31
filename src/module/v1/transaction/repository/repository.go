package repository

import (
	transactionPostgre "modular-monolithic/module/v1/transaction/repository/postgresql"

	"git.motiolabs.com/library/motiolibs/mcarrier"
)

type TransactionRepository struct {
	Carrier            *mcarrier.Carrier
	TransactionPostgre transactionPostgre.ITransactionPostgre
}

func NewRepository(carrier *mcarrier.Carrier) TransactionRepository {
	transactionPostgre := transactionPostgre.NewTransactionPostgre(carrier)

	return TransactionRepository{
		Carrier:            carrier,
		TransactionPostgre: transactionPostgre,
	}
}
