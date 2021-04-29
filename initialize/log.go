package initialize

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setCore(filename string, level zapcore.Level) (core zapcore.Core) {
	// Encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// WriteSyncer
	lumberJackLogger := &lumberjack.Logger{
		Filename: filename,
		MaxSize:  10,
		Compress: false,
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)

	// LevelEnabler
	levelEnabler := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= level
	})

	// set core
	core = zapcore.NewCore(encoder, writeSyncer, levelEnabler)
	return
}

func loggerInit() {
	core := setCore("logs/xx.log", zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	zap.ReplaceGlobals(logger)
}
