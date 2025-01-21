package configuration

import (
	"strings"

	"github.com/kingstonduy/go-core/logger"
	"github.com/kingstonduy/go-core/logger/logrus"
	rotateLog "github.com/kingstonduy/go-core/logger/logrus/hooks/rotate_log"
	"github.com/kingstonduy/go-core/trace"
	custom_logger "github.com/kingstonduy/order-service/internal/pkg/logger"
)

func GetLogger(cfg *Configuration, tracer trace.Tracer) logger.Logger {
	var (
		lConfig = cfg.LoggerConfig
	)
	level, err := logger.GetLevel(lConfig.LogLevel)
	if err != nil {
		level = logger.InfoLevel
	}

	hook, _ := rotateLog.NewRotateLogHook(
		rotateLog.WithRotateLogFilePattern("./access_log_%Y%m%d.log"),
	)

	if strings.ToUpper(cfg.LoggerConfig.LogLevel) == "DEBUG" {
		return custom_logger.NewcustomLogger(
			logger.WithLevel(level),
			logger.WithTracer(tracer),
			logrus.WithHooks(hook),
			logger.WithMaskedSensitiveData(false),
		)
	} else {
		return logrus.NewLogrusLogger(
			logger.WithLevel(level),
			logger.WithTracer(tracer),
			logrus.WithHooks(hook),
			logger.WithMaskedSensitiveData(false),
		)
	}

}
