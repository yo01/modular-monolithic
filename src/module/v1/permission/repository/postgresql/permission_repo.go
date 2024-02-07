package postgresql

import (
	"net/http"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/permission/dto"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type IPermissionPostgre interface {
	Select() (resp []model.Permission, merr merror.Error)
	SelectByID(id string) (resp *model.Permission, merr merror.Error)
	Insert(data dto.CreatePermissionRequest) (merr merror.Error)
	Update(data dto.UpdatePermissionRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)

	// ADDITIONAL
	SelectByRoleID(roleID string) (resp *model.Permission, merr merror.Error)
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
	sqlQuery += utils.BuildQuery(pageRequest, "p", nil)

	rows, err := r.Carrier.Library.Sqlx.Queryx(sqlQuery)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
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
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
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
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}
	defer row.Close()

	var role model.Permission

	for row.Next() {
		if err := row.StructScan(&role); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
	}

	return &role, merr
}

func (r permissionPostgre) Insert(data dto.CreatePermissionRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	if _, err := r.Carrier.Library.Sqlx.Queryx(INSERT_PERMISSION, id, data.RoleID, pq.StringArray(data.ListAPI)); err != nil {
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	return merr
}

func (r permissionPostgre) Update(data dto.UpdatePermissionRequest, id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	_, err := r.Carrier.Library.Sqlx.Queryx(UPDATE_PERMISSION, id, data.RoleID, pq.StringArray(data.ListAPI), authUser.ID)
	if err != nil {
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	return merr
}

func (r permissionPostgre) Destroy(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	_, err := r.Carrier.Library.Sqlx.Queryx(SOFT_DELETE_PERMISSION, id, authUser.ID)
	if err == nil {
		zap.S().Error(err)
		return merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}

	return merr
}

func (r permissionPostgre) SelectByRoleID(roleID string) (resp *model.Permission, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_PERMISSION_BY_ROLE_ID, roleID)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  http.StatusInternalServerError,
			Error: err,
		}
	}
	defer row.Close()

	var role model.Permission

	for row.Next() {
		if err := row.StructScan(&role); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  http.StatusInternalServerError,
				Error: err,
			}
		}
	}

	return &role, merr
}
