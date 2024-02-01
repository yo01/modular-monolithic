package postgresql

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/user/dto"
	"modular-monolithic/security/middleware"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"
)

type IUserPostgre interface {
	Select() (resp []model.User, merr merror.Error)
	SelectByID(id string) (resp *model.User, merr merror.Error)
	Insert(data dto.CreateUserRequest) (merr merror.Error)
	Update(data dto.UpdateUserRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)

	// ADDITIONAL
	SelectByEmail(email string) (resp *model.User, merr merror.Error)
}

type userPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewUserPostgre(carrier *mcarrier.Carrier) userPostgre {
	return userPostgre{
		Carrier: carrier,
	}
}

func (r userPostgre) Select() (resp []model.User, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_USER)
	if err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.StructScan(&user); err != nil {
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return users, merr
}

func (r userPostgre) SelectByID(id string) (resp *model.User, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_USER_BY_ID, id)
	if err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var user model.User

	for row.Next() {
		if err := row.StructScan(&user); err != nil {
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &user, merr
}

func (r userPostgre) Insert(data dto.CreateUserRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_USER, id, data.Email, &data.Password, data.FullName, data.RoleID)
	if row == nil {
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r userPostgre) Update(data dto.UpdateUserRequest, id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_USER, id, data.FullName, authUser.ID, authUser.FullName)
	if row == nil {
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r userPostgre) Destroy(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_USER, id, authUser.ID, authUser.FullName)
	if row == nil {
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

// ADDITIONAL
func (r userPostgre) SelectByEmail(email string) (resp *model.User, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_USER_BY_EMAIL, email)
	if err != nil {
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var user model.User

	for row.Next() {
		if err := row.StructScan(&user); err != nil {
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &user, merr
}
