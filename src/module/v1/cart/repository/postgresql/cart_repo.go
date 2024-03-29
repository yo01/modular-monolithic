package postgresql

import (
	"net/http"

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
	Select() (resp []model.Cart, merr merror.Error)
	SelectByID(id string) (resp []model.Cart, merr merror.Error)
	Insert(data dto.CreateCartRequest) (resp *model.Cart, merr merror.Error)
	Update(data dto.UpdateCartRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)

	// ADDITIONAL
	SelectOneByID(id string) (resp *model.Cart, merr merror.Error)
	UpdateFlagIsSuccess(isSuccess bool, id string) (merr merror.Error)
}

type cartPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewCartPostgre(carrier *mcarrier.Carrier) cartPostgre {
	return cartPostgre{
		Carrier: carrier,
	}
}

func (r cartPostgre) Select() (resp []model.Cart, merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	pageRequest := r.Carrier.Context.Value(middleware.PageRequestCtxKey).(*model.PageRequest)

	// MAIN VARIABLE
	sqlQuery := SELECT_CART
	sqlQuery += utils.BuildQuery(pageRequest, "c", []string{
		"c.product_name",
	})

	rows, err := r.Carrier.Library.Sqlx.Queryx(sqlQuery)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
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
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
		Carts = append(Carts, cart)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
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
			Code:  http.StatusInternalServerError,
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
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
		Carts = append(Carts, cart)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
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
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}
	defer rows.Close()

	var cart model.Cart

	for rows.Next() {
		if err := rows.StructScan(&cart); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
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
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}
	defer row.Close()

	var cart model.Cart

	for row.Next() {
		if err := row.StructScan(&cart); err != nil {
			zap.S().Error(err)
			return resp, merror.Error{
				Code:  http.StatusInternalServerError,
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
			Code:  http.StatusInternalServerError,
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
			Code:  http.StatusInternalServerError,
			Error: row.Err(),
		}
	}

	return merr
}

// ADDITIONAL
func (r cartPostgre) UpdateFlagIsSuccess(isSuccess bool, id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_CART_IS_SUCCESS, id, isSuccess, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: row.Err(),
		}
	}

	return merr
}
