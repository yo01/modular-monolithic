package postgresql

import (
	"fmt"
	"strings"

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
	Select(pagination *model.PageRequest) (resp []model.User, merr merror.Error)
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

func (r userPostgre) Select(pagination *model.PageRequest) (resp []model.User, merr merror.Error) {
	// MAIN VARIABLE
	sqlQuery := SELECT_USER
	offset := (pagination.Page - 1) * pagination.PerPage

	for _, filter := range pagination.Filter {
		// Loop through inner map
		for key, valueMap := range filter {
			for operator, value := range valueMap {
				sqlQuery += fmt.Sprintf(" AND %s %s %s", fmt.Sprintf("u.%v", key), operator, utils.GetSQLValue(operator, value))
			}
		}
	}

	if pagination.Search != "" {
		// MAIN VARIABLE
		fields := []string{
			"u.email",
			"u.full_name",
		}

		if len(fields) == 1 {
			sqlQuery += fmt.Sprintf("AND %s ILIKE '%%%s%%' ", fields[0], pagination.Search)
		} else {
			var conditions []string

			// Add conditions based on non-empty filter criteria
			for _, field := range fields {
				condition := fmt.Sprintf("%s %s '%%%s%%'", field, "ILIKE", pagination.Search)
				conditions = append(conditions, condition)
			}

			// Combine conditions with OR and wrap in parentheses
			conditionClause := "(" + strings.Join(conditions, " OR ") + ")"

			sqlQuery += "AND " + conditionClause + " "
		}
	}

	if pagination.Sort != "" {
		sqlQuery += fmt.Sprintf("ORDER BY %v ", fmt.Sprintf("u.%v", pagination.Sort))
	}

	if pagination.PerPage != 0 {
		sqlQuery += fmt.Sprintf("LIMIT %v ", pagination.PerPage)
	}

	sqlQuery += fmt.Sprintf("OFFSET %v ", offset)

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

func (r userPostgre) SelectByID(id string) (resp *model.User, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_USER_BY_ID, id)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var user model.User

	for row.Next() {
		if err := row.StructScan(&user); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &user, merr
}

func (r userPostgre) Insert(data dto.CreateUserRequest) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_USER, id, data.Email, &data.Password, data.FullName, data.RoleID, authUser.ID)
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
func (r userPostgre) SelectByEmail(email string) (resp *model.User, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_USER_BY_EMAIL, email)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var user model.User

	for row.Next() {
		if err := row.StructScan(&user); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &user, merr
}
