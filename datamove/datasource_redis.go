package datamove

import (
	"context"
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/redis"
	"github.com/team-ide/go-tool/util"
)

func NewDataSourceRedis() *DataSourceRedis {
	return &DataSourceRedis{
		DataSourceBase:       &DataSourceBase{},
		DataSourceRedisParam: &DataSourceRedisParam{},
	}
}

type DataSourceRedisParam struct {
	RedisDatabase      int    `json:"redisDatabase"`
	RedisKeyPattern    string `json:"redisKeyPattern"`
	RedisKeyName       string `json:"redisKeyName"`
	RedisKeyScript     string `json:"redisKeyScript"`
	RedisFieldName     string `json:"redisFieldName"`
	RedisFieldScript   string `json:"redisFieldScript"`
	RedisValueName     string `json:"redisValueName"`
	RedisValueTypeName string `json:"redisValueTypeName"`
	RedisValueByData   bool   `json:"redisValueByData"`
	RedisValueType     string `json:"redisValueType"`
}

type DataSourceRedis struct {
	*DataSourceBase
	*DataSourceRedisParam
	Service  redis.IService
	lastData *Data
}

func (this_ *DataSourceRedis) Stop(progress *Progress) {

}

func (this_ *DataSourceRedis) ReadStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceRedis) Read(progress *Progress, dataChan chan *Data) (err error) {

	this_.lastData = &Data{
		DataType: DataTypeCols,
	}
	var redisKeyPattern = this_.RedisKeyPattern
	if redisKeyPattern == "" {
		redisKeyPattern = "*"
	}
	param := &redis.Param{
		Ctx:      context.Background(),
		Database: this_.RedisDatabase,
	}

	res, err := this_.Service.Keys(redisKeyPattern, param)
	if err != nil {
		return
	}
	if res != nil && res.KeyList != nil {
		for _, keyInfo := range res.KeyList {
			if progress.ShouldStop() {
				return
			}
			err = this_.ReadKey(progress, dataChan, keyInfo, param)
			if err != nil {
				return
			}
		}
	}

	if this_.lastData != nil && this_.lastData.Total > 0 {
		this_.lastData.columnList = &this_.ColumnList
		dataChan <- this_.lastData
		this_.lastData = &Data{
			DataType: DataTypeCols,
		}
	}

	return
}
func (this_ *DataSourceRedis) ReadKey(progress *Progress, dataChan chan *Data, keyInfo *redis.KeyInfo, param *redis.Param) (err error) {

	var valuesList [][]interface{}
	defer func() {
		if err != nil {
			progress.ReadCount.AddError(1, err)
			if progress.ErrorContinue {
				err = nil
				return
			}
		} else {
			if len(valuesList) == 0 {
				return
			}
			this_.lastData.ColsList = append(this_.lastData.ColsList, valuesList...)
			this_.lastData.Total++
			progress.ReadCount.AddSuccess(1)
			if this_.lastData.Total >= progress.BatchNumber {
				this_.lastData.columnList = &this_.ColumnList
				dataChan <- this_.lastData
				this_.lastData = &Data{
					DataType: DataTypeCols,
				}
			}
		}

	}()
	param.Database = keyInfo.Database
	valueInfo, err := this_.Service.GetValueInfo(keyInfo.Key, param)
	if err != nil {
		return
	}
	if valueInfo.Value == nil {
		return
	}
	var values []interface{}

	if valueInfo.ValueType == "hash" {
		keyValue := valueInfo.Value.(map[string]string)
		for k, v := range keyValue {
			data := map[string]interface{}{}
			data[this_.RedisKeyName] = keyInfo.Key
			data[this_.RedisValueTypeName] = valueInfo.ValueType
			data[this_.RedisFieldName] = k
			data[this_.RedisValueName] = v
			if this_.RedisValueByData {
				_ = util.JSONDecodeUseNumber([]byte(v), &data)
			}
			if this_.FillColumn {
				this_.fullColumnListByData(progress, data)
			}
			values, err = this_.DataToValues(progress, data)
			if err != nil {
				return
			}
			valuesList = append(valuesList, values)
		}
	} else if valueInfo.ValueType == "list" || valueInfo.ValueType == "set" {
		vs := valueInfo.Value.([]string)
		for k, v := range vs {
			data := map[string]interface{}{}
			data[this_.RedisKeyName] = keyInfo.Key
			data[this_.RedisValueTypeName] = valueInfo.ValueType
			data[this_.RedisFieldName] = k
			data[this_.RedisValueName] = v
			if this_.RedisValueByData {
				_ = util.JSONDecodeUseNumber([]byte(v), &data)
			}
			if this_.FillColumn {
				this_.fullColumnListByData(progress, data)
			}
			values, err = this_.DataToValues(progress, data)
			if err != nil {
				return
			}
			valuesList = append(valuesList, values)
		}
	} else {
		data := map[string]interface{}{}
		data[this_.RedisKeyName] = keyInfo.Key
		data[this_.RedisValueTypeName] = valueInfo.ValueType
		data[this_.RedisValueName] = valueInfo.Value
		if this_.RedisValueByData {
			v, ok := valueInfo.Value.(string)
			if ok {
				_ = util.JSONDecodeUseNumber([]byte(v), &data)
			}
		}
		if this_.FillColumn {
			this_.fullColumnListByData(progress, data)
		}
		values, err = this_.DataToValues(progress, data)
		if err != nil {
			return
		}
		valuesList = append(valuesList, values)
	}

	return
}

