package datamove

import (
	"encoding/json"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/javascript"
	"github.com/team-ide/go-tool/util"
	"strings"
)

func NewDataSourceBase(columnList []*Column) *DataSourceBase {
	return &DataSourceBase{
		ColumnList: columnList,
	}
}

type DataSourceBase struct {
	ColumnList    []*Column
	script        *javascript.Script
	ScriptContext map[string]interface{}
}

type Column struct {
	*dialect.ColumnModel
	Value       string `json:"value"`
	SubProperty bool   `json:"subProperty"`
}

func (this_ *DataSourceBase) GetColumnList() []*Column {

	return this_.ColumnList
}
func (this_ *DataSourceBase) GetColumnNames() []string {
	var columnNames []string
	list := this_.GetColumnList()
	for _, c := range list {
		columnNames = append(columnNames, c.ColumnName)
	}
	return columnNames
}

func (this_ *DataSourceBase) GetDialectColumnList() []*dialect.ColumnModel {
	var columns []*dialect.ColumnModel
	list := this_.GetColumnList()
	for _, c := range list {
		if c.ColumnModel != nil && c.ColumnName != "" {
			columns = append(columns, c.ColumnModel)
		}
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

func (this_ *DataSourceBase) SetScriptContextData(data map[string]interface{}) {
	this_.initScriptContext()
	if data != nil {
		for key, value := range data {
			_ = this_.script.Set(key, value)
		}
	}
}

func (this_ *DataSourceBase) SetScriptContext(key string, value interface{}) {
	this_.initScriptContext()
	_ = this_.script.Set(key, value)
}

func (this_ *DataSourceBase) GetValueByScript(script string) (interface{}, error) {
	this_.initScriptContext()
	if script == "" {
		return nil, nil
	}
	return this_.script.GetScriptValue(script)
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
	list := this_.GetColumnList()
	for index, c := range list {
		if c.ColumnName == "" {
			continue
		}
		var v interface{}
		if vSize > index {
			v = cols[index]
		}
		res = append(res, v)
	}
	return
}

func (this_ *DataSourceBase) DataToValues(progress *Progress, data map[string]interface{}) (res []interface{}, err error) {

	list := this_.GetColumnList()
	subData := map[string]map[string]interface{}{}
	for _, column := range list {
		v := data[column.ColumnName]
		if column.SubProperty {
			v = nil
			ss := strings.Split(column.ColumnName, ".")
			if len(ss) == 2 {
				if find, ok := subData[ss[0]]; ok {
					v = find[ss[1]]
				} else {
					if find, ok := data[ss[0]]; ok && find != nil {
						bs, _ := json.Marshal(find)
						sub := map[string]interface{}{}
						_ = json.Unmarshal(bs, &sub)
						subData[ss[0]] = sub
						v = sub[ss[1]]
					}

				}
			}
		}
		res = append(res, v)
	}
	return
}

func (this_ *DataSourceBase) ValuesToData(progress *Progress, cols []interface{}) (data map[string]interface{}, err error) {

	list := this_.GetColumnList()
	data = map[string]interface{}{}
	vSize := len(cols)
	for index, c := range list {
		if c.ColumnName == "" {
			continue
		}
		var v interface{}
		if vSize > index {
			v = cols[index]
		}
		data[c.ColumnName] = v
	}
	return
}
