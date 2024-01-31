package utils

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
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
