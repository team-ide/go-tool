package javascript

import (
	"fmt"
	"github.com/team-ide/go-tool/util"
	"testing"
)

func TestScript(t *testing.T) {
	var num = 9999999999999999
	fmt.Println(num)
	script := "1 + getUUID()"
	context := NewContext()

	res, err := Run(script, context)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	script = `
let num = 9999999999999999
console.log(getStringValue(num))

let dir = getRootDir()
console.log(dir)
let paths = loadDirFilenames(dir)
console.log(paths)
let filePath = dir+"/test.txt"
let bs = readFile(filePath)
console.log(bs.length + "---" + getStringValue(bs))
let content = readFileString(filePath)
return content
`
	res, err = RunScript(script, context)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func TestNumber(t *testing.T) {

	context := NewContext()
	script := `
4611686027042965191 + 11
`
	res, err := Run(script, context)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(util.GetStringValue(res))
	fmt.Println(util.StringToInt64(util.GetStringValue(res)))
}
