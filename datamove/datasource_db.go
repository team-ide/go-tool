package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
)

func NewDataSourceDb() *DataSourceDb {
	return &DataSourceDb{
		DataSourceBase:     &DataSourceBase{},
		DataSourceSqlParam: &DataSourceSqlParam{},
		DataSourceDbParam:  &DataSourceDbParam{},
	}
}

type DataSourceDbParam struct {
	SelectSql        string `json:"selectSql,omitempty"`
	ShouldSelectPage bool   `json:"shouldSelectPage,omitempty"`
}

type DataSourceDb struct {
	*dialect.ParamModel
	*DataSourceBase
	*DataSourceDbParam
	*DataSourceSqlParam

	Service db.IService
}

func (this_ *DataSourceDb) GetParam() *dialect.ParamModel {

	if this_.ParamModel == nil {
		this_.ParamModel = &dialect.ParamModel{}
	}
	return this_.ParamModel
}

func (this_ *DataSourceDb) Stop(progress *Progress) {

}

func (this_ *DataSourceDb) ReadStart(progress *Progress) (err error) {

	if this_.SelectSql == "" {
		this_.SelectSql, _, err = this_.Service.GetDialect().DataListSelectSql(this_.GetParam(), "", this_.TableName, this_.GetDialectColumnList(), nil, nil)
		if err != nil {
			return
		}
		this_.ShouldSelectPage = true
	}

	countSql, e := dialect.FormatCountSql(this_.SelectSql)
	if e == nil {
		t, _ := this_.Service.Count(countSql, nil)
		progress.DataTotal += t
	}

	if len(this_.ColumnList) == 0 {
		str := strings.TrimSpace(this_.SelectSql)
		if this_.ShouldSelectPage {
			str = this_.Service.GetDialect().PackPageSql(str, 1, 1)
		}

		rows, e := this_.Service.GetDb().Query(str)
		if e == nil {
			defer func() { _ = rows.Close() }()
			columnTypes, _ := rows.ColumnTypes()
			if columnTypes != nil {
				for _, columnType := range columnTypes {
					column := &Column{
						ColumnModel: &dialect.ColumnModel{},
					}
					column.ColumnName = columnType.Name()
					if precision, scale, ok := columnType.DecimalSize(); ok {
						column.ColumnPrecision = int(precision)
						column.ColumnScale = int(scale)
					}
					column.ColumnDataType = columnType.DatabaseTypeName()
					if length, ok := columnType.Length(); ok {
						column.ColumnLength = int(length)
					}
					if nullable, ok := columnType.Nullable(); ok {
						column.ColumnNotNull = !nullable
					}

					this_.ColumnList = append(this_.ColumnList, column)
				}
			}
		}
	}

	return
}

func (this_ *DataSourceDb) Read(progress *Progress, dataChan chan *Data) (err error) {

	pageSize := progress.BatchNumber

	pageNo := 1

	var list []map[string]interface{}
	var doQuery func() (err error)
	doQuery = func() (err error) {
		if progress.ShouldStop() {
			return
		}

		selectSql := strings.TrimSpace(this_.SelectSql)
		var hasPage = false
		if this_.ShouldSelectPage {
			selectSql = this_.Service.GetDialect().PackPageSql(selectSql, int(pageSize), pageNo)
			hasPage = true
		}
		util.Logger.Info("datasource db read",
			zap.Any("selectSql", selectSql),
		)
		list, err = this_.Service.QueryMap(selectSql, nil)
		if err != nil {
			util.Logger.Error("datasource db read error",
				zap.Error(err),
			)
			return
		}

		var lastData = &Data{
			DataType: DataTypeCols,
		}
		for _, data := range list {

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

		lastData.columnList = &this_.ColumnList
		dataChan <- lastData
		if hasPage && lastData.Total >= pageSize {
			pageNo++
			err = doQuery()
			if err != nil {
				return
			}
		}
		return
	}

	err = doQuery()

	return
}

func (this_ *DataSourceDb) ReadEnd(progress *Progress) (err error) {
	return
}

func (this_ *DataSourceDb) WriteStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceDb) Write(progress *Progress, data *Data) (err error) {

	if this_.FillColumn && data.columnList != nil {
		this_.fullColumnListByColumnList(progress, data.columnList)
	}

	var sqlList []string
	var paramList [][]interface{}
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
					ss, ps, _, _, e := this_.Service.GetDialect().DataListInsertSql(this_.GetParam(), "", this_.TableName, this_.GetDialectColumnList(), []map[string]interface{}{d})
					if e != nil {
						progress.WriteCount.AddError(1, e)
						if !progress.ErrorContinue {
							err = e
							return
						}
					} else {
						sqlList = append(sqlList, ss...)
						paramList = append(paramList, ps...)
					}
				}

			}

		}
		break
	case DataTypeSql:
		data.Total = int64(len(data.SqlList))
		if data.Total > 0 {
			for _, s := range data.SqlList {
				sqlList = append(sqlList, s)
				paramList = append(paramList, []interface{}{})
			}
		}
		break
	case DataTypeSqlAndParams:
		data.Total = int64(len(data.SqlAndParams))
		if data.Total > 0 {
			for _, s := range data.SqlAndParams {
				sqlList = append(sqlList, s.Sql)
				paramList = append(paramList, s.Params)
			}
		}
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
		return
	}
	if len(sqlList) > 0 {
		size := int64(len(sqlList))
		_, e := this_.Service.Execs(sqlList, paramList)
		if e != nil {
			progress.WriteCount.AddError(size, e)
			if !progress.ErrorContinue {
				err = e
				return
			}
		} else {
			progress.WriteCount.AddSuccess(size)
		}
	}
	return
}

func (this_ *DataSourceDb) WriteEnd(progress *Progress) (err error) {
	return
}
