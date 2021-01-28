package base

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"resk.micro/infra"
)

var (
	logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

type LogStarter struct {
	infra.BaseStarter
}

const (
	Console = "console"
	File    = "file"
)

var (
	Leavel = zap.InfoLevel
	Target = Console
)

func Logger() *zap.Logger {
	return logger
}

func (l *LogStarter) Init(ctx infra.StarterContext) {
	logger, _ = zap.NewProduction()
	logger.Sugar()

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/test.log",
		MaxSize:    1024,
		MaxBackups: 10,
		MaxAge:     1,
		Compress:   true,
	})

	var writeSyncer zapcore.WriteSyncer
	writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w)
	/*	if Target == Console {
			writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),w)
		}
		if Target == File {
			writeSyncer = zapcore.NewMultiWriteSyncer(w)
		}*/

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		writeSyncer,
		Leavel,
	)
	logger = zap.New(core, zap.AddCaller())
	Sugar = logger.Sugar()
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
