package postgresql

import (
	"fmt"
	"strings"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/transaction/dto"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type ITransactionPostgre interface {
	Select(pagination *model.PageRequest) (resp []model.Transaction, merr merror.Error)
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

func (r transactionPostgre) Select(pagination *model.PageRequest) (resp []model.Transaction, merr merror.Error) {
	// MAIN VARIABLE
	sqlQuery := SELECT_TRANSACTION
	offset := (pagination.Page - 1) * pagination.PerPage

	for _, filter := range pagination.Filter {
		// Loop through inner map
		for key, valueMap := range filter {
			for operator, value := range valueMap {
				sqlQuery += fmt.Sprintf(" AND %s %s %s", fmt.Sprintf("t.%v", key), operator, utils.GetSQLValue(operator, value))
			}
		}
	}

	if pagination.Search != "" {
		// MAIN VARIABLE
		fields := []string{
			"t.is_success_payment",
			"t.invoice_number",
		}

		if len(fields) == 1 {
			sqlQuery += fmt.Sprintf("AND %s ILIKE '%%%s%%' ", fields[0], pagination.Search)
		} else {
			var conditions []string

			// Add conditions based on non-empty filter criteria
			for _, field := range fields {
				condition := fmt.Sprintf("%s %s '%%%s%%'", field, "ILIKE", pagination.Search)
				conditions = append(conditions, condition)
			}

			// Combine conditions with OR and wrap in parentheses
			conditionClause := "(" + strings.Join(conditions, " OR ") + ")"

			sqlQuery += "AND " + conditionClause + " "
		}
	}

	if pagination.Sort != "" {
		sqlQuery += fmt.Sprintf("ORDER BY %v ", fmt.Sprintf("t.%v", pagination.Sort))
	}

	if pagination.PerPage != 0 {
		sqlQuery += fmt.Sprintf("LIMIT %v ", pagination.PerPage)
	}

	sqlQuery += fmt.Sprintf("OFFSET %v ", offset)

	rows, err := r.Carrier.Library.Sqlx.Queryx(sqlQuery)
	if err != nil {
		zap.S().Error(err)
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
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
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
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var transaction model.Transaction

	for row.Next() {
		if err := row.StructScan(&transaction); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &transaction, merr
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
			Code:  500,
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
			Code:  500,
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
			Code:  500,
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
