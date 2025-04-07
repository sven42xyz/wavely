package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger(debug bool) {
	var logEncoding string

	level := zap.NewAtomicLevel()
	if debug {
		level.SetLevel(zap.DebugLevel)
		logEncoding = "console"
	} else {
		level.SetLevel(zap.InfoLevel)
		logEncoding = "json"
	}

	cfg := zap.Config{
		Level:            level,
		Encoding:         logEncoding,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	Log, err = cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("Fehler beim Initialisieren des Loggers: %v", err))
	}
}
