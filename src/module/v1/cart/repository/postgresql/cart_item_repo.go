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

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type ICartItemPostgre interface {
	Select(pagination *model.PageRequest) (resp []model.CartItem, merr merror.Error)
	SelectByID(id string) (resp *model.CartItem, merr merror.Error)
	Insert(data dto.CreateCartItemRequest) (merr merror.Error)
	Update(data dto.UpdateCartItemRequest, id, cartID string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)
}

type cartItemPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewCartItemPostgre(carrier *mcarrier.Carrier) cartItemPostgre {
	return cartItemPostgre{
		Carrier: carrier,
	}
}

func (r cartItemPostgre) Select(pagination *model.PageRequest) (resp []model.CartItem, merr merror.Error) {
	// MAIN VARIABLE
	sqlQuery := SELECT_CART_ITEM
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
		fields := []string{}

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

	var CartItems []model.CartItem

	for rows.Next() {
		var cartItem model.CartItem
		if err := rows.StructScan(&cartItem); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		CartItems = append(CartItems, cartItem)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return CartItems, merr
}

func (r cartItemPostgre) SelectByID(id string) (resp *model.CartItem, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_CART_ITEM_BY_ID, id)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var cartItem model.CartItem

	for row.Next() {
		if err := row.StructScan(&cartItem); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &cartItem, merr
}

func (r cartItemPostgre) Insert(data dto.CreateCartItemRequest) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	// GENERATE NEW UUID
	id := uuid.New()
	cartUUID := utils.ParseStringToUUID(data.CartID)
	productUUID := utils.ParseStringToUUID(data.ProductID)

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_CART_ITEM, id, cartUUID, productUUID, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r cartItemPostgre) Update(data dto.UpdateCartItemRequest, id, cartID string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_CART_ITEM, id, cartID, data.ProductID, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r cartItemPostgre) Destroy(cartID string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_CART_ITEM, cartID, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}
