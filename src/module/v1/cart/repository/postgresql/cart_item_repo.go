package postgresql

import (
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
	Select(pageRequest *model.PageRequest) (resp []model.CartItem, merr merror.Error)
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

func (r cartItemPostgre) Select(pageRequest *model.PageRequest) (resp []model.CartItem, merr merror.Error) {
	// MAIN VARIABLE
	sqlQuery := SELECT_CART_ITEM
	sqlQuery += utils.BuildQuery(pageRequest, "c", nil)

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

	if _, err := r.Carrier.Library.Sqlx.Queryx(INSERT_CART_ITEM, id, cartUUID, productUUID, authUser.ID); err != nil {
		zap.S().Error(err)
		return merror.Error{
			Code:  500,
			Error: err,
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
