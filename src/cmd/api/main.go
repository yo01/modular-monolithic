package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"modular-monolithic/app"
	"modular-monolithic/app/database/sqlx"
	"modular-monolithic/app/logger"
	"modular-monolithic/config"
	"modular-monolithic/module"

	"git.motiolabs.com/library/motiolibs/mcarrier"

	"go.uber.org/zap"

	"github.com/rs/cors"
)

func main() {
	//Check config
	if err := config.CheckConfig(); err.Error != nil {
		panic(fmt.Sprintf("%s: %+v", err.Message, err.Error))
	}
	cfg := config.Get()

	// Init App Library
	sqlxDB := sqlx.InitPostgreConnection(cfg)
	logLib := logger.InitLogger()

	// Init Carrier
	carrier := mcarrier.New()
	carrier.Library = mcarrier.AppLibrary{
		Sqlx:   sqlxDB,
		Logger: logLib,
	}

	//Set App Routes
	router := app.InitRouter(cfg)

	//cors
	c := cors.New(cors.Options{
		AllowedMethods:   []string{"POST, GET, OPTIONS, PUT, DELETE"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept, Content-Type, Content-Length, Accept-Encoding, X-Api-Key, Authorization"},
	})

	// Set App Config
	appConfig := app.AppConfig{
		Router:  router,
		Config:  cfg,
		Carrier: &carrier,
	}

	// Add dependency injection
	module.Inject(appConfig)

	logLib.Info("Starting server...", nil)

	//serve server
	srv := &http.Server{
		Addr: fmt.Sprintf(
			"%s:%d",
			cfg.AppUrl,
			cfg.AppPort,
		),
		Handler: c.Handler(router),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	logLib.Info(fmt.Sprintf("starting application { %v } on port :%v", appConfig.Config.AppName, appConfig.Config.AppPort), nil)

	go listenAndServe(srv)
	waitForShutdown(srv)
}

func listenAndServe(apiServer *http.Server) {
	if err := apiServer.ListenAndServe(); err != nil {
		zap.S().Error(err)
		logger.Log.Fatal("unable to server", err)
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)

	<-sig

	logger.Log.Info("shutting down", nil)

	if err := apiServer.Shutdown(context.Background()); err != nil {
		zap.S().Error(err)
		logger.Log.Error("Error", err)
	}

	logger.Log.Info("shutdown complete", nil)
}
