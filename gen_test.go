package main

import (
	"fmt"
	"github.com/team-ide/go-tool/javascript"
	"github.com/team-ide/go-tool/util"
	"os"
	"strings"
	"testing"
)

func TestGen(t *testing.T) {
	var err error

	rootDir := util.GetRootDir()
	utilDir := rootDir + "/util/"
	filenames, err := util.LoadDirFilenames(utilDir)
	if err != nil {
		panic("LoadDirFilenames error:" + err.Error())
	}
	var funcInfoList []*javascript.FuncInfo
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
			var lastComment string
			var i = row - 1
			for {
				if !strings.HasPrefix(lines[i], "//") {
					break
				}
				lastComment = lines[i]
				commandLines = append(commandLines, lastComment)
				i--
			}
			vv := []rune(funcName)
			if vv[0] >= 97 && vv[0] <= 122 {
				continue
			}
			var fS = "// " + funcName + " "
			if !strings.HasPrefix(lastComment, fS) {
				continue
			}
			funcInfo := &javascript.FuncInfo{
				Name:    funcName,
				Comment: lastComment[len(fS):],
			}
			for i = len(commandLines) - 1; i >= 0; i-- {
				fmt.Println(commandLines[i])
			}
			fmt.Println("funcName", funcName)
			funcInfoList = append(funcInfoList, funcInfo)

		}
	}

	fmt.Println("--------------------", "func info list", "----------")

	genContent := `package javascript

import "github.com/team-ide/go-tool/util"

func init() {
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
	AddFunc(&FuncInfo{
		Name:    "` + name + `",
		Comment: "` + comment + `",
		Func:    util.` + funcInfo.Name + `,
	})
`
		fmt.Println(funcInfo.Name, ":", funcInfo.Comment)
	}
	genContent += `
}`

	f, err := os.Create(rootDir + "/javascript/func_util.go")
	if err != nil {
		panic("os.Create error:" + err.Error())
	}
	_, _ = f.WriteString(genContent)
}
