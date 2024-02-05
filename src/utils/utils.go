package utils

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"modular-monolithic/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func GetID(r *http.Request) (ID string) {
	vars := mux.Vars(r)
	ID = vars["id"]

	return ID
}

func NullBoolToBool(nb sql.NullBool) bool {
	// Check if the value is valid
	if nb.Valid {
		return nb.Bool
	}

	// Decide on a default value when the value is NULL
	return false // You can choose any default value
}

func ParseStringToUUID(value string) (resp uuid.UUID) {
	resp, err := uuid.Parse(value)
	if err != nil {
		zap.S().Error(err)
		return
	}

	return
}

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
			filter, _ = validateAndReturnFilterMap(value)
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
		Filter:   filter,
	}
}

func validateAndReturnFilterMap(filter []string) ([]map[string]map[string]interface{}, error) {
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
