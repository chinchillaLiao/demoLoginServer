package datatbases

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

var DebugLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      false,       // Don't include params in the SQL log
		Colorful:                  false,       // Disable color

	},
)
var DevLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Warn, // Log level
		IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      false,       // Don't include params in the SQL log
		Colorful:                  false,       // Disable color

	},
)
var ProdLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold:             time.Second,  // Slow SQL threshold
		LogLevel:                  logger.Error, // Log level
		IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      false,        // Don't include params in the SQL log
		Colorful:                  false,        // Disable color

	},
)

func EnvLogger() logger.Interface {
	ENVIRONMENT := os.Getenv("ENVIRONMENT")
	switch ENVIRONMENT {
	case "DEBUG":
		return DebugLogger
	case "DEV":
		return DevLogger
	case "PROD":
		return ProdLogger
	}
	return ProdLogger
}
