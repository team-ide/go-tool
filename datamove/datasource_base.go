package datamove

import (
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"strings"
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
func (this_ *DataSourceBase) StringsToValues(progress *Progress, cols []string) (res []interface{}, err error) {
	vSize := len(cols)
	for index, _ := range this_.ColumnList {
		var v string
		if vSize > index {
			v = cols[index]
		}
		if progress.ReplaceSeparators != nil {
			for rK, rV := range progress.ReplaceSeparators {
				v = strings.ReplaceAll(v, rV, rK)
			}
		}
		if progress.ShouldTrimSpace {
			v = strings.TrimSpace(v)
		}
		res = append(res, v)
	}
	return
}

func (this_ *DataSourceBase) ValuesToValues(progress *Progress, cols []interface{}) (res []interface{}, err error) {
	vSize := len(cols)
	for index, _ := range this_.ColumnList {
		var v interface{}
		if vSize > index {
			v = cols[index]
		}
		if progress.ReplaceSeparators != nil {
			for rK, rV := range progress.ReplaceSeparators {
				if sV, sOk := v.(string); sOk {
					v = strings.ReplaceAll(sV, rV, rK)
				}
			}
		}
		if progress.ShouldTrimSpace {
			if sV, sOk := v.(string); sOk {
				v = strings.TrimSpace(sV)
			}
		}
		res = append(res, v)
	}
	return
}

func (this_ *DataSourceBase) ValuesToStrings(progress *Progress, cols []interface{}) (res []string, err error) {
	values, err := this_.ValuesToValues(progress, cols)
	if err != nil {
		return
	}
	for _, v := range values {
		res = append(res, util.GetStringValue(v))
	}
	return
}

func (this_ *DataSourceBase) DataToValues(progress *Progress, data map[string]interface{}) (res []interface{}, err error) {

	for _, column := range this_.ColumnList {
		v := data[column.ColumnName]
		if progress.ReplaceSeparators != nil {
			for rK, rV := range progress.ReplaceSeparators {
				if sV, sOk := v.(string); sOk {
					v = strings.ReplaceAll(sV, rV, rK)
				}
			}
		}
		if progress.ShouldTrimSpace {
			if sV, sOk := v.(string); sOk {
				v = strings.TrimSpace(sV)
			}
		}
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
		if progress.ReplaceSeparators != nil {
			for rK, rV := range progress.ReplaceSeparators {
				if sV, sOk := v.(string); sOk {
					v = strings.ReplaceAll(sV, rV, rK)
				}
			}
		}
		if progress.ShouldTrimSpace {
			if sV, sOk := v.(string); sOk {
				v = strings.TrimSpace(sV)
			}
		}
		data[column.ColumnName] = v
	}
	return
}
