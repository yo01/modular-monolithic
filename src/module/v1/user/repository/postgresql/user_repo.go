package postgresql

import (
	"modular-monolithic/model"
	"modular-monolithic/module/v1/user/dto"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"
	"go.uber.org/zap"

	"github.com/google/uuid"
)

type IUserPostgre interface {
	Select() (resp []model.User, merr merror.Error)
	SelectByID(id string) (resp []model.User, merr merror.Error)
	Insert(data dto.CreateUserRequest) (merr merror.Error)
	Update(data dto.UpdateUserRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)

	// ADDITIONAL
	SelectByEmail(email string) (resp []model.User, merr merror.Error)
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
	// GET DATA FROM CONTEXT MIDDLEWARE
	pageRequest := r.Carrier.Context.Value(middleware.PageRequestCtxKey).(*model.PageRequest)

	// MAIN VARIABLE
	sqlQuery := SELECT_USER
	sqlQuery += utils.BuildQuery(pageRequest, "u", []string{
		"u.email",
		"u.full_name",
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

	var users []model.User

	for rows.Next() {
		var user model.User
		if err := rows.StructScan(&user); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return users, merr
}

func (r userPostgre) SelectByID(id string) (resp []model.User, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_USER_BY_ID, id)
	if err != nil {
		zap.S().Error(err)
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
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return users, merr
}

func (r userPostgre) Insert(data dto.CreateUserRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_USER, id, data.Email, &data.Password, data.FullName, data.RoleID)
	if row == nil {
		zap.S().Error(row.Err())
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

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_USER, id, data.FullName, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
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

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_USER, id, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

// ADDITIONAL
func (r userPostgre) SelectByEmail(email string) (resp []model.User, merr merror.Error) {
	rows, err := r.Carrier.Library.Sqlx.Queryx(SELECT_USER_BY_EMAIL, email)
	if err != nil {
		zap.S().Error(err)
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
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}

	return users, merr
}
