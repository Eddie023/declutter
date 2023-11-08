package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	NoColor bool
	Level   zapcore.Level
}

func New(opts ...func(*Options)) *zap.SugaredLogger {
	o := Options{
		Level: zapcore.InfoLevel,
	}

	for _, opt := range opts {
		opt(&o)
	}

	level := zap.NewAtomicLevel()
	level.SetLevel(o.Level)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	if !o.NoColor {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.Lock(os.Stderr),
		level,
	))
	defer logger.Sync()

	return logger.Sugar()
}

// withLogLevel will set the default log level to Debug.
func WithLogLevel(l zapcore.Level) func(*Options) {
	return func(o *Options) {
		o.Level = l
	}
}

// withNoColor will show log outputs without color.
func WithNoColor() func(*Options) {
	return func(o *Options) {
		o.NoColor = true
	}
}
