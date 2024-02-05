package logger

import (
	"modular-monolithic/config"

	"git.motiolabs.com/library/motiolibs/mlogger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *mlogger.Logger

func InitLogger() *mlogger.Logger {
	log := mlogger.InitLogrus(config.Get().AppMode, config.Get().AppName)
	Log = log

	// Init zap log
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)

	return log
}
