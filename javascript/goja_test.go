package javascript

import (
	"fmt"
	"testing"
)

func TestScript(t *testing.T) {
	script := "1 + getUUID()"
	context := NewContext()

	res, err := Run(script, context)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	script = `
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
