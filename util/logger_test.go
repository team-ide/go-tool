package util

import "testing"

var (
	skipLogger = NewLoggerByCallerSkip(1)
)

func TestLogger(t *testing.T) {
	Logger.Debug("this is debug log")
	Logger.Info("this is info log")
	Logger.Warn("this is warn log")
	Logger.Error("this is error log")
	logA()
}

func logA() {
	skipLogger.Info("logA message")
	logB()
}

func logB() {
	skipLogger.Info("logB message")
}