func (this_ *DataSourceRedis) ReadEnd(progress *Progress) (err error) {
	return
}

func (this_ *DataSourceRedis) WriteStart(progress *Progress) (err error) {

	return
}

func (this_ *DataSourceRedis) Write(progress *Progress, data *Data) (err error) {

	if this_.FillColumn && data.columnList != nil {
		this_.fullColumnListByColumnList(progress, data.columnList)
	}
	if this_.FillColumn && data.columnList != nil {
		this_.fullColumnListByColumnList(progress, data.columnList)
	}

	param := &redis.Param{
		Ctx:      context.Background(),
		Database: this_.RedisDatabase,
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
					var key string
					var field string
					var value string
					if this_.RedisKeyScript != "" || this_.RedisFieldScript != "" {
						this_.SetScriptContextData(d)
					}
					if this_.RedisKeyScript != "" {
						key, e = this_.GetStringValueByScript(this_.RedisKeyScript)
					} else if this_.RedisKeyName != "" {
						key = util.GetStringValue(d[this_.RedisKeyName])
					}
					if e == nil && key == "" {
						e = errors.New("key is empty")
					}
					if e == nil {
						if this_.RedisValueType == "hash" {
							if this_.RedisFieldScript != "" {
								field, e = this_.GetStringValueByScript(this_.RedisFieldScript)
							} else if this_.RedisFieldName != "" {
								field = util.GetStringValue(d[this_.RedisFieldName])
							}
							if e == nil && field == "" {
								e = errors.New("hash field is empty")
							}
						}
					}
					if e == nil {
						if this_.RedisValueByData {
							value = util.GetStringValue(d)
						} else if this_.RedisValueName != "" {
							value = util.GetStringValue(d[this_.RedisValueName])
						}
					}
					if e == nil {
						if this_.RedisValueType == "hash" {
							e = this_.Service.HashSet(key, field, value, param)
						} else if this_.RedisValueType == "list" {
							e = this_.Service.ListPush(key, value, param)
						} else if this_.RedisValueType == "set" {
							e = this_.Service.SetAdd(key, value, param)
						} else if this_.RedisValueType == "string" {
							e = this_.Service.Set(key, value, param)
						} else {
							e = errors.New("不支持值类型[" + this_.RedisValueType + "]")
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

func (this_ *DataSourceRedis) WriteEnd(progress *Progress) (err error) {
	return
}
