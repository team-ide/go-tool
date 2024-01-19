package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/elasticsearch"
	"github.com/team-ide/go-tool/util"
)

func NewDataSourceEs() *DataSourceEs {
	return &DataSourceEs{
		DataSourceBase:    &DataSourceBase{},
		DataSourceEsParam: &DataSourceEsParam{},
	}
}

type DataSourceEsParam struct {
	IndexName     string `json:"indexName"`
	IndexIdName   string `json:"indexIdName"`
	IndexIdScript string `json:"indexIdScript"`
}

type DataSourceEs struct {
	*DataSourceBase
	*DataSourceEsParam

	Service elasticsearch.IService
}

func (this_ *DataSourceEs) Stop(progress *Progress) {

}

func (this_ *DataSourceEs) ReadStart(progress *Progress) (err error) {

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

			data["_id"] = h.Id
			data["_source"] = h.Source

			e := util.JSONDecodeUseNumber([]byte(h.Source), &data)
			if e != nil {
				progress.WriteCount.AddError(1, e)
				if !progress.ErrorContinue {
					err = e
					return
				}
			} else {
				if this_.FillColumn {
					this_.fullColumnListByData(progress, data)
				}
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
		lastData.columnList = &this_.ColumnList
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

	if this_.FillColumn && data.columnList != nil {
		this_.fullColumnListByColumnList(progress, data.columnList)
	}

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
						if this_.IndexIdName == "" {
							e = errors.New("index id name is empty")
						} else {
							id = util.GetStringValue(d[this_.IndexIdName])
						}
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
