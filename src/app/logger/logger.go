package logger

import (
	"modular-monolithic/config"

	"git.motiolabs.com/library/motiolibs/mlogger"
)

var Log *mlogger.Logger

func InitLogger() *mlogger.Logger {
	log := mlogger.InitLogrus(config.Get().AppMode, config.Get().AppName)
	Log = log

	return log
}
