package pkg

import (
	"fmt"
	"os"

	"github.com/irvanrifai/go-clean-architecture/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLog *zap.Logger

func InitLog() {
	appName := config.GetAppName()
	appEnv := config.GetEnv()
	logPath := fmt.Sprintf("%s.log", appName)
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("failed open log file: %v", err))
	}

	var level zapcore.Level
	if appEnv != "production" {
		level = zap.DebugLevel
	} else {
		level = zap.InfoLevel
	}

	var encoderConfig zapcore.EncoderConfig
	if appEnv != "production" {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	consoleDebugging := zapcore.Lock(os.Stdout)
	fileDebugging := zapcore.AddSync(f)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleDebugging, level),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileDebugging, level),
	)

	ZapLog = zap.New(core, zap.AddCaller())
}
