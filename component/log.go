package component

import (
	"fmt"
	"time"

	conf "github.com/gedelumbung/go-movie/config"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func GetLogger(config *conf.Configuration) *log.Logger {
	if logger != nil {
		return logger
	}

	var f string = fmt.Sprintf("%s/error.%s.log", config.LOG.Dir, time.Now().Format("20060102"))
	logger = log.New()

	var env string = config.SITE.Env
	if env == "production" {
		logger.Level = log.ErrorLevel
	} else {
		logger.Level = log.DebugLevel
	}

	logger.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: f,
		log.InfoLevel:  f,
		log.WarnLevel:  f,
		log.ErrorLevel: f,
		log.FatalLevel: f,
		log.PanicLevel: f,
	}, &log.JSONFormatter{}))
	return logger
}
