package main

import (
	"fmt"
	"github.com/team-ide/go-tool/javascript/context_map"
	"github.com/team-ide/go-tool/util"
	"os"
	"strings"
	"testing"
)

func TestGenVarUtil(t *testing.T) {
	var err error

	rootDir := util.GetRootDir()
	utilDir := rootDir + "/util/"
	filenames, err := util.LoadDirFilenames(utilDir)
	if err != nil {
		panic("LoadDirFilenames error:" + err.Error())
	}
	var funcInfoList []*context_map.FuncInfo
	for _, filename := range filenames {
		if strings.HasSuffix(filename, "_test.go") {
			continue
		}
		var lines []string
		lines, err = util.ReadLine(utilDir + filename)
		fmt.Println("---------------", filename, "---------------")
		for row, line := range lines {
			if !strings.HasPrefix(line, "func ") {
				continue
			}
			if row == 0 {
				continue
			}
			funcName := line[len("func "):strings.Index(line, "(")]
			var comment string
			var i = row - 1
			for {
				if !strings.HasPrefix(lines[i], "//") {
					break
				}
				str := lines[i]
				str = strings.TrimSpace(str[2:])
				if comment != "" {
					comment = str + "\n" + comment
				} else {
					comment = str
				}
				i--
			}
			vv := []rune(funcName)
			if len(vv) == 0 {
				continue
			}
			if vv[0] >= 97 && vv[0] <= 122 {
				continue
			}
			fmt.Println(comment)
			var fS = funcName + " "
			if !strings.HasPrefix(comment, fS) {
				continue
			}
			s := line[strings.Index(line, "(")+1:]
			argS := s[:strings.Index(s, ")")]
			var returnS = s[strings.Index(s, ")")+1:]
			if strings.Contains(s, "(") {
				returnS = s[strings.Index(s, "(")+1:]
			}
			argS = strings.TrimSpace(argS)
			returnS = strings.TrimSpace(returnS)
			returnS = strings.TrimRight(returnS, "{")
			returnS = strings.TrimSpace(returnS)
			returnS = strings.TrimRight(returnS, ")")
			returnS = strings.TrimSpace(returnS)

			funcInfo := &context_map.FuncInfo{
				Name:    funcName,
				Comment: comment[len(fS):],
			}
			argSS := strings.Split(argS, ",")
			var params []*context_map.FuncVarInfo
			for _, arg := range argSS {
				arg = strings.TrimSpace(arg)
				as := strings.Split(arg, " ")
				var name = strings.TrimSpace(as[0])
				var argType string
				if len(as) == 2 {
					argType = strings.TrimSpace(as[1])
				}
				if argType != "" {
					for _, param := range params {
						if param.Type == "" {
							param.Type = argType
						}
					}
				}
				funcInfo.Params = append(funcInfo.Params, &context_map.FuncVarInfo{
					Name: name,
					Type: argType,
				})
			}
			returnSS := strings.Split(returnS, ",")
			for _, re := range returnSS {
				re = strings.TrimSpace(re)
				as := strings.Split(re, " ")
				var argType = strings.TrimSpace(as[0])
				var name = "res"
				if len(as) == 2 {
					name = strings.TrimSpace(as[0])
					argType = strings.TrimSpace(as[1])
				}
				if strings.EqualFold(name, "err") || strings.EqualFold(argType, "error") {
					funcInfo.HasError = true
					continue
				}
				funcInfo.Return = &context_map.FuncVarInfo{
					Name: name,
					Type: argType,
				}
			}
			fmt.Println("funcName", funcName)
			fmt.Println("argS", argS)
			fmt.Println("returnS", returnS)
			fmt.Println("funcInfo:", util.GetStringValue(funcInfo))
			funcInfoList = append(funcInfoList, funcInfo)

		}
	}

	fmt.Println("--------------------", "func info list", "----------")

	genContent := `package builder_golang

import (
	"maker/parser_tm"
)

func (this_ *Context) initUtilVar() {
	utilSpace := this_.NewVarSpace()
	utilSpace.PackImpl = "github.com/team-ide/go-tool/util"
	utilSpace.PackAsName = "util"
	this_.AddVar("util", utilSpace)

`
	for _, funcInfo := range funcInfoList {
		//comment := funcInfo.Comment
		//comment = strings.ReplaceAll(comment, `"`, `\"`)
		name := funcInfo.Name
		//vv := []rune(name)
		//if vv[1] >= 97 && vv[1] <= 122 {
		//	name = util.FirstToLower(name)
		//}
		var argStr string
		for _, arg := range funcInfo.Params {
			argType := arg.Type
			if argType == "" || argType == "interface{}" {
				argType = "any"
			}
			argStr += `
			{Name: "` + arg.Name + `", Type: parser_tm.NewBindingTypeName("` + argType + `")},`
		}
		var returnStr string
		if funcInfo.Return != nil {
			var ret = `parser_tm.NewBindingTypeName("` + funcInfo.Return.Type + `")`
			if strings.HasPrefix(funcInfo.Return.Type, "[]") {
				ret = `parser_tm.NewBindingTypeList("` + strings.TrimPrefix(funcInfo.Return.Type, "[]") + `")`
			}

			returnStr += `
		Return: &parser_tm.FuncReturnNode{Name: "` + funcInfo.Return.Name + `", Type: ` + ret + `},`
		}
		if funcInfo.HasError {
			returnStr += `
		HasError: true,`
		}
		genContent += `
	this_.AddVar("` + name + `", &VarFunc{
		VarBase:    utilSpace.VarBase,
		ScriptName:  "` + name + `",
		Args: []*parser_tm.FuncArgNode{` + argStr + `
		},` + returnStr + `
	})
	utilSpace.AddVar("` + name + `", &VarFunc{
		ScriptName:  "` + name + `",
		Args: []*parser_tm.FuncArgNode{` + argStr + `
		},` + returnStr + `
	})
`
	}
	genContent += `
	return
}`

	f, err := os.Create(rootDir + "/builder_golang/var_util.go")
	if err != nil {
		panic("os.Create error:" + err.Error())
	}
	_, _ = f.WriteString(genContent)
}
