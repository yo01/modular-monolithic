package postgresql

import (
	"fmt"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/menu/dto"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

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
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_MENU)
	if err != nil {
		return nil, merror.Error{
			Error: err,
		}
	}
	defer rows.Close()

	var Menus []model.Menu

	for rows.Next() {
		var menu model.Menu
		if err := rows.StructScan(&menu); err != nil {
			return nil, merror.Error{
				Error: err,
			}
		}
		Menus = append(Menus, menu)
	}

	if err := rows.Err(); err != nil {
		return nil, merror.Error{
			Error: err,
		}
	}

	return Menus, merr
}

func (r menuPostgre) SelectByID(id string) (resp *model.Menu, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_MENU_BY_ID, id)
	if err != nil {
		return nil, merror.Error{
			Error: err,
		}
	}
	defer row.Close()

	var menu model.Menu

	for row.Next() {
		if err := row.StructScan(&menu); err != nil {
			return nil, merror.Error{
				Error: err,
			}
		}
	}

	return &menu, merr
}

func (r menuPostgre) Insert(data dto.CreateMenuRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_MENU, id, data.Name)
	if row == nil {
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}

func (r menuPostgre) Update(data dto.UpdateMenuRequest, id string) (merr merror.Error) {
	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_MENU, id, data.Name)
	if row == nil {
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}

func (r menuPostgre) Destroy(id string) (merr merror.Error) {
	row, _ := r.Carrier.Library.Sqlx.Exec(DELETE_MENU, id)

	rowInt, _ := row.RowsAffected()
	if rowInt == 0 {
		return merror.Error{
			Error: fmt.Errorf("No menu found with ID %v to delete", id),
		}
	}

	return merr
}
