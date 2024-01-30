package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetID(r *http.Request) (ID string) {
	vars := mux.Vars(r)
	ID = vars["id"]

	return ID
}
