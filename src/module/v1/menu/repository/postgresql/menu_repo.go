package postgresql

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/menu/dto"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
)

type IMenuPostgre interface {
	Select() (resp []model.Menu, merr merror.Error)
	SelectByID(id string) (resp *model.Menu, merr merror.Error)
	Insert(data dto.CreateMenuRequest) (merr merror.Error)
	Update(data dto.UpdateMenuRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)
}

type menuPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewMenuPostgre(carrier *mcarrier.Carrier) menuPostgre {
	return menuPostgre{
		Carrier: carrier,
	}
}

func (r menuPostgre) Select() (resp []model.Menu, merr merror.Error) {
	// GET DATA FROM CONTEXT MIDDLEWARE
	pageRequest := r.Carrier.Context.Value(middleware.PageRequestCtxKey).(*model.PageRequest)

	// MAIN VARIABLE
	sqlQuery := SELECT_MENU
	sqlQuery += utils.BuildQuery(pageRequest, "m", nil)

	rows, err := r.Carrier.Library.Sqlx.Queryx(sqlQuery)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer rows.Close()

	var Menus []model.Menu

	for rows.Next() {
		var menu model.Menu
		if err := rows.StructScan(&menu); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		Menus = append(Menus, menu)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Error: err,
		}
	}

	return Menus, merr
}

func (r menuPostgre) SelectByID(id string) (resp *model.Menu, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_MENU_BY_ID, id)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var menu model.Menu

	for row.Next() {
		if err := row.StructScan(&menu); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &menu, merr
}

func (r menuPostgre) Insert(data dto.CreateMenuRequest) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	// GENERATE NEW UUID
	id := uuid.New()

	_, err := r.Carrier.Library.Sqlx.Queryx(INSERT_MENU, id, data.Name, authUser.ID)
	if err != nil {
		zap.S().Error(err)
		return merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return merr
}

func (r menuPostgre) Update(data dto.UpdateMenuRequest, id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_MENU, id, data.Name, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r menuPostgre) Destroy(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_MENU, id, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}
