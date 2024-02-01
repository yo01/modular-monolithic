package app

import (
	"net/http"

	"modular-monolithic/config"

	"git.motiolabs.com/library/motiolibs/mcarrier"
	"git.motiolabs.com/library/motiolibs/mhttp"

	"github.com/gorilla/mux"
)

// AppConfig - Application config
type AppConfig struct {
	Router  *mux.Router
	Config  *config.Config
	Carrier *mcarrier.Carrier
}

// InitRouter - Create mux router
func InitRouter(cfg *config.Config) *mux.Router {
	router := mux.NewRouter()

	//handle route not found
	router.NotFoundHandler = http.HandlerFunc(mhttp.NotFoundHandler)

	//recovery handler
	router.Use(mhttp.RecoveryHandler)

	return router
}
