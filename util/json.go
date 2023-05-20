package util

import (
	"bytes"
	"encoding/json"
)

func JSONDecodeUseNumber(bs []byte, obj interface{}) (err error) {
	d := json.NewDecoder(bytes.NewReader(bs))
	d.UseNumber()
	err = d.Decode(obj)
	return
}

// ObjToJson 对象 转 json 字符串
// ObjToJson(obj)
func ObjToJson(obj interface{}) (res string, err error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return
	}
	res = string(bs)
	return
}

// JsonToMap json 字符串 转 map对象
// JsonToMap("{\"a\":1}")
func JsonToMap(str string) (res map[string]interface{}, err error) {
	res = map[string]interface{}{}
	err = JSONDecodeUseNumber([]byte(str), &res)
	return
}

// JsonToObj json 字符串 转 对象
// JsonToObj("{\"a\":1}", &obj)
func JsonToObj(str string, obj interface{}) (err error) {
	err = JSONDecodeUseNumber([]byte(str), obj)
	return
}
