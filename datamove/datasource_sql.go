package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"io"
)

func NewDataSourceSql() *DataSourceSql {
	return &DataSourceSql{
		DataSourceBase: &DataSourceBase{},
		DataSourceFile: &DataSourceFile{},
	}
}

type DataSourceSql struct {
	*DataSourceBase
	*DataSourceFile
	*dialect.ParamModel
	DialectType string `json:"databaseType"`
	OwnerName   string `json:"ownerName"`
	TableName   string `json:"tableName"`
	dia_        dialect.Dialect

	sqlList []string
}

func (this_ *DataSourceSql) GetDialect() dialect.Dialect {
	if this_.dia_ == nil {
		if this_.DialectType == "" {
			this_.dia_, _ = dialect.NewDialect("mysql")
		} else {
			this_.dia_, _ = dialect.NewDialect(this_.DialectType)
		}
	}
	return this_.dia_
}
func (this_ *DataSourceSql) GetParam() *dialect.ParamModel {

	if this_.ParamModel == nil {
		this_.ParamModel = &dialect.ParamModel{}
	}
	return this_.ParamModel
}

func (this_ *DataSourceSql) ReadStart(progress *Progress) (err error) {
	if err != nil {
		return
	}
	err = this_.DataSourceFile.ReadStart(progress)
	if err != nil {
		return
	}

	// 读取 大概有多少SQL语句

	file, err := this_.GetReadFile()
	if err != nil {
		return
	}
	bs, err := io.ReadAll(file)
	if err != nil {
		return
	}

	this_.sqlList = this_.GetDialect().SqlSplit(string(bs))
	progress.DataTotal += int64(len(this_.sqlList))

	return
}
func (this_ *DataSourceSql) WriteStart(progress *Progress) (err error) {
	if err != nil {
		return
	}
	err = this_.DataSourceFile.WriteStart(progress)
	if err != nil {
		return
	}

	return
}

func (this_ *DataSourceSql) Read(progress *Progress, dataChan chan *Data) (err error) {
	pageSize := progress.BatchNumber

	var lastData = &Data{
		DataType: DataTypeSql,
	}
	for _, sqlInfo := range this_.sqlList {
		if progress.ShouldStop() {
			return
		}

		lastData.SqlList = append(lastData.SqlList, sqlInfo)
		lastData.Total++
		progress.ReadCount.AddSuccess(1)
		if lastData.Total >= pageSize {
			dataChan <- lastData
			lastData = &Data{
				DataType: DataTypeSql,
			}
		}
	}
	if lastData.Total > 0 {
		dataChan <- lastData
	}
	return
}

func (this_ *DataSourceSql) Write(progress *Progress, data *Data) (err error) {

	buf, err := this_.GetWriteBuf()
	if err != nil {
		return
	}
	var sqlList []string
	switch data.DataType {
	case DataTypeCols:
		data.Total = int64(len(data.ColsList))
		if data.Total > 0 {
			var dataList []map[string]interface{}
			for _, cols := range data.ColsList {
				d, e := this_.ValuesToData(progress, cols)
				if e != nil {
					progress.WriteCount.AddError(1, e)
					if !progress.ErrorContinue {
						err = e
						return
					}
				} else {
					dataList = append(dataList, d)
				}
			}
			param := this_.GetParam()
			param.AppendSqlValue = new(bool)
			*param.AppendSqlValue = true
			ss, _, _, _, _ := this_.GetDialect().DataListInsertSql(param, this_.OwnerName, this_.TableName, this_.GetDialectColumnList(), dataList)
			sqlList = append(sqlList, ss...)
		}
		break
	case DataTypeSql:
		data.Total = int64(len(data.SqlList))
		if data.Total > 0 {
			sqlList = data.SqlList
		}
		break
	default:
		err = errors.New(fmt.Sprintf("当前数据类型[%d]，不支持写入", data.DataType))
		return
	}
	for _, s := range sqlList {
		if s == "" {
			continue
		}
		progress.WriteCount.AddSuccess(1)
		bs := []byte(s + ";\n")
		n, e := buf.Write(bs)
		if e != nil {
			util.Logger.Error("sql write error", zap.Error(e))
		} else {
			size := len(bs)
			if n != size {
				fmt.Println(s)
				fmt.Println("n:", n, ",size:", size)
			}
		}
	}
	_ = buf.Flush()
	return
}
