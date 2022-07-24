package logger

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

func New() (*logr.Logger, error) {
	//config := zap.Config{
	//	Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
	//	Development: false,
	//	Encoding:    "json",
	//	//DisableCaller: true,
	//	EncoderConfig: zapcore.EncoderConfig{
	//		TimeKey:        "timestamp",
	//		LevelKey:       "level",
	//		NameKey:        "logger",
	//		MessageKey:     "message",
	//		StacktraceKey:  "stacktrace",
	//		LineEnding:     zapcore.DefaultLineEnding,
	//		EncodeLevel:    zapcore.LowercaseLevelEncoder,
	//		EncodeTime:     zapcore.EpochMillisTimeEncoder,
	//		EncodeDuration: zapcore.SecondsDurationEncoder,
	//	},
	//	OutputPaths:      []string{"stdout"},
	//	ErrorOutputPaths: []string{"stderr"},
	//}
	config := zap.NewProductionConfig()

	l, err := config.Build()
	if err != nil {
		return nil, err
	}

	zl := zapr.NewLogger(l)

	return &zl, nil
}
