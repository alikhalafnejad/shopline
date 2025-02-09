package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(debug bool) {
	var cfg zap.Config

	if debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

func Sync() {
	_ = Logger.Sync()
}
