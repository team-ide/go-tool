package db

import (
	"github.com/team-ide/go-dialect/dialect"
	"reflect"
	"strings"
	"sync"
)

func WarpTemplate[T any](t T, opts *TemplateOptions) (res *Template[T]) {
	if opts.DialectParam == nil {
		opts.DialectParam = &dialect.ParamModel{}
	}
	if opts.Dialect == nil {
		if opts.Service != nil {
			opts.Dialect = opts.Service.GetDialect()
		} else {
			opts.Dialect, _ = dialect.NewDialect("mysql")
		}
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
	Dialect             dialect.Dialect
	DialectParam        *dialect.ParamModel
	Service             IService
	ColumnTagName       string `json:"columnTagName"`    // 字段 tag 注解名称 默认 `column:"xx"`
	UseJsonTagName      bool   `json:"useJsonTagName"`   // 如果 字段 tag 注解 未找到 使用 json 注解 默认为 false
	UseFieldName        bool   `json:"useFieldName"`     // 如果 以上都未配置 使用 字段名称 默认为 false
	StrictColumnName    bool   `json:"strictColumnName"` // 严格的字段大小写 默认 false `userId 将 匹配 Userid UserId USERID`
	structInfoCache     map[reflect.Type]*StructInfo
	structInfoCacheLock sync.Mutex
	StringEmptyUseNull  bool `json:"stringEmptyUseNull"`
	NumberZeroUseNull   bool `json:"numberZeroUseNull"`
}

type StructInfo struct {
	structColumns     []*FieldColumn
	structColumnMap   map[string]*FieldColumn
	structColumnLower map[string]*FieldColumn
}

type FieldColumn struct {
	Field      reflect.StructField
	Index      int
	ColumnName string
	IsString   bool
	IsNumber   bool
	IsBool     bool
}
type Template[T any] struct {
	*TemplateOptions
	t            T
	objType      reflect.Type
	objValueType reflect.Type
	structInfo   *StructInfo
}

func (this_ *Template[T]) SelectOne(table string, whereSql string, whereParam any) (res T, err error) {

	return
}

func (this_ *TemplateOptions) GetStructInfo(structType reflect.Type) (info *StructInfo) {
	for structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}
	this_.structInfoCacheLock.Lock()
	if this_.structInfoCache == nil {
		this_.structInfoCache = map[reflect.Type]*StructInfo{}
	}
	var ok bool
	info, ok = this_.structInfoCache[structType]
	if ok {
		this_.structInfoCacheLock.Unlock()
		return
	}
	defer this_.structInfoCacheLock.Unlock()
	info = &StructInfo{}
	this_.structInfoCache[structType] = info
	info.structColumnMap = map[string]*FieldColumn{}
	info.structColumnLower = map[string]*FieldColumn{}
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldColumn := &FieldColumn{
			Field: field,
			Index: i,
		}
		var str string
		var columnName string
		var tag = this_.ColumnTagName
		if tag == "" {
			tag = "column"
		}
		str = field.Tag.Get(tag)
		if str == "" && this_.UseJsonTagName {
			str = field.Tag.Get("json")
		}
		if str == "" && this_.UseFieldName {
			str = field.Name
		}
		if str != "" && str != "-" {
			ss := strings.Split(str, ",")
			columnName = ss[0]
		}
		if columnName != "" {
			fT := field.Type
			fTKind := fT.Kind()
			for fTKind == reflect.Ptr {
				fT = fT.Elem()
				fTKind = fT.Kind()
			}
			fieldColumn.IsString = fTKind == reflect.String
			fieldColumn.IsNumber = fTKind >= reflect.Int && fTKind <= reflect.Uint64
			fieldColumn.IsBool = fTKind == reflect.Bool
			fieldColumn.ColumnName = columnName
			info.structColumns = append(info.structColumns, fieldColumn)
			info.structColumnMap[columnName] = fieldColumn
			info.structColumnLower[strings.ToLower(columnName)] = fieldColumn
		}
	}
	return
}
