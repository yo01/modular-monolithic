package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"modular-monolithic/model"
)

func GeneratePaginationFromRequest(r *http.Request) model.PageRequest {
	// Initializing default
	perPage := 10
	page := 1
	paginate := 1
	sort := "created_at asc"
	search := ""
	filter := []map[string]map[string]interface{}{}
	query := r.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "per_page":
			perPage, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		case "filter[]":
			filter, _ = ValidateAndReturnFilterMap(value)
			break
		case "paginate":
			paginate, _ = strconv.Atoi(queryValue)
			break
		case "search":
			search = queryValue
		}
	}

	return model.PageRequest{
		Paginate: paginate,
		Page:     page,
		PerPage:  perPage,
		Sort:     sort,
		Search:   search,
		Filters:  filter,
	}
}

func ValidateAndReturnFilterMap(filter []string) ([]map[string]map[string]interface{}, error) {
	var mappingFilter []map[string]map[string]interface{}

	for _, v := range filter {
		splits := strings.Split(v, "|")

		// SET DATA
		if len(splits) > 0 {
			field, operator, value := splits[0], splits[1], splits[2]

			// SET DATA TO MAPPING FILTER
			mappingFilter = append(mappingFilter, CreateFilterFromMap(field, operator, value))
		}
	}

	return mappingFilter, nil
}

func CreateFilterFromMap(key, operator string, value interface{}) map[string]map[string]interface{} {
	// MAIN VARIABLE
	res := map[string]map[string]interface{}{}

	// CONVERT FROM STRING TO MAP
	innerMap := map[string]interface{}{}
	innerMap[operator] = value
	res[key] = innerMap

	return res
}

// GetSQLValue returns the SQL-formatted value based on the operator
func GetSQLValue(operator string, value interface{}) string {
	switch strings.ToUpper(operator) {
	case "=", ">", "<", ">=", "<=", "<>", "!=":
		return fmt.Sprintf("'%v'", value)
	case "LIKE", "ILIKE":
		return fmt.Sprintf("'%s'", "%"+fmt.Sprintf("%v", value)+"%")
	// Add more cases for other operators as needed
	default:
		return ""
	}
}

func BuildFilterCondition(filters []map[string]map[string]interface{}, initial string) string {
	var conditions []string

	for _, filter := range filters {
		// Loop through inner map
		for key, valueMap := range filter {
			for operator, value := range valueMap {
				condition := fmt.Sprintf("%s %s %s", fmt.Sprintf("%v.%v", initial, key), operator, GetSQLValue(operator, value))
				conditions = append(conditions, condition)
			}
		}
	}

	return strings.Join(conditions, " AND ")
}

func BuildSearchCondition(fields []string, searchValue string) string {
	if len(fields) == 1 {
		return fmt.Sprintf("%s ILIKE '%%%s%%'", fields[0], searchValue)
	}

	var conditions []string

	// Add conditions based on non-empty filter criteria
	for _, field := range fields {
		condition := fmt.Sprintf("%s %s '%%%s%%'", field, "ILIKE", searchValue)
		conditions = append(conditions, condition)
	}

	// Combine conditions with OR and wrap in parentheses
	conditionClause := "(" + strings.Join(conditions, " OR ") + ")"

	return conditionClause
}

func BuildOrderByClause(initial, sortField string) string {
	return fmt.Sprintf("ORDER BY %s ", fmt.Sprintf("%s.%s", initial, sortField))
}

func BuildLimitClause(perPage int) string {
	if perPage != 0 {
		return fmt.Sprintf("LIMIT %d ", perPage)
	}
	return ""
}

func BuildOffsetClause(page, perPage int) string {
	offset := (page - 1) * perPage
	return fmt.Sprintf("OFFSET %d ", offset)
}

func BuildQuery(pageRequest *model.PageRequest, initial string, searchFields []string) string {
	sqlQuery := ""

	// FILTER
	if pageRequest.Filters != nil {
		if filterCondition := BuildFilterCondition(pageRequest.Filters, initial); filterCondition != "" {
			sqlQuery += "AND " + filterCondition + " "
		}
	}

	// SEARCH
	if pageRequest.Search != "" && len(searchFields) > 0 {
		searchCondition := BuildSearchCondition(searchFields, pageRequest.Search)
		sqlQuery += "AND " + searchCondition + " "
	}

	// SORT
	if pageRequest.Sort != "" {
		orderByClause := BuildOrderByClause(initial, pageRequest.Sort)
		sqlQuery += orderByClause
	}

	// LIMIT
	limitClause := BuildLimitClause(pageRequest.PerPage)
	sqlQuery += limitClause

	// OFFSET
	offsetClause := BuildOffsetClause(pageRequest.Page, pageRequest.PerPage)
	sqlQuery += offsetClause

	return sqlQuery
}
