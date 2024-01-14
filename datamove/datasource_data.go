package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
)

func NewDataSourceData() *DataSourceData {
	return &DataSourceData{
		DataSourceBase: &DataSourceBase{},
	}
}

type DataSourceData struct {
	*DataSourceBase
	DataList []map[string]interface{}
}

func (this_ *DataSourceData) GetDataList() []map[string]interface{} {
	return this_.DataList
}

func (this_ *DataSourceData) Stop(progress *Progress) {

}

func (this_ *DataSourceData) ReadStart(progress *Progress) (err error) {

	size := int64(len(this_.GetDataList()))
	progress.DataTotal += size
	if size > 0 {
		err = this_.initColumnListByData(progress, this_.DataList[0])
		if err != nil {
			return
		}
	}

	return
}

func (this_ *DataSourceData) initColumnListByData(progress *Progress, data map[string]interface{}) (err error) {

	var titles []string
	if data != nil {
		for k := range data {
			titles = append(titles, k)
		}
	}
	if len(this_.ColumnList) == 0 {
		for _, title := range titles {
			column := &Column{
				ColumnModel: &dialect.ColumnModel{},
			}
			column.ColumnName = title
			this_.ColumnList = append(this_.ColumnList, column)
		}
	}
	if len(this_.ColumnList) != len(titles) {
		err = errors.New(fmt.Sprintf("字段长度[%d]头部长度[%d]，长度不一致", len(this_.ColumnList), len(titles)))
		return
	}

	return
}
func (this_ *DataSourceData) Read(progress *Progress, dataChan chan *Data) (err error) {

	pageSize := progress.BatchNumber

	var lastData = &Data{
		DataType: DataTypeCols,
	}
	for _, data := range this_.DataList {
		if progress.ShouldStop() {
			return
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

func (this_ *DataSourceData) ReadEnd(progress *Progress) (err error) {
	return
}

func (this_ *DataSourceData) WriteStart(progress *Progress) (err error) {
	this_.DataList = nil
	return
}

func (this_ *DataSourceData) Write(progress *Progress, data *Data) (err error) {

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
					this_.DataList = append(this_.DataList, d)
					progress.WriteCount.AddSuccess(1)
				}

			}

		}
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
		return
	}
	progress.WriteCount.AddSuccess(data.Total)
	return
}

func (this_ *DataSourceData) WriteEnd(progress *Progress) (err error) {
	return
}
