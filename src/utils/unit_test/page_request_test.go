package unit_test

import (
	"net/http"
	"testing"

	"modular-monolithic/model"
	"modular-monolithic/utils"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePaginationFromRequest(t *testing.T) {
	// Prepare a sample HTTP request with query parameters
	req, err := http.NewRequest("GET", "/example", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("per_page", "20")
	q.Add("page", "1")
	q.Add("paginate", "1")
	q.Add("sort", "name desc")
	q.Add("search", "example")
	q.Add("filter[]", "field1|operator1|value1")
	req.URL.RawQuery = q.Encode()

	// Call the function with the sample request
	pageRequest := utils.GeneratePaginationFromRequest(req)

	// Define the expected PageRequest struct
	expectedPageRequest := model.PageRequest{
		Paginate: 1,
		Page:     1,
		PerPage:  20,
		Sort:     "name desc",
		Search:   "example",
		Filters: []map[string]map[string]interface{}{
			{"field1": {"operator1": "value1"}},
		},
	}

	// Assert that the returned PageRequest matches the expected one
	assert.Equal(t, expectedPageRequest, pageRequest, "Generated PageRequest should match expected")
}

func TestCreateFilterFromMap(t *testing.T) {
	// Test case: create filter from map with string value
	key := "field1"
	operator := "eq"
	value := "value1"

	expectedFilter := map[string]map[string]interface{}{
		key: {
			operator: value,
		},
	}

	// Call the function with the test values
	filter := utils.CreateFilterFromMap(key, operator, value)

	// Assert that the returned filter matches the expected filter
	assert.Equal(t, expectedFilter, filter, "Generated filter should match expected")

	// Test case: create filter from map with integer value
	key = "field2"
	operator = "gt"
	intValue := 10

	expectedFilter = map[string]map[string]interface{}{
		key: {
			operator: intValue,
		},
	}

	// Call the function with the test values
	filter = utils.CreateFilterFromMap(key, operator, intValue)

	// Assert that the returned filter matches the expected filter
	assert.Equal(t, expectedFilter, filter, "Generated filter should match expected")
}

func TestGetSQLValue(t *testing.T) {
	var value interface{}

	// Test case: operator "="
	value = 10
	expected := "'10'"
	result := utils.GetSQLValue("=", value)
	assert.Equal(t, expected, result, "SQL value for operator '=' should be '10'")

	// Test case: operator ">"
	value = 20
	expected = "'20'"
	result = utils.GetSQLValue(">", value)
	assert.Equal(t, expected, result, "SQL value for operator '>' should be '20'")

	// Test case: operator "LIKE"
	value = "test"
	expected = "'%test%'"
	result = utils.GetSQLValue("LIKE", value)
	assert.Equal(t, expected, result, "SQL value for operator 'LIKE' should be '%test%'")

	// Test case: unknown operator
	value = "unknown"
	expected = ""
	result = utils.GetSQLValue("unknown", value)
	assert.Equal(t, expected, result, "SQL value for unknown operator should be empty string")
}

func TestBuildFilterCondition(t *testing.T) {
	// Test case: single filter condition with '=' operator
	filter := []map[string]map[string]interface{}{
		{"field1": {"=": "value1"}},
	}
	initial := "table"
	expected := "table.field1 = 'value1'"
	result := utils.BuildFilterCondition(filter, initial)
	assert.Equal(t, expected, result, "Generated filter condition should match expected")

	// Test case: multiple filter conditions with different operators
	filter = []map[string]map[string]interface{}{
		{"field2": {"=": 10}},
		{"field3": {">": 20}},
		{"field4": {"LIKE": "test"}},
	}
	initial = "table"
	expected = "table.field2 = '10' AND table.field3 > '20' AND table.field4 LIKE '%test%'"
	result = utils.BuildFilterCondition(filter, initial)
	assert.Equal(t, expected, result, "Generated filter conditions should match expected")
}

func TestBuildSearchCondition(t *testing.T) {
	// Test case: single field
	fields := []string{"field1"}
	searchValue := "value1"
	expected := "field1 ILIKE '%value1%'"
	result := utils.BuildSearchCondition(fields, searchValue)
	assert.Equal(t, expected, result, "Generated search condition should match expected")

	// Test case: multiple fields
	fields = []string{"field2", "field3", "field4"}
	searchValue = "test"
	expected = "(field2 ILIKE '%test%' OR field3 ILIKE '%test%' OR field4 ILIKE '%test%')"
	result = utils.BuildSearchCondition(fields, searchValue)
	assert.Equal(t, expected, result, "Generated search condition should match expected")
}

func TestBuildOrderByClause(t *testing.T) {
	// Test case: single sort field
	initial := "table"
	sortField := "field1"
	expected := "ORDER BY table.field1 "
	result := utils.BuildOrderByClause(initial, sortField)
	assert.Equal(t, expected, result, "Generated ORDER BY clause should match expected")

	// Test case: multiple sort fields
	initial = "schema"
	sortField = "field2"
	expected = "ORDER BY schema.field2 "
	result = utils.BuildOrderByClause(initial, sortField)
	assert.Equal(t, expected, result, "Generated ORDER BY clause should match expected")
}

func TestBuildLimitClause(t *testing.T) {
	// Test case: non-zero perPage
	perPage := 10
	expected := "LIMIT 10 "
	result := utils.BuildLimitClause(perPage)
	assert.Equal(t, expected, result, "Generated LIMIT clause should match expected")

	// Test case: zero perPage
	perPage = 0
	expected = ""
	result = utils.BuildLimitClause(perPage)
	assert.Equal(t, expected, result, "Generated LIMIT clause should be empty for zero perPage")
}

func TestBuildOffsetClause(t *testing.T) {
	// Test case: page 1, perPage 10
	page := 1
	perPage := 10
	expected := "OFFSET 0 "
	result := utils.BuildOffsetClause(page, perPage)
	assert.Equal(t, expected, result, "Generated OFFSET clause should match expected for page 1")

	// Test case: page 2, perPage 10
	page = 2
	perPage = 10
	expected = "OFFSET 10 "
	result = utils.BuildOffsetClause(page, perPage)
	assert.Equal(t, expected, result, "Generated OFFSET clause should match expected for page 2")
}

func TestBuildQuery(t *testing.T) {
	// Define sample page request
	pageRequest := &model.PageRequest{
		Paginate: 1,
		Page:     1,
		PerPage:  10,
		Sort:     "created_at asc",
		Search:   "example",
		Filters: []map[string]map[string]interface{}{
			{"field1": {"=": "value1"}},
			{"field2": {">": 20}},
		},
	}

	// Define sample initial value and search fields
	initial := "table"
	searchFields := []string{"field1", "field2"}

	// Call the function with the sample values
	sqlQuery := utils.BuildQuery(pageRequest, initial, searchFields)

	// Define expected SQL query based on the sample values
	expectedSQLQuery := "AND table.field1 = 'value1' AND table.field2 > '20' AND (field1 ILIKE '%example%' OR field2 ILIKE '%example%') ORDER BY table.created_at asc LIMIT 10 OFFSET 0 "

	// Assert that the generated SQL query matches the expected SQL query
	assert.Equal(t, expectedSQLQuery, sqlQuery, "Generated SQL query should match expected")
}
