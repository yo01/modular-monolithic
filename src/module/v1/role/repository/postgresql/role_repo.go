package postgresql

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/role/dto"
	"modular-monolithic/security/middleware"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"
)

type IRolePostgre interface {
	Select() (resp []model.Role, merr merror.Error)
	SelectByID(id string) (resp *model.Role, merr merror.Error)
	Insert(data dto.CreateRoleRequest) (merr merror.Error)
	Update(data dto.UpdateRoleRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)
}

type rolePostgre struct {
	Carrier *mcarrier.Carrier
}

func NewRolePostgre(carrier *mcarrier.Carrier) rolePostgre {
	return rolePostgre{
		Carrier: carrier,
	}
}

func (r rolePostgre) Select() (resp []model.Role, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_ROLE)
	if err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer rows.Close()

	var roles []model.Role

	for rows.Next() {
		var role model.Role
		if err := rows.StructScan(&role); err != nil {
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return roles, merr
}

func (r rolePostgre) SelectByID(id string) (resp *model.Role, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_ROLE_BY_ID, id)
	if err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var role model.Role

	for row.Next() {
		if err := row.StructScan(&role); err != nil {
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &role, merr
}

func (r rolePostgre) Insert(data dto.CreateRoleRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_ROLE, id, data.Name)
	if row == nil {
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r rolePostgre) Update(data dto.UpdateRoleRequest, id string) (merr merror.Error) {
	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_ROLE, id, data.Name)
	if row == nil {
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r rolePostgre) Destroy(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_ROLE, id, authUser.ID, authUser.FullName)
	if row == nil {
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}
