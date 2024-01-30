package postgresql

import (
	"fmt"
	"modular-monolithic/model"
	"modular-monolithic/module/v1/permission/dto"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
	"github.com/google/uuid"
)

type IPermissionPostgre interface {
	Select() (resp []model.Permission, merr merror.Error)
	SelectByID(id string) (resp *model.Permission, merr merror.Error)
	Insert(data dto.CreatePermissionRequest) (merr merror.Error)
	Update(data dto.UpdatePermissionRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)
}

type permissionPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewPermissionPostgre(carrier *mcarrier.Carrier) permissionPostgre {
	return permissionPostgre{
		Carrier: carrier,
	}
}

func (r permissionPostgre) Select() (resp []model.Permission, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_PERMISSION)
	if err != nil {
		return nil, merror.Error{
			Error: err,
		}
	}
	defer rows.Close()

	var permissions []model.Permission

	for rows.Next() {
		var permission model.Permission
		if err := rows.StructScan(&permission); err != nil {
			return nil, merror.Error{
				Error: err,
			}
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, merror.Error{
			Error: err,
		}
	}

	return permissions, merr
}

func (r permissionPostgre) SelectByID(id string) (resp *model.Permission, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_PERMISSION_BY_ID, id)
	if err != nil {
		return nil, merror.Error{
			Error: err,
		}
	}
	defer row.Close()

	var role model.Permission

	for row.Next() {
		if err := row.StructScan(&role); err != nil {
			return nil, merror.Error{
				Error: err,
			}
		}
	}

	return &role, merr
}

func (r permissionPostgre) Insert(data dto.CreatePermissionRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_PERMISSION, id, data.Name)
	if row == nil {
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}

func (r permissionPostgre) Update(data dto.UpdatePermissionRequest, id string) (merr merror.Error) {
	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_PERMISSION, id, data.Name)
	if row == nil {
		return merror.Error{
			Error: row.Err(),
		}
	}

	return merr
}

func (r permissionPostgre) Destroy(id string) (merr merror.Error) {
	row, _ := r.Carrier.Library.Sqlx.Exec(DELETE_PERMISSION, id)

	rowInt, _ := row.RowsAffected()
	if rowInt == 0 {
		return merror.Error{
			Error: fmt.Errorf("No permission found with ID %v to delete", id),
		}
	}

	return merr
}
