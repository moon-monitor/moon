package log

import (
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/moon-monitor/moon/pkg/config"
)

func New(isDev bool, cfg *config.Log) (*zap.SugaredLogger, error) {
	level, err := zap.ParseAtomicLevel(cfg.GetLevel())
	if err != nil {
		level = zap.NewAtomicLevel()
	}
	zapCfg := zap.Config{
		Level:             level,
		Development:       isDev,
		DisableCaller:     cfg.GetDisableCaller(),
		DisableStacktrace: cfg.GetDisableStacktrace(),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: cfg.GetFormat(),
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "ts",
			LevelKey:      "level",
			NameKey:       "ns",
			CallerKey:     "caller",
			FunctionKey:   zapcore.OmitKey,
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel: func() zapcore.LevelEncoder {
				if cfg.GetEnableColor() {
					return zapcore.CapitalColorLevelEncoder
				}
				return zapcore.LowercaseLevelEncoder
			}(),
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{cfg.GetOutput()},
		ErrorOutputPaths: []string{cfg.GetOutput()},
	}
	logger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

type sugaredLogger struct {
	logger *zap.SugaredLogger
}

func (s *sugaredLogger) Log(level log.Level, keyvals ...any) error {
	switch level {
	case log.LevelDebug:
		s.logger.Debug(keyvals...)
	case log.LevelInfo:
		s.logger.Info(keyvals...)
	case log.LevelWarn:
		s.logger.Warn(keyvals...)
	case log.LevelError:
		s.logger.Error(keyvals...)
	case log.LevelFatal:
		s.logger.Fatal(keyvals...)
	default:
		s.logger.Info(keyvals...)
	}
	return nil
}

func WithSugaredLogger(logger *zap.SugaredLogger) log.Logger {
	return &sugaredLogger{logger: logger}
}
