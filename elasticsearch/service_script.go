package elasticsearch

import (
	"errors"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"reflect"
)

func NewServiceScript(config interface{}) (res map[string]interface{}, err error) {
	if config == nil {
		err = errors.New("config is null")
		util.Logger.Error("NewServiceScript error", zap.Error(err))
		return
	}
	var service IService
	if c1, ok := config.(*Config); ok {
		service, err = New(c1)
		if err != nil {
			util.Logger.Error("NewServiceScript error", zap.Error(err))
			return
		}
	} else {
		if c2, ok := config.(Config); ok {
			service, err = New(&c2)
			if err != nil {
				util.Logger.Error("NewServiceScript error", zap.Error(err))
				return
			}
		} else {
			var c3 = &Config{}
			err = util.ObjToObjByJson(config, c3)
			if err != nil {
				err = errors.New("config to config by json error:" + err.Error())
				util.Logger.Error("NewServiceScript error", zap.Error(err))
				return
			}
			service, err = New(c3)
			if err != nil {
				util.Logger.Error("NewServiceScript error", zap.Error(err))
				return
			}
		}
	}
	res = map[string]interface{}{}
	res["service"] = service

	t := reflect.TypeOf(service)
	v := reflect.ValueOf(service)
	size := t.NumMethod()
	for i := 0; i < size; i++ {
		method := t.Method(i)

		name := method.Name
		res[name] = v.Method(i).Interface()
		res[util.FirstToLower(name)] = v.Method(i).Interface()
	}

	return
}
