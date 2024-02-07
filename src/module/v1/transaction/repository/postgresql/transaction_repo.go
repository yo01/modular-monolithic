package postgresql

import (
	"net/http"

	"modular-monolithic/constant"
	"modular-monolithic/model"
	"modular-monolithic/module/v1/transaction/dto"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ITransactionPostgre interface {
	Select() (resp []model.Transaction, merr merror.Error)
	SelectByID(id string) (resp []model.Transaction, merr merror.Error)
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
	// GET DATA FROM CONTEXT MIDDLEWARE
	pageRequest := r.Carrier.Context.Value(middleware.PageRequestCtxKey).(*model.PageRequest)
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)

	// MAIN VARIABLE
	sqlQuery := ""
	var rows *sqlx.Rows

	sqlQuery += utils.BuildQuery(pageRequest, "t", []string{
		"t.is_success_payment",
		"t.invoice_number",
	})

	if auth.User.Role.Name == constant.RoleAdmin {
		sqlQuery = SELECT_TRANSACTION_ADMIN

		result, err := r.Carrier.Library.Sqlx.Queryx(sqlQuery)
		if err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
		rows = result
		defer rows.Close()
	} else if auth.User.Role.Name == constant.RoleLearner {
		sqlQuery = SELECT_TRANSACTION_LEARNER

		result, err := r.Carrier.Library.Sqlx.Queryx(sqlQuery, auth.User.ID)
		if err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
		rows = result
		defer rows.Close()
	}

	var transactions []model.Transaction

	for rows.Next() {
		var transaction model.Transaction
		if err := rows.StructScan(&transaction); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	return transactions, merr
}

func (r transactionPostgre) SelectByID(id string) (resp []model.Transaction, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_TRANSACTION_BY_ID, id)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}
	defer rows.Close()

	var transactions []model.Transaction

	for rows.Next() {
		var transaction model.Transaction
		if err := rows.StructScan(&transaction); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	return transactions, merr
}

func (r transactionPostgre) Insert(data dto.CreateTransactionRequest) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	// GENERATE NEW UUID
	id := uuid.New()
	cartUUID := utils.ParseStringToUUID(data.CartID)

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_TRANSACTION, id, cartUUID, true, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: row.Err(),
		}
	}

	return merr
}

func (r transactionPostgre) Update(data dto.UpdateTransactionRequest, id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_TRANSACTION, id, "", authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: row.Err(),
		}
	}

	return merr
}

func (r transactionPostgre) Destroy(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_TRANSACTION, id, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: row.Err(),
		}
	}

	return merr
}

func (r transactionPostgre) Payment(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_TRANSACTION_PAYMENT, id, authUser.ID, true)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}
