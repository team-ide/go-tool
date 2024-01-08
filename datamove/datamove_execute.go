package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func (this_ *Executor) execute() (err error) {

	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			util.Logger.Error("execute panic error", zap.Error(err))
		}
	}()

	if this_.Source.IsDb() {
		if this_.Target.IsDb() {
			err = this_.dbToDb()
		} else if this_.Target.IsSql() {
			err = this_.dbToSql()
		} else if this_.Target.IsExcel() {
			err = this_.dbToExcel()
		} else if this_.Target.IsTxt() {
			err = this_.dbToTxt()
		} else {
			err = errors.New(fmt.Sprintf("不支持的 目标 类型[%s]", this_.Target.Type))
			util.Logger.Error("execute error", zap.Error(err))
			return
		}
	} else {
		err = errors.New(fmt.Sprintf("不支持的 源 类型[%s]", this_.Target.Type))
		util.Logger.Error("execute error", zap.Error(err))
		return
	}
	return
}
