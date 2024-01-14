package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/elasticsearch"
	"github.com/team-ide/go-tool/util"
	"strings"
)

func NewDataSourceEs() *DataSourceEs {
	return &DataSourceEs{
		DataSourceBase: &DataSourceBase{},
	}
}

type DataSourceEs struct {
	*DataSourceBase
	IndexName     string `json:"indexName"`
	IndexIdName   string `json:"indexIdName"`
	IndexIdScript string `json:"indexIdScript"`
	SelectSql     string `json:"selectSql"`

	Service elasticsearch.IService
}

func (this_ *DataSourceEs) Stop(progress *Progress) {

}

func (this_ *DataSourceEs) ReadStart(progress *Progress) (err error) {

	if this_.SelectSql == "" {

		var names = this_.GetColumnNames()
		this_.SelectSql += "SELECT "
		if len(names) == 0 {
			this_.SelectSql += "* "
		} else {
			this_.SelectSql += strings.Join(names, ",")
		}
		this_.SelectSql += "FROM " + this_.IndexName
	}

	countSql, e := dialect.FormatCountSql(this_.SelectSql)
	if e == nil {
		r, _ := this_.Service.QuerySql(countSql)
		if r != nil && len(r.Rows) > 0 && len(r.Rows[0]) > 0 {
			progress.DataTotal += util.StringToInt64(util.GetStringValue(r.Rows[0][0]))
		}
	}

	if len(this_.ColumnList) == 0 {
		pageSql := this_.SelectSql + " LIMIT 0"

		r, _ := this_.Service.QuerySql(pageSql)
		if r != nil {
			for _, columnType := range r.Columns {
				column := &Column{
					ColumnModel: &dialect.ColumnModel{},
				}
				column.ColumnName = columnType.Name
				column.ColumnDataType = columnType.Type
				this_.ColumnList = append(this_.ColumnList, column)
			}
		}
	}

	return
}

func (this_ *DataSourceEs) Read(progress *Progress, dataChan chan *Data) (err error) {

	pageSize := progress.BatchNumber
	var doQuery func() (err error)

	var scrollId string

	doQuery = func() (err error) {

		if progress.ShouldStop() {
			return
		}

		r, err := this_.Service.Scroll(this_.IndexName, scrollId, int(pageSize), nil, nil)
		if err != nil {
			return
		}
		scrollId = r.ScrollId

		var lastData = &Data{
			DataType: DataTypeCols,
		}
		for _, h := range r.Hits {

			data := map[string]interface{}{}

			e := util.JSONDecodeUseNumber([]byte(h.Source), &data)
			if e != nil {
				progress.WriteCount.AddError(1, e)
				if !progress.ErrorContinue {
					err = e
					return
				}
			} else {
				values, e := this_.DataToValues(progress, data)
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
				}
			}

		}

		dataChan <- lastData
		if lastData.Total >= pageSize {
			err = doQuery()
			if err != nil {
				return
			}
		}

		return
	}

	err = doQuery()
	if err != nil {
		return
	}

	return
}

func (this_ *DataSourceEs) ReadEnd(progress *Progress) (err error) {
	return
}

func (this_ *DataSourceEs) WriteStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceEs) Write(progress *Progress, data *Data) (err error) {

	switch data.DataType {
	case DataTypeCols:
		data.Total = int64(len(data.ColsList))
		if data.Total > 0 {
			for _, cols := range data.ColsList {
				d, e := this_.ValuesToData(progress, cols)
				if e != nil {
					progress.WriteCount.AddError(1, e)
					if !progress.ErrorContinue {
						err = e
						return
					}
				} else {
					var id string
					if this_.IndexIdScript != "" {
						this_.SetScriptContextData(d)
						id, e = this_.GetStringValueByScript(this_.IndexIdScript)
					} else {
						id = util.GetStringValue(d[this_.IndexIdName])
					}
					if e == nil {
						if id == "" {
							e = errors.New("id is empty")
						} else {
							_, e = this_.Service.InsertNotWait(this_.IndexName, id, d)
						}
					}
					if e != nil {
						progress.WriteCount.AddError(1, e)
						if !progress.ErrorContinue {
							err = e
							return
						}
					} else {
						progress.WriteCount.AddSuccess(1)
					}
				}

			}

		}
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
		return
	}
	return
}

func (this_ *DataSourceEs) WriteEnd(progress *Progress) (err error) {
	return
}
