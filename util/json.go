package util

import (
	jsoniter "github.com/json-iterator/go"
	"io"
)

var JsonConfigUseNumber = jsoniter.Config{
	EscapeHTML: false,
	UseNumber:  true,
}.Froze()

func JSONDecodeUseNumber(bs []byte, obj any) (err error) {
	err = JsonConfigUseNumber.Unmarshal(bs, obj)
	return
}

func JSONDecode(bs []byte, obj any) (err error) {
	err = JsonConfigDefault.Unmarshal(bs, obj)
	return
}

func JSONDecodeByReader(reader io.Reader, obj any) (err error) {
	decoder := JsonConfigDefault.NewDecoder(reader)
	err = decoder.Decode(obj)
	return
}

func JSONEncoderByWriter(writer io.Writer, obj any) (err error) {
	encoder := JsonConfigDefault.NewEncoder(writer)
	err = encoder.Encode(obj)
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

// JsonConfigDefault the default API
var JsonConfigDefault = jsoniter.Config{
	EscapeHTML: false,
}.Froze()

// ObjToJsonBytes 对象 转 json Buffer
// ObjToJsonBytes(obj)
func ObjToJsonBytes(obj any) (bs []byte, err error) {
	bs, err = JsonConfigDefault.Marshal(obj)
	if err != nil {
		return
	}
	return
}

// ObjToMarshalIndent 对象 转 json Buffer
// ObjToMarshalIndent(obj)
func ObjToMarshalIndent(obj any, prefix string, indent string) (bs []byte, err error) {
	bs, err = JsonConfigDefault.MarshalIndent(obj, prefix, indent)
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

func JsonBytesToObj(bs []byte, obj any) (err error) {
	err = JSONDecode(bs, obj)
	return
}
