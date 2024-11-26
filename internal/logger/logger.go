package logger

import (
	"go.uber.org/zap"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupZapLogger(env string) *zap.Logger {
	var config zap.Config // Переменная для хранения конфигурации логгера

	switch env {
	case envLocal:
		config = zap.NewDevelopmentConfig()
	case envDev:
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case envProd:
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	default: // If env config is invalid, set prod settings by default due to security
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	logger, err := config.Build() // Создание логгера на основе конфигурации
	if err != nil {
		panic(err)
	}

	logger.Sync() // Ensure all buffered logs are written

	return logger

}
