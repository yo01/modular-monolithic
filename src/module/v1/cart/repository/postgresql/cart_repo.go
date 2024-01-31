package postgresql

import (
	"fmt"
	"modular-monolithic/model"
	"modular-monolithic/module/v1/cart/dto"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
	"github.com/google/uuid"
)

type ICartPostgre interface {
	Select() (resp []model.Cart, merr merror.Error)
	SelectByID(id string) (resp *model.Cart, merr merror.Error)
	Insert(data dto.CreateCartRequest) (merr merror.Error)
	Update(data dto.UpdateCartRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)
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

func (r cartPostgre) SelectByID(id string) (resp *model.Cart, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_CART_BY_ID, id)
	if err != nil {
		return nil, merror.Error{
			Error: err,
		}
	}
	defer row.Close()

	var cart model.Cart

	for row.Next() {
		if err := row.StructScan(&cart); err != nil {
			return nil, merror.Error{
				Error: err,
			}
		}
	}

	return &cart, merr
}

func (r cartPostgre) Insert(data dto.CreateCartRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_CART, id, data.ProductID)
	if row == nil {
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}

func (r cartPostgre) Update(data dto.UpdateCartRequest, id string) (merr merror.Error) {
	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_CART, id, data.ProductID)
	if row == nil {
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}

func (r cartPostgre) Destroy(id string) (merr merror.Error) {
	row, _ := r.Carrier.Library.Sqlx.Exec(DELETE_CART, id)

	rowInt, _ := row.RowsAffected()
	if rowInt == 0 {
		return merror.Error{
			Error: fmt.Errorf("No cart found with ID %v to delete", id),
		}
	}

	return merr
}
