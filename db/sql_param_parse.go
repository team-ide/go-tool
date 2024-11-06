package db

import (
	"errors"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"reflect"
	"regexp"
	"strings"
)

var (
	SqlParamParserDefaultParamBefore = "\\${"
	SqlParamParserDefaultParamAfter  = "}"
)

type SqlParamParser struct {
	appends      []string
	appendParams []any
	NotParseArg  bool // 不是 不解析参数值 只解析SQL
	ParamBefore  string
	ParamAfter   string
}

type SqlParamParserInfo struct {
	Sql        string
	Args       []any
	ParamNames []string
}

func NewSqlParamParser(sql string, param any) (res *SqlParamParser) {
	res = &SqlParamParser{}
	res.Append(sql, param)
	return
}

func (this_ *TemplateOptions) SqlParamParser(sql string, param any) (res *SqlParamParser) {
	res = NewSqlParamParser(sql, param)
	return
}

func (this_ *SqlParamParser) Append(sql string, param any) *SqlParamParser {
	this_.appends = append(this_.appends, sql)
	this_.appendParams = append(this_.appendParams, param)
	return this_
}

func (this_ *SqlParamParser) Parse() (res string, args []any, err error) {
	info, err := this_.ParseInfo()
	if err != nil {
		return
	}
	res = info.Sql
	args = info.Args
	return
}

func (this_ *SqlParamParser) ParseInfo() (info *SqlParamParserInfo, err error) {
	if this_.ParamBefore == "" {
		this_.ParamBefore = SqlParamParserDefaultParamBefore
	}
	if this_.ParamAfter == "" {
		this_.ParamAfter = SqlParamParserDefaultParamAfter
	}
	info = &SqlParamParserInfo{}
	var parseSql string
	var parseArgs []any
	for i := 0; i < len(this_.appends); i++ {
		sql := this_.appends[i]
		param := this_.appendParams[i]

		var paramValues []any
		var paramMap map[string]any
		if !this_.NotParseArg {
			paramMap = this_.GetParamData(param)
			if paramMap == nil {
				paramValues = []any{param}
			}
		}

		parseSql, parseArgs, err = this_.parse(info, sql, paramValues, paramMap)
		if err != nil {
			err = errors.New("SQL[" + sql + "]解析异常，" + err.Error())
			return
		}
		info.Sql += parseSql
		info.Args = append(info.Args, parseArgs...)

	}
	return
}

func (this_ *SqlParamParser) parse(info *SqlParamParserInfo, sql string, paramValues []any, paramMap map[string]any) (res string, args []any, err error) {
	if sql == "" {
		return
	}
	text := ""
	var re *regexp.Regexp
	expr := `(` + this_.ParamBefore + `)(.+?)(` + this_.ParamAfter + `)`
	re, _ = regexp.Compile(expr)
	indexList := re.FindAllIndex([]byte(sql), -1)
	var lastIndex = 0
	var paramIndex int
	for _, indexes := range indexList {
		text += sql[lastIndex:indexes[0]]

		lastIndex = indexes[1]

		paramStr := sql[indexes[0]:indexes[1]]
		subMatch := re.FindStringSubmatch(paramStr)

		paramName := subMatch[len(subMatch)-2]
		paramName = strings.TrimSpace(paramName)
		info.ParamNames = append(info.ParamNames, paramName)
		if this_.NotParseArg {
			args = append(args, nil)
		} else {
			arg, find := this_.parseArg(paramName, paramIndex, paramValues, paramMap)
			if !find {
				err = errors.New("参数[" + paramName + "]不存在")
				return
			}
			args = append(args, arg)
		}
		text += "?"
		paramIndex++
	}
	text += sql[lastIndex:]

	res = text
	return
}

func (this_ *SqlParamParser) parseArg(paramName string, paramIndex int, paramValues []any, paramMap map[string]any) (arg any, find bool) {
	if paramMap != nil {
		arg, find = paramMap[strings.ToLower(paramName)]
	} else {
		if paramIndex < len(paramValues)-1 {
			arg = paramValues[paramIndex]
			find = true
		}
	}
	return
}

func (this_ *SqlParamParser) GetParamData(param any) (data map[string]any) {
	if param == nil {
		return
	}
	data = map[string]any{}
	ok := this_.appendData(data, "", param)
	if !ok {
		data = nil
		return
	}
	return
}
func (this_ *SqlParamParser) appendData(data map[string]any, fieldName string, v any) (append bool) {
	if v == nil {
		return
	}
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
