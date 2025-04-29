package util

import (
	"bytes"
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
)

func JSONDecodeUseNumber(bs []byte, obj any) (err error) {
	d := json.NewDecoder(bytes.NewReader(bs))
	d.UseNumber()
	err = d.Decode(obj)
	return
}

func JSONDecode(bs []byte, obj any) (err error) {
	err = json.Unmarshal(bs, obj)
	return
}

func ObjToObjByJson(obj any, toObj any) (err error) {
	bs, err := ObjToJsonBytes(obj)
	if err != nil {
		return
	}
	err = JSONDecode(bs, toObj)
	if err != nil {
		return
	}
	return
}

// ObjToJson 对象 转 json 字符串
// ObjToJson(obj)
func ObjToJson(obj any) (res string, err error) {
	bs, err := ObjToJsonBytes(obj)
	if err != nil {
		return
	}
	res = string(bs)
	return
}

// ObjToJsonBytes 对象 转 json Buffer
// ObjToJsonBytes(obj)
func ObjToJsonBytes(obj any) (bs []byte, err error) {
	var j = jsoniter.ConfigFastest
	bs, err = j.Marshal(obj)
	if err != nil {
		return
	}
	return
}

// JsonToMap json 字符串 转 map对象
// JsonToMap("{\"a\":1}")
func JsonToMap(str string) (res map[string]any, err error) {
	res = map[string]any{}
	err = JSONDecode([]byte(str), &res)
	return
}

// JsonToObj json 字符串 转 对象
// JsonToObj("{\"a\":1}", &obj)
func JsonToObj(str string, obj any) (err error) {
	err = JSONDecode([]byte(str), obj)
	return
}
