package main

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"os"
	"strings"
	"testing"
)

func TestMod(t *testing.T) {
	var modPath = `D:\Code\linkdood\vrv-job\job-engine\go.mod`
	var modDir = os.Getenv("GOPATH")
	modDir = util.FormatPath(modDir)
	modDir = modDir + `/pkg/mod`
	err := outModTree(0, modPath, modDir)
	if err != nil {
		panic(err)
	}
}

func outModTree(leven int, modPath string, modDir string) (err error) {
	lines, err := util.ReadLine(modPath)
	if err != nil {
		err = errors.New("outModTree ReadLine mod path [" + modPath + "] error:" + err.Error())
		return
	}
	var isInRequire = false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if isInRequire {
			if strings.HasPrefix(line, ")") {
				isInRequire = false
				break
			}
		} else {
			if strings.HasPrefix(line, "require (") {
				isInRequire = true
				continue
			}
		}
		if !isInRequire {
			continue
		}
		ss := strings.Split(line, " ")
		if len(ss) != 2 {
			continue
		}

		for i := 0; i < leven; i++ {
			fmt.Printf("\t")
		}
		fmt.Printf("%s@%s\n", ss[0], ss[1])
		subModPath := modDir + "/" + ss[0] + "@" + ss[1] + "/go.mod"
		exi, _ := util.PathExists(subModPath)
		if exi {
			err = outModTree(leven+1, subModPath, modDir)
			if err != nil {
				return
			}
		}
	}
	return
}
