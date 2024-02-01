package postgresql

import (
	"fmt"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/transaction/dto"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"
)

type ITransactionPostgre interface {
	Select() (resp []model.Transaction, merr merror.Error)
	SelectByID(id string) (resp *model.Transaction, merr merror.Error)
	Insert(data dto.CreateTransactionRequest) (merr merror.Error)
	Update(data dto.UpdateTransactionRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)

	// ADDITIONAL
	Payment(id string) (merr merror.Error)
}

type transactionPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewTransactionPostgre(carrier *mcarrier.Carrier) transactionPostgre {
	return transactionPostgre{
		Carrier: carrier,
	}
}

func (r transactionPostgre) Select() (resp []model.Transaction, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_TRANSACTION)
	if err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer rows.Close()

	var transactions []model.Transaction

	for rows.Next() {
		var transaction model.Transaction
		if err := rows.StructScan(&transaction); err != nil {
			return nil, merror.Error{
				Error: err,
			}
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return transactions, merr
}

func (r transactionPostgre) SelectByID(id string) (resp *model.Transaction, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_TRANSACTION_BY_ID, id)
	if err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var transaction model.Transaction

	for row.Next() {
		if err := row.StructScan(&transaction); err != nil {
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &transaction, merr
}

func (r transactionPostgre) Insert(data dto.CreateTransactionRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_TRANSACTION, id, data.Name)
	if row == nil {
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}

func (r transactionPostgre) Update(data dto.UpdateTransactionRequest, id string) (merr merror.Error) {
	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_TRANSACTION, id, data.Name)
	if row == nil {
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}

func (r transactionPostgre) Destroy(id string) (merr merror.Error) {
	row, _ := r.Carrier.Library.Sqlx.Exec(DELETE_TRANSACTION, id)

	rowInt, _ := row.RowsAffected()
	if rowInt == 0 {
		return merror.Error{
			Code:  404,
			Error: fmt.Errorf("No transaction found with ID %v to delete", id),
		}
	}

	return merr
}

func (r transactionPostgre) Payment(id string) (merr merror.Error) {
	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_TRANSACTION_PAYMENT, id, true)
	if row == nil {
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}
