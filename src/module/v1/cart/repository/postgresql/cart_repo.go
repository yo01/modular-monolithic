package postgresql

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/cart/dto"
	"modular-monolithic/security/middleware"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
	"github.com/google/uuid"
)

type ICartPostgre interface {
	Select() (resp []model.Cart, merr merror.Error)
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

func (r cartPostgre) Select() (resp []model.Cart, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_CART)
	if err != nil {
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
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		Carts = append(Carts, cart)
	}

	if err := rows.Err(); err != nil {
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
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		Carts = append(Carts, cart)
	}

	if err := rows.Err(); err != nil {
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
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer rows.Close()

	var cart model.Cart

	for rows.Next() {
		if err := rows.StructScan(&cart); err != nil {
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	if err := rows.Err(); err != nil {
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
		return resp, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var cart model.Cart

	for row.Next() {
		if err := row.StructScan(&cart); err != nil {
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
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}
