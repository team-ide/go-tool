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
		name := filename[0:strings.Index(filename, "/")]
		serviceName := name + "Service"
		moduleName := name + "Module"
		serviceInfo := &context_map.ServiceInfo{
			Name:   serviceName,
			Module: moduleName,
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
			funcInfo := &context_map.FuncInfo{
				Name:    funcName,
				Comment: comment[len(fS):],
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
		name := serviceInfo.Name
		vv := []rune(name)
		if vv[1] >= 97 && vv[1] <= 122 {
			name = util.FirstToLower(name)
		}
		genContent += `
	` + serviceInfo.Module + `.Service = &context_map.ServiceInfo{
		Name:    "` + name + `",
		Comment: ` + "`" + comment + "`" + `,
		FuncList: []*context_map.FuncInfo{
`
		for _, funcInfo := range serviceInfo.FuncList {
			comment := funcInfo.Comment
			name := funcInfo.Name
			vv := []rune(name)
			if vv[1] >= 97 && vv[1] <= 122 {
				name = util.FirstToLower(name)
			}
			genContent += `
			{
				Name:    "` + name + `",
				Comment: ` + "`" + comment + "`" + `,
			},`
			fmt.Println(serviceInfo.Name, ".", funcInfo.Name, ":", funcInfo.Comment)
		}

		genContent += `
		},
	}
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
