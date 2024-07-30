package logger

import (
	"fmt"

	"github.com/ramizkhan99/dio-be/internal/config"
	"go.uber.org/zap"
)

var logger *zap.Logger

func InitializeLogger() *zap.Logger {
	var err error
	env := config.GetEnv()
	switch env {
	case "DEV":
		logger, err = zap.NewDevelopment()
	case "PROD":
		logger, err = zap.NewProduction()
	default:
		panic(fmt.Errorf("failed to identify application environment: %s", env))
	}
	if err != nil {
		panic(fmt.Errorf("failed to initialize Zap logger: %w", err))
	}
	defer logger.Sync()
	logger.Sugar().Infof("Logger initialized for %s environment", env)
	return logger
}

func GetLogger() *zap.Logger {
	return logger
}

func GetSugaredLogger() *zap.SugaredLogger {
	return logger.Sugar()
}
