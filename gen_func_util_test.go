package main

import (
	"fmt"
	"github.com/team-ide/go-tool/javascript/context_map"
	"github.com/team-ide/go-tool/util"
	"os"
	"strings"
	"testing"
)

func TestGenFuncUtil(t *testing.T) {
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
			var commandLines []string
			var startComment string
			var comment string
			var i = row - 1
			for {
				if !strings.HasPrefix(lines[i], "//") {
					break
				}
				startComment = lines[i]
				startComment = strings.TrimSpace(startComment[2:])
				commandLines = append(commandLines, startComment)
				if comment != "" {
					comment = startComment + "\n" + comment
				} else {
					comment = startComment
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
			funcInfo := &context_map.FuncInfo{
				Name:    funcName,
				Comment: comment[len(fS):],
			}
			fmt.Println("funcName", funcName)
			funcInfoList = append(funcInfoList, funcInfo)

		}
	}

	fmt.Println("--------------------", "func info list", "----------")

	genContent := `package javascript

import (
	"github.com/team-ide/go-tool/javascript/context_map"
	"github.com/team-ide/go-tool/util"
)

func init() {
	context_map.AddModule(&context_map.ModuleInfo{
		Name:    "util",
		Comment: "工具模块",
		FuncList: []*context_map.FuncInfo{
`
	for _, funcInfo := range funcInfoList {
		comment := funcInfo.Comment
		comment = strings.ReplaceAll(comment, `"`, `\"`)
		name := funcInfo.Name
		vv := []rune(name)
		if vv[1] >= 97 && vv[1] <= 122 {
			name = util.FirstToLower(name)
		}
		genContent += `
		{
			Name:    "` + name + `",
			Comment: ` + "`" + comment + "`" + `,
			Func:    util.` + funcInfo.Name + `,
		},`
		fmt.Println(funcInfo.Name, ":", funcInfo.Comment)
	}
	genContent += `
		},
	})
}`

	f, err := os.Create(rootDir + "/javascript/func_util.go")
	if err != nil {
		panic("os.Create error:" + err.Error())
	}
	_, _ = f.WriteString(genContent)
}
