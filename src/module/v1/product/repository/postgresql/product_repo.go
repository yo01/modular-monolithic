package postgresql

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/product/dto"
	"modular-monolithic/security/middleware"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
	"github.com/google/uuid"
)

type IProductPostgre interface {
	Select() (resp []model.Product, merr merror.Error)
	SelectByID(id string) (resp *model.Product, merr merror.Error)
	Insert(data dto.CreateProductRequest) (merr merror.Error)
	Update(data dto.UpdateProductRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)
}

type productPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewProductPostgre(carrier *mcarrier.Carrier) productPostgre {
	return productPostgre{
		Carrier: carrier,
	}
}

func (r productPostgre) Select() (resp []model.Product, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_PRODUCT)
	if err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product
		if err := rows.StructScan(&product); err != nil {
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, merror.Error{
			Error: err,
		}
	}

	return products, merr
}

func (r productPostgre) SelectByID(id string) (resp *model.Product, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_PRODUCT_BY_ID, id)
	if err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var product model.Product

	for row.Next() {
		if err := row.StructScan(&product); err != nil {
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &product, merr
}

func (r productPostgre) Insert(data dto.CreateProductRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_PRODUCT, id, data.Name)
	if row == nil {
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r productPostgre) Update(data dto.UpdateProductRequest, id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_PRODUCT, id, data.Name, authUser.ID, authUser.FullName)
	if row == nil {
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r productPostgre) Destroy(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_PRODUCT, id, authUser.ID, authUser.FullName)
	if row == nil {
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}
