package postgresql

import (
	"fmt"
	"strings"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/product/dto"
	"modular-monolithic/security/middleware"
	"modular-monolithic/utils"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/merror"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

type IProductPostgre interface {
	Select(pagination *model.PageRequest) (resp []model.Product, merr merror.Error)
	SelectByID(id string) (resp *model.Product, merr merror.Error)
	Insert(data dto.CreateProductRequest) (merr merror.Error)
	Update(data dto.UpdateProductRequest, id string) (merr merror.Error)
	Destroy(id string) (merr merror.Error)
}

type productPostgre struct {
	Carrier *mcarrier.Carrier
}

func NewProductPostgre(carrier *mcarrier.Carrier) productPostgre {
	return productPostgre{
		Carrier: carrier,
	}
}

func (r productPostgre) Select(pagination *model.PageRequest) (resp []model.Product, merr merror.Error) {
	// MAIN VARIABLE
	sqlQuery := SELECT_PRODUCT
	offset := (pagination.Page - 1) * pagination.PerPage

	for _, filter := range pagination.Filter {
		// Loop through inner map
		for key, valueMap := range filter {
			for operator, value := range valueMap {
				sqlQuery += fmt.Sprintf(" AND %s %s %s", fmt.Sprintf("p.%v", key), operator, utils.GetSQLValue(operator, value))
			}
		}
	}

	if pagination.Search != "" {
		// MAIN VARIABLE
		fields := []string{
			"p.name",
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
		sqlQuery += fmt.Sprintf("ORDER BY %v ", fmt.Sprintf("p.%v", pagination.Sort))
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

	var products []model.Product

	for rows.Next() {
		var product model.Product
		if err := rows.StructScan(&product); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Error: err,
		}
	}

	return products, merr
}

func (r productPostgre) SelectByID(id string) (resp *model.Product, merr merror.Error) {
	row, err := r.Carrier.Library.Sqlx.Queryx(SELECT_PRODUCT_BY_ID, id)
	if err != nil {
		zap.S().Error(err)
		return nil, merror.Error{
			Code:  500,
			Error: err,
		}
	}
	defer row.Close()

	var product model.Product

	for row.Next() {
		if err := row.StructScan(&product); err != nil {
			zap.S().Error(err)
			return nil, merror.Error{
				Code:  500,
				Error: err,
			}
		}
	}

	return &product, merr
}

func (r productPostgre) Insert(data dto.CreateProductRequest) (merr merror.Error) {
	// GENERATE NEW UUID
	id := uuid.New()

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, INSERT_PRODUCT, id, data.Name)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r productPostgre) Update(data dto.UpdateProductRequest, id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, UPDATE_PRODUCT, id, data.Name, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}

func (r productPostgre) Destroy(id string) (merr merror.Error) {
	// MAIN VARIABLE
	auth := r.Carrier.Context.Value(middleware.AuthUserCtxKey).(*model.Auth)
	authUser := auth.User

	row := r.Carrier.Library.Sqlx.QueryRowxContext(r.Carrier.Context, SOFT_DELETE_PRODUCT, id, authUser.ID)
	if row == nil {
		zap.S().Error(row.Err())
		return merror.Error{
			Code:  500,
			Error: row.Err(),
		}
	}

	return merr
}
