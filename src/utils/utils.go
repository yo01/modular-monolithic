package utils

import (
	"database/sql"
	"net/http"

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
