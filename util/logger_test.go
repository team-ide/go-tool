package util

import "testing"

func TestLogger(t *testing.T) {
	Logger.Debug("this is debug log")
	Logger.Info("this is info log")
	Logger.Warn("this is warn log")
	Logger.Error("this is error log")
}
