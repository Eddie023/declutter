package internal

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLogger(isDebugMode bool) *zap.SugaredLogger {
	atom := zap.NewAtomicLevel()

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = ""

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.Lock(os.Stderr),
		atom,
	))
	defer logger.Sync()

	if isDebugMode {
		atom.SetLevel(zap.DebugLevel)
	}

	return logger.Sugar()
}
