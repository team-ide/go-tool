package db

import (
	"reflect"
	"sync"
)

func WarpTemplate[T any](t T, opts *TemplateOptions) (res *Template[T]) {
	if opts.structColumnsCache == nil {
		opts.structColumnsCache = make(map[reflect.Type]map[string]*FieldColumn)
	}
	if opts.structColumnsLowerCache == nil {
		opts.structColumnsLowerCache = make(map[reflect.Type]map[string]*FieldColumn)
	}
	res = &Template[T]{}
	res.TemplateOptions = opts
	res.t = t
	res.objType = reflect.TypeOf(t)

	res.objValueType = res.objType
	for res.objValueType.Kind() == reflect.Ptr {
		res.objValueType = res.objValueType.Elem()
	}
	return
}

type TemplateOptions struct {
	Service                 IService
	ColumnTagName           string `json:"columnTagName"`    // 字段 tag 注解名称 默认 `column:"xx"`
	UseJsonTagName          bool   `json:"useJsonTagName"`   // 如果 字段 tag 注解 未找到 使用 json 注解 默认为 false
	UseFieldName            bool   `json:"useFieldName"`     // 如果 以上都未配置 使用 字段名称 默认为 false
	StrictColumnName        bool   `json:"strictColumnName"` // 严格的字段大小写 默认 false `userId 将 匹配 Userid UserId USERID`
	structColumnsCache      map[reflect.Type]map[string]*FieldColumn
	structColumnsLowerCache map[reflect.Type]map[string]*FieldColumn
	structCacheLock         sync.Mutex
}

type Template[T any] struct {
	*TemplateOptions
	t                  T
	objType            reflect.Type
	objValueType       reflect.Type
	structColumns      map[string]*FieldColumn
	structColumnsLower map[string]*FieldColumn
}

func (this_ *Template[T]) SelectOne(table string, whereSql string, whereParam any) (res T, err error) {

	return
}
