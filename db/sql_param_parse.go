package db

import (
	"errors"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"reflect"
	"regexp"
	"strings"
)

type SqlParamParser struct {
	*TemplateOptions
	sqlBuilder   string
	appends      []string
	appendParams []any
}

func (this_ *TemplateOptions) SqlParamParser(sql string, param any) (res *SqlParamParser) {
	res = &SqlParamParser{
		TemplateOptions: this_,
	}
	res.Append(sql, param)
	return
}

func (this_ *SqlParamParser) Append(sql string, param any) *SqlParamParser {
	this_.appends = append(this_.appends, sql)
	this_.appendParams = append(this_.appendParams, param)
	return this_
}

func (this_ *SqlParamParser) Parse() (res string, args []any, err error) {
	var parseSql string
	var parseArgs []any
	for i := 0; i < len(this_.appends); i++ {
		sql := this_.appends[i]
		param := this_.appendParams[i]

		var paramValues []any
		util.Logger.Debug("parse sql", zap.Any("sql", sql), zap.Any("param", param))
		paramMap := this_.GetParamData(param)
		if paramMap == nil {
			paramValues = []any{param}
		}

		parseSql, parseArgs, err = this_.parse(sql, paramValues, paramMap)
		if err != nil {
			err = errors.New("SQL[" + sql + "]解析异常，" + err.Error())
			return
		}
		res += parseSql
		args = append(args, parseArgs...)

	}
	return
}

func (this_ *SqlParamParser) parse(sql string, paramValues []any, paramMap map[string]any) (res string, args []any, err error) {
	if sql == "" {
		return
	}
	text := ""
	var re *regexp.Regexp
	re, _ = regexp.Compile(`[$]+{(.+?)}`)
	indexList := re.FindAllIndex([]byte(sql), -1)
	var lastIndex = 0
	var paramIndex int
	for _, indexes := range indexList {
		text += sql[lastIndex:indexes[0]]

		lastIndex = indexes[1]

		paramName := sql[indexes[0]+2 : indexes[1]-1]

		arg, find := this_.parseArg(paramName, paramIndex, paramValues, paramMap)
		if !find {
			err = errors.New("参数[" + paramName + "]不存在")
			return
		}
		args = append(args, arg)
		text += "?"
		paramIndex++
	}
	text += sql[lastIndex:]

	res = text
	return
}

func (this_ *SqlParamParser) parseArg(paramName string, paramIndex int, paramValues []any, paramMap map[string]any) (arg any, find bool) {
	if paramMap != nil {
		arg, find = paramMap[strings.ToLower(strings.TrimSpace(paramName))]
	} else {
		if paramIndex < len(paramValues)-1 {
			arg = paramValues[paramIndex]
			find = true
		}
	}
	return
}

func (this_ *SqlParamParser) GetParamData(param any) (data map[string]any) {
	data = map[string]any{}
	ok := this_.appendData(data, "", param)
	if !ok {
		data = nil
		return
	}
	return
}
func (this_ *SqlParamParser) appendData(data map[string]any, fieldName string, v any) (append bool) {

	objV := reflect.ValueOf(v)
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	if objV.Kind() == reflect.Map {
		this_.appendMapData(data, fieldName, objV)
		append = true
	} else if objV.Kind() == reflect.Struct {
		this_.appendStructData(data, fieldName, objV)
		append = true
	} else {
		if fieldName != "" {
			(data)[strings.ToLower(fieldName)] = v
			append = true
			util.Logger.Debug("parse sql append data add field", zap.Any("fieldName", fieldName), zap.Any("value", v))
		}
	}
	return
}

func (this_ *SqlParamParser) appendMapData(data map[string]any, parentName string, objV reflect.Value) {

	for _, kV := range objV.MapKeys() {
		if kV.Type().Kind() != reflect.String {
			continue
		}
		k := kV.String()
		vV := objV.MapIndex(kV)
		kName := k
		if parentName != "" {
			kName = parentName + "." + k
		}
		this_.appendData(data, kName, vV.Interface())
	}
	return
}

func (this_ *SqlParamParser) appendStructData(data map[string]any, parentName string, objV reflect.Value) {

	objT := objV.Type()
	for i := 0; i < objV.NumField(); i++ {
		vV := objV.Field(i)
		k := objT.Field(i).Name
		kName := k
		if parentName != "" {
			kName = parentName + "." + k
		}
		this_.appendData(data, kName, vV.Interface())
	}
	return
}
