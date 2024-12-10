package backend

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const defaultLoggerLevel = zap.InfoLevel

type LoggerConfiguration struct {
	LogLevel string `yaml:"log_level"`
	LogFile  string `yaml:"log_file"`
}

func NewLogger(lc LoggerConfiguration) *zap.SugaredLogger {
	lvl, err := zap.ParseAtomicLevel(lc.LogLevel)
	if err != nil {
		lvl = zap.NewAtomicLevelAt(defaultLoggerLevel)
	}

	var encoder zapcore.EncoderConfig
	var config zap.Config

	env := os.Getenv("APP_ENVIRONMENT")
	if env == DevelopmentEnvironment {
		encoder = zap.NewDevelopmentEncoderConfig()
		config = zap.NewDevelopmentConfig()
	} else {
		encoder = zap.NewProductionEncoderConfig()
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig = encoder
	config.Level = lvl
	config.OutputPaths = []string{"stdout", lc.LogFile}

	l := zap.Must(config.Build())
	return l.Sugar()
}
