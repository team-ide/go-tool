package datamove

import (
	"github.com/team-ide/go-dialect/dialect"
)

type DataSourceBase struct {
	SkipNames  []string  `json:"skipNames"`
	ColumnList []*Column `json:"columnList"`
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
