package javascript

import (
	"fmt"
	"github.com/team-ide/go-tool/javascript/context_map"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	context_map.AddModule(&context_map.ModuleInfo{
		Name:    "logger",
		Comment: "日志",
		FuncList: []*context_map.FuncInfo{
			{
				Name:    "debug",
				Comment: "输出 debug 级别日志",
				Func: func(args ...interface{}) {
					msg, fields := formatZapArgs(args...)
					util.Logger.Debug(fmt.Sprint(msg...), fields...)
				},
			},
			{
				Name:    "info",
				Comment: "输出 info 级别日志",
				Func: func(args ...interface{}) {
					msg, fields := formatZapArgs(args...)
					util.Logger.Info(fmt.Sprint(msg...), fields...)
				},
			},
			{
				Name:    "warn",
				Comment: "输出 warn 级别日志",
				Func: func(args ...interface{}) {
					msg, fields := formatZapArgs(args...)
					util.Logger.Warn(fmt.Sprint(msg...), fields...)
				},
			},
			{
				Name:    "error",
				Comment: "输出 error 级别日志",
				Func: func(args ...interface{}) {
					msg, fields := formatZapArgs(args...)
					util.Logger.Error(fmt.Sprint(msg...), fields...)
				},
			},
			{
				Name:    "any",
				Comment: "创建 zap 参数 zap.Any(key, value)",
				Func: func(key, value interface{}) zap.Field {
					return zap.Any(util.GetStringValue(key), value)
				},
			},
		},
	})
}

func formatZapArgs(args ...interface{}) (msg []interface{}, fields []zap.Field) {
	for _, arg := range args {
		switch tV := arg.(type) {
		case zap.Field:
			fields = append(fields, tV)
		default:
			msg = append(msg, tV)
		}
	}

	return
}
