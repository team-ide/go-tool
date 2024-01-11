package datamove

import (
	"errors"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/javascript"
	"github.com/team-ide/go-tool/util"
)

type DataSourceBase struct {
	SkipNames     []string  `json:"skipNames"`
	ColumnList    []*Column `json:"columnList"`
	script        *javascript.Script
	ScriptContext map[string]interface{}
}

type Column struct {
	*dialect.ColumnModel
	Value string `json:"value"`
}

func (this_ *DataSourceBase) GetColumnNames() []string {
	var columnNames []string
	for _, c := range this_.ColumnList {
		columnNames = append(columnNames, c.ColumnName)
	}
	return columnNames
}

func (this_ *DataSourceBase) GetDialectColumnList() []*dialect.ColumnModel {
	var columns []*dialect.ColumnModel
	for _, c := range this_.ColumnList {
		columns = append(columns, c.ColumnModel)
	}
	return columns
}

func (this_ *DataSourceBase) initScriptContext() {
	if this_.script == nil {
		this_.script, _ = javascript.NewScript()
		if this_.ScriptContext != nil {
			for k, v := range this_.ScriptContext {
				_ = this_.script.Set(k, v)
			}
		}
	}
}

func (this_ *DataSourceBase) SetScriptContext(key string, value interface{}) {
	this_.initScriptContext()
	_ = this_.script.Set(key, value)
}

func (this_ *DataSourceBase) GetValueByScript(script string) (res interface{}, err error) {
	this_.initScriptContext()
	if script == "" {
		return
	}
	res, err = this_.script.GetScriptValue(script)
	if err != nil {
		err = errors.New("script [" + script + "] error:" + err.Error())
		return
	}
	return
}

func (this_ *DataSourceBase) GetStringValueByScript(script string) (res string, err error) {
	r, err := this_.GetValueByScript(script)
	if err != nil {
		return
	}
	res = util.GetStringValue(r)
	return
}

func (this_ *DataSourceBase) ValuesToValues(progress *Progress, cols []interface{}) (res []interface{}, err error) {
	vSize := len(cols)
	for index, _ := range this_.ColumnList {
		var v interface{}
		if vSize > index {
			v = cols[index]
		}
		res = append(res, v)
	}
	return
}

func (this_ *DataSourceBase) DataToValues(progress *Progress, data map[string]interface{}) (res []interface{}, err error) {

	for _, column := range this_.ColumnList {
		v := data[column.ColumnName]
		res = append(res, v)
	}
	return
}

func (this_ *DataSourceBase) ValuesToData(progress *Progress, cols []interface{}) (data map[string]interface{}, err error) {

	data = map[string]interface{}{}
	vSize := len(cols)
	for index, column := range this_.ColumnList {
		var v interface{}
		if vSize > index {
			v = cols[index]
		}
		data[column.ColumnName] = v
	}
	return
}
