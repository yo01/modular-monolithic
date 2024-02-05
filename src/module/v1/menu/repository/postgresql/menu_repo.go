package postgresql

import (
	"fmt"
	"strings"

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
	Select(pagination *model.PageRequest) (resp []model.Menu, merr merror.Error)
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

func (r menuPostgre) Select(pagination *model.PageRequest) (resp []model.Menu, merr merror.Error) {
	// MAIN VARIABLE
	sqlQuery := SELECT_MENU
	offset := (pagination.Page - 1) * pagination.PerPage

	for _, filter := range pagination.Filter {
		// Loop through inner map
		for key, valueMap := range filter {
			for operator, value := range valueMap {
				sqlQuery += fmt.Sprintf(" AND %s %s %s", fmt.Sprintf("m.%v", key), operator, utils.GetSQLValue(operator, value))
			}
		}
	}

	if pagination.Search != "" {
		// MAIN VARIABLE
		fields := []string{
			"m.name",
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
		sqlQuery += fmt.Sprintf("ORDER BY %v ", fmt.Sprintf("m.%v", pagination.Sort))
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

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_MENU, id, data.Name, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
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
