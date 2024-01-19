package datamove

import (
	"errors"
	"fmt"
)

func NewDataSourceScript() *DataSourceScript {
	return &DataSourceScript{
		DataSourceBase:        &DataSourceBase{},
		DataSourceScriptParam: &DataSourceScriptParam{},
	}
}

type DataSourceScriptParam struct {
	Total int64 `json:"total"`
}

type DataSourceScript struct {
	*DataSourceBase
	*DataSourceScriptParam
}

func (this_ *DataSourceScript) Stop(progress *Progress) {

}

func (this_ *DataSourceScript) ReadStart(progress *Progress) (err error) {

	this_.SetScriptContext("total", this_.Total)
	this_.SetScriptContext("progress", progress)

	progress.DataTotal += this_.Total

	return
}

func (this_ *DataSourceScript) Read(progress *Progress, dataChan chan *Data) (err error) {

	pageSize := progress.BatchNumber

	var lastData = &Data{
		DataType: DataTypeCols,
	}
	for index := int64(0); index < this_.Total; index++ {
		if progress.ShouldStop() {
			return
		}
		values, e := this_.getValues(progress, index)
		if e != nil {
			progress.ReadCount.AddError(1, e)
			if !progress.ErrorContinue {
				err = e
				return
			}
		} else {
			lastData.ColsList = append(lastData.ColsList, values)
			lastData.Total++
			progress.ReadCount.AddSuccess(1)
			if lastData.Total >= pageSize {
				dataChan <- lastData
				lastData = &Data{
					DataType: DataTypeCols,
				}
			}
		}
	}
	if lastData.Total > 0 {
		dataChan <- lastData
	}

	return
}

func (this_ *DataSourceScript) getValues(progress *Progress, index int64) (values []interface{}, err error) {
	this_.SetScriptContext("index", index)
	this_.SetScriptContext("$index", index)

	var value interface{}
	for _, c := range this_.ColumnList {

		value, err = this_.GetStringValueByScript(c.Value)
		if err != nil {
			return
		}
		this_.SetScriptContext(c.ColumnName, value)

		values = append(values, value)
	}
	return
}

func (this_ *DataSourceScript) ReadEnd(progress *Progress) (err error) {
	return
}

func (this_ *DataSourceScript) WriteStart(progress *Progress) (err error) {
	return
}

func (this_ *DataSourceScript) Write(progress *Progress, data *Data) (err error) {

	err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
	return
}

func (this_ *DataSourceScript) WriteEnd(progress *Progress) (err error) {
	return
}
