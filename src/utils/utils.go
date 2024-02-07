package utils

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func GetID(r *http.Request) (ID string) {
	vars := mux.Vars(r)
	ID = vars["id"]

	return
}

func GetSubRouterName(r *http.Request) (subRouterName string) {
	vars := mux.CurrentRoute(r)
	subRouterName = vars.GetName()

	return
}

func ParseStringToUUID(value string) (resp uuid.UUID) {
	resp, err := uuid.Parse(value)
	if err != nil {
		zap.S().Error(err)
		return
	}

	return
}

func InStringArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
