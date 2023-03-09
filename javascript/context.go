package javascript

import (
	"github.com/dop251/goja"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func NewContext() map[string]interface{} {
	baseContext := map[string]interface{}{}

	for _, funcInfo := range FuncList {
		baseContext[funcInfo.Name] = funcInfo.Func
	}

	return baseContext
}

type Script struct {
	dataContext map[string]interface{}
	vm          *goja.Runtime
}

func (this_ *Script) Set(name string, value interface{}) (err error) {
	this_.dataContext[name] = value
	err = this_.vm.Set(name, value)
	if err != nil {
		return
	}
	return
}
func (this_ *Script) GetScriptValue(script string) (value interface{}, err error) {
	if script == "" {
		value = ""
		return
	}

	var scriptValue goja.Value
	if scriptValue, err = this_.vm.RunString(script); err != nil {
		util.Logger.Error("表达式执行异常", zap.Any("script", script), zap.Error(err))
		return
	}
	value = scriptValue.Export()
	return
}

func (this_ *Script) GetStringScriptValue(script string) (value string, err error) {

	var scriptValue interface{}
	scriptValue, err = this_.GetScriptValue(script)
	if scriptValue != nil {
		value = util.GetStringValue(scriptValue)
		return
	}
	return
}
func NewScript() (script *Script, err error) {

	return NewScriptByParent(nil)
}
func NewScriptByParent(parent *Script) (script *Script, err error) {
	script = &Script{}
	script.vm = goja.New()
	script.dataContext = make(map[string]interface{})
	scriptContext := NewContext()
	for key, value := range scriptContext {
		err = script.vm.Set(key, value)
		if err != nil {
			return
		}
	}
	if parent != nil {
		for key, value := range parent.dataContext {
			err = script.Set(key, value)
			if err != nil {
				return
			}
		}
	}

	return
}
