package postgresql

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/permission/dto"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

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
	// GET DATA FROM CONTEXT MIDDLEWARE
	pageRequest := r.Carrier.Context.Value(middleware.PageRequestCtxKey).(*model.PageRequest)

	// MAIN VARIABLE
	sqlQuery := SELECT_PERMISSION
	sqlQuery += utils.BuildQuery(pageRequest, "p", []string{
		"p.name",
	})

	rows, err := r.Carrier.Library.Sqlx.Queryx(sqlQuery)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer rows.Close()

	var permissions []model.Permission

	for rows.Next() {
		var permission model.Permission
		if err := rows.StructScan(&permission); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return permissions, merr
}

func (r permissionPostgre) SelectByID(id string) (resp *model.Permission, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_PERMISSION_BY_ID, id)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var role model.Permission

	for row.Next() {
		if err := row.StructScan(&role); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
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
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r permissionPostgre) Update(data dto.UpdatePermissionRequest, id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_PERMISSION, id, data.Name, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r permissionPostgre) Destroy(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_PERMISSION, id, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}
