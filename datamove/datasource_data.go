package datamove

import (
	"errors"
	"fmt"
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

func (this_ *DataSourceData) Read(progress *Progress, dataChan chan *Data) (err error) {

	pageSize := progress.BatchNumber

	var lastData = &Data{
		DataType: DataTypeCols,
	}
	for _, data := range this_.DataList {
		if progress.ShouldStop() {
			return
		}

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
			if lastData.Total >= pageSize {
				lastData.columnList = &this_.ColumnList
				dataChan <- lastData
				lastData = &Data{
					DataType: DataTypeCols,
				}
			}
		}
	}
	if lastData.Total > 0 {
		lastData.columnList = &this_.ColumnList
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
