package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

var (
	Logger *zap.Logger
)

func init() {

	// 默认 日志 输出在 控制台
	// 日志格式如下
	// [2023-05-17 11:17:33.063] [DEBUG] [util/logger.go:55] [logger test debug] [{"arg1": 1, "arg2": "2"}]
	// [2023-05-17 11:17:33.067] [INFO ] [util/logger.go:56] [logger test info] [{"arg1": 1, "arg2": "2"}]
	// [2023-05-17 11:17:33.067] [WARN ] [util/logger.go:57] [logger test warn] [{"arg1": 1, "arg2": "2"}]
	// [2023-05-17 11:17:33.067] [ERROR] [util/logger.go:58] [logger test error] [{"arg1": 1, "arg2": "2"}]

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "S",
		//FunctionKey:      "F",
		ConsoleSeparator: "] [",
		LineEnding:       "]\n",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("[" + t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(StrPadRight(strings.ToUpper(l.String()), 5, " "))
		},
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendFloat64(float64(d) / float64(time.Second))
		},
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(caller.TrimmedPath())
		},
		EncodeName: func(s string, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(strings.ToUpper(s))
		},
	}
	level := zapcore.DebugLevel
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(level),
	)
	Logger = zap.New(
		core,
		// 表示 输出 文件名 以及 行号
		zap.AddCaller(),
		// 表示 输出 堆栈跟踪 传入 level 表示 在哪个级别下输出
		zap.AddStacktrace(zapcore.ErrorLevel),
		//zap.AddCallerSkip(0),
	)
	//Logger.Debug("logger test debug", zap.Any("arg1", 1), zap.Any("arg2", "2"))
	//Logger.Info("logger test info", zap.Any("arg1", 1), zap.Any("arg2", "2"))
	//Logger.Warn("logger test warn", zap.Any("arg1", 1), zap.Any("arg2", "2"))
	//Logger.Error("logger test error", zap.Any("arg1", 1), zap.Any("arg2", "2"))
}

// GetLogger 获取logger输出对象
func GetLogger() *zap.Logger {

	return Logger
}
