package logging

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	cfg := zap.Config{
		OutputPaths:   []string{"./log/test.log", "stdout"},
		EncoderConfig: zap.NewProductionEncoderConfig(),
		Encoding:      "json",
		Level:         zap.NewAtomicLevelAt(zap.DebugLevel),
	}

	loggerFile, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer loggerFile.Sync()

	level := loggerFile.Level()

	loggerFile.Info("level", zap.String("level", level.String()))
	loggerFile.Error("Error")
	loggerFile.Info("failed to fetch URL",
		zap.String("url", "url"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
