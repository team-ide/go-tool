package main

import (
	"fmt"
	"github.com/team-ide/go-tool/javascript/context_map"
	"github.com/team-ide/go-tool/util"
	"os"
	"strings"
	"testing"
)

func TestGenContextService(t *testing.T) {
	var err error

	rootDir := util.GetRootDir()
	utilDir := rootDir + "/"
	filenames, err := util.LoadDirFilenames(utilDir)
	if err != nil {
		panic("LoadDirFilenames error:" + err.Error())
	}
	var serviceList []*context_map.ServiceInfo
	for _, filename := range filenames {
		if !strings.HasSuffix(filename, "iface.go") {
			continue
		}
		var lines []string
		lines, err = util.ReadLine(utilDir + filename)
		fmt.Println("---------------", filename, "---------------")
		serviceName := filename[0:strings.Index(filename, "/")]
		serviceName += "Service"
		serviceInfo := &context_map.ServiceInfo{
			Name: serviceName,
		}
		serviceList = append(serviceList, serviceInfo)
		var isInIFace bool
		for row, line := range lines {
			if strings.HasPrefix(line, "type IService interface {") {
				isInIFace = true
			}
			if !isInIFace {
				continue
			}
			if isInIFace {
				if strings.HasPrefix(line, "}") {
					break
				}
			}
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "// ") {
				continue
			}
			if strings.Index(line, "(") < 0 {
				continue
			}

			funcName := line[0:strings.Index(line, "(")]
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
			if len(vv) == 0 {
				continue
			}
			if vv[0] >= 97 && vv[0] <= 122 {
				continue
			}
			var fS = "// " + funcName + " "
			//fmt.Println("lastComment", lastComment, ",fS", fS)
			if !strings.HasPrefix(lastComment, fS) {
				continue
			}
			funcInfo := &context_map.FuncInfo{
				Name:    funcName,
				Comment: lastComment[len(fS):],
			}
			for i = len(commandLines) - 1; i >= 0; i-- {
				fmt.Println(commandLines[i])
			}
			fmt.Println(serviceName, ".", funcName)
			serviceInfo.FuncList = append(serviceInfo.FuncList, funcInfo)

		}
	}

	fmt.Println("--------------------", "service info list", "----------")

	genContent := `package context_service

import (
	"github.com/team-ide/go-tool/javascript/context_map"
)

func init() {
`
	for _, serviceInfo := range serviceList {
		comment := serviceInfo.Comment
		comment = strings.ReplaceAll(comment, `"`, `\"`)
		name := serviceInfo.Name
		vv := []rune(name)
		if vv[1] >= 97 && vv[1] <= 122 {
			name = util.FirstToLower(name)
		}
		genContent += `
	context_map.AddService(&context_map.ServiceInfo{
		Name:    "` + name + `",
		Comment: "` + comment + `",
		FuncList: []*context_map.FuncInfo{
`
		for _, funcInfo := range serviceInfo.FuncList {
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
				Comment: "` + comment + `",
			},`
			fmt.Println(serviceInfo.Name, ".", funcInfo.Name, ":", funcInfo.Comment)
		}

		genContent += `
		},
	})
`
	}
	genContent += `
}`

	f, err := os.Create(rootDir + "/javascript/context_service/service_func.go")
	if err != nil {
		panic("os.Create error:" + err.Error())
	}
	_, _ = f.WriteString(genContent)
}
