package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	log *zap.Logger

	LogOutput = "LOG_OUTPUT"
	LogLevel  = "LOG_LEVEL"
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutPutLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	err := log.Sync()
	if err != nil {
		return
	}
}

func Error(message string, err error, tags ...zap.Field) {
	log.Info(message, tags...)
	err = log.Sync()
	if err != nil {
		return
	}
}

func Debug(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	err = log.Sync()
	if err != nil {
		return
	}
}

func getOutPutLogs() string {
	otput := strings.ToLower(strings.TrimSpace(os.Getenv(LogOutput)))
	if otput == "" {
		return "stdout"
	}
	return otput
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LogLevel))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel

	}
}
