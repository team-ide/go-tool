package util

import (
	"bytes"
	"encoding/json"
)

func JSONDecodeUseNumber(bs []byte, obj any) (err error) {
	d := json.NewDecoder(bytes.NewReader(bs))
	d.UseNumber()
	err = d.Decode(obj)
	return
}

func ObjToObjByJson(obj any, toObj any) (err error) {
	b, err := ObjToJsonBuffer(obj)
	if err != nil {
		return
	}
	err = JSONDecodeUseNumber(b.Bytes(), toObj)
	if err != nil {
		return
	}
	return
}

// ObjToJson 对象 转 json 字符串
// ObjToJson(obj)
func ObjToJson(obj any) (res string, err error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false) // 关闭 HTML 转义
	err = encoder.Encode(obj)
	if err != nil {
		return
	}
	res = buf.String()
	return
}

// ObjToJsonBuffer 对象 转 json Buffer
// ObjToJsonBuffer(obj)
func ObjToJsonBuffer(obj any) (buf bytes.Buffer, err error) {
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false) // 关闭 HTML 转义
	err = encoder.Encode(obj)
	if err != nil {
		return
	}
	return
}

// ObjToJsonBytes 对象 转 json Buffer
// ObjToJsonBytes(obj)
func ObjToJsonBytes(obj any) (bs []byte, err error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false) // 关闭 HTML 转义
	err = encoder.Encode(obj)
	if err != nil {
		return
	}
	bs = buf.Bytes()
	return
}

// JsonToMap json 字符串 转 map对象
// JsonToMap("{\"a\":1}")
func JsonToMap(str string) (res map[string]any, err error) {
	res = map[string]any{}
	err = JSONDecodeUseNumber([]byte(str), &res)
	return
}

// JsonToObj json 字符串 转 对象
// JsonToObj("{\"a\":1}", &obj)
func JsonToObj(str string, obj any) (err error) {
	err = JSONDecodeUseNumber([]byte(str), obj)
	return
}
