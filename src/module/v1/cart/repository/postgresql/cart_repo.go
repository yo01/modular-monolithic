package postgresql

import (
	"fmt"
	"strings"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/cart/dto"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
	"github.com/google/uuid"

	"go.uber.org/zap"
)

type ICartPostgre interface {
	Select(pagination *model.PageRequest) (resp []model.Cart, merr merror.Error)
	SelectByID(id string) (resp []model.Cart, merr merror.Error)
	Insert(data dto.CreateCartRequest) (resp *model.Cart, merr merror.Error)
	Update(data dto.UpdateCartRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)

	SelectOneByID(id string) (resp *model.Cart, merr merror.Error)
}

type cartPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewCartPostgre(carrier *mcarrier.Carrier) cartPostgre {
	return cartPostgre{
		Carrier: carrier,
	}
}

func (r cartPostgre) Select(pagination *model.PageRequest) (resp []model.Cart, merr merror.Error) {
	// MAIN VARIABLE
	sqlQuery := SELECT_CART
	offset := (pagination.Page - 1) * pagination.PerPage

	for _, filter := range pagination.Filter {
		// Loop through inner map
		for key, valueMap := range filter {
			for operator, value := range valueMap {
				sqlQuery += fmt.Sprintf(" AND %s %s %s", fmt.Sprintf("c.%v", key), operator, utils.GetSQLValue(operator, value))
			}
		}
	}

	if pagination.Search != "" {
		// MAIN VARIABLE
		fields := []string{
			"c.product_name",
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
		sqlQuery += fmt.Sprintf("ORDER BY %v ", fmt.Sprintf("c.%v", pagination.Sort))
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

	var Carts []model.Cart

	for rows.Next() {
		var cart model.Cart
		if err := rows.StructScan(&cart); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		Carts = append(Carts, cart)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return Carts, merr
}

func (r cartPostgre) SelectByID(id string) (resp []model.Cart, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_CART_BY_ID, id)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer rows.Close()

	var Carts []model.Cart

	for rows.Next() {
		var cart model.Cart
		if err := rows.StructScan(&cart); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		Carts = append(Carts, cart)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return Carts, merr
}

func (r cartPostgre) SelectOneByID(id string) (resp *model.Cart, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_ONE_CART_BY_ID, id)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer rows.Close()

	var cart model.Cart

	for rows.Next() {
		if err := rows.StructScan(&cart); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return &cart, merr
}

func (r cartPostgre) Insert(data dto.CreateCartRequest) (resp *model.Cart, merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User
	id := uuid.New()

	row, err := r.Carrier.Library.Sqlx.Queryx(INSERT_CART, id, authUser.ID)
	if err != nil {
		zap.S().Error(err)
		return resp, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var cart model.Cart

	for row.Next() {
		if err := row.StructScan(&cart); err != nil {
			zap.S().Error(err)
			return resp, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &cart, merr
}

func (r cartPostgre) Update(data dto.UpdateCartRequest, id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_CART, id, data.ProductID, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r cartPostgre) Destroy(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_CART, id, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}
