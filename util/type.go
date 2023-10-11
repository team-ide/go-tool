package util

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func GetValueByType(valueType reflect.Type, data interface{}) (res interface{}, err error) {

	//fmt.Println("valueType:", valueType)
	var isPtr = valueType.Kind() == reflect.Ptr
	var setValue reflect.Value
	if isPtr {
		setValue = reflect.New(valueType.Elem())
	} else {
		setValue = reflect.New(valueType)
	}

	res = setValue.Interface()
	if data == nil {
		if !isPtr {
			res = reflect.ValueOf(res).Elem().Interface()
		} else {
			res = nil
		}
		return
	}
	// 赋值
	switch res.(type) {
	case string:
		res = GetStringValue(data)
		break
	case *string:
		sV := GetStringValue(data)
		res = &sV
	case bool:
		sV := GetStringValue(data)
		if sV != "" {
			res, err = strconv.ParseBool(sV)
		}
		break
	case *bool:
		sV := GetStringValue(data)
		var b bool
		if sV != "" {
			b, err = strconv.ParseBool(sV)
			if err != nil {
				return
			}
		}
		res = &b
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
		sV := GetStringValue(data)
		if sV != "" {
			err = json.Unmarshal([]byte(sV), &res)
		}
		break
	default:
		sV := GetStringValue(data)
		if sV != "" {
			err = json.Unmarshal([]byte(sV), &res)
		}
		break
	}
	if !isPtr {
		res = reflect.ValueOf(res).Elem().Interface()
	}
	//fmt.Println("resultType:", reflect.TypeOf(res))
	return
}
