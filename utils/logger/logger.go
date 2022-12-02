package logger

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

const (
	// Values used in config files
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Error = "error"
)

func SetLevel(level string) {
	// Production configs
	// For dev config, using zap default
	rawJSON := []byte(`{
	  "level": "info",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
		"disableCaller": false,
		"disableStacktrace": false,
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	var err error
	err = json.Unmarshal(rawJSON, &cfg)

	var logger *zap.Logger
	if level == Debug {
		logger, err = zap.NewDevelopment()
	} else {
		switch level {
		case Info:
			cfg.Level.SetLevel(zap.InfoLevel)
		case Warn:
			cfg.Level.SetLevel(zap.WarnLevel)
		case Error:
			cfg.Level.SetLevel(zap.ErrorLevel)
		}
		logger = zap.Must(cfg.Build())
	}
	defer logger.Sync()

	if err != nil {
		log.Panic("Failed to construct logger: %w", err)
	}

	Log = logger.Sugar()
	Log.Info("Logger construction succeeded")
}
