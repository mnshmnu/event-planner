package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(env string) (logger *zap.Logger, err error) {
	var zapConfig zap.Config

	// Set the logger config
	if env == "PROD" {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}

	// Set the encoder config
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err = zapConfig.Build()
	if err != nil {
		log.Fatalln("Error building logger", err)
		return nil, err
	}
	zap.ReplaceGlobals(logger)

	return logger, nil
}
