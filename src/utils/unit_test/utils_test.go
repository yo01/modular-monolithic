package unit_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"modular-monolithic/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetID(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/menus/{id}", func(w http.ResponseWriter, r *http.Request) {
		ID := utils.GetID(r)

		assert.Equal(t, "test-id", ID)
	})

	req, _ := http.NewRequest(http.MethodGet, "/menus/test-id", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
}

func TestGetSubRouterName(t *testing.T) {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/menus").Subrouter()
	subRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		subRouterName := utils.GetSubRouterName(r)

		assert.Equal(t, "menus", subRouterName)
	}).Name("menus")

	req, _ := http.NewRequest(http.MethodGet, "/menus/123", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
}

func TestParseStringToUUID(t *testing.T) {
	// Test case: valid UUID string
	validUUIDString := "c9bf9e57-1685-4c89-bafb-ff5af830be8a"
	expectedUUID, _ := uuid.Parse(validUUIDString)
	parsedUUID := utils.ParseStringToUUID(validUUIDString)
	assert.Equal(t, expectedUUID, parsedUUID, "Parsed UUID should match expected UUID")

	// Test case: invalid UUID string
	invalidUUIDString := "invalid-uuid"
	expectedNilUUID := uuid.Nil
	parsedNilUUID := utils.ParseStringToUUID(invalidUUIDString)
	assert.Equal(t, expectedNilUUID, parsedNilUUID, "Parsed UUID should be nil for invalid input string")
}

func TestInStringArray(t *testing.T) {
	// Test case: value exists in the array
	array := []string{"apple", "banana", "orange", "grape"}
	existingValue := "banana"
	assert.True(t, utils.InStringArray(existingValue, array), "Expected value to exist in the array")

	// Test case: value does not exist in the array
	nonExistingValue := "pineapple"
	assert.False(t, utils.InStringArray(nonExistingValue, array), "Expected value not to exist in the array")

	// Test case: empty array
	emptyArray := []string{}
	assert.False(t, utils.InStringArray(existingValue, emptyArray), "Expected value not to exist in an empty array")
}
