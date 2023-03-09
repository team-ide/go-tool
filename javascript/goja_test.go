package javascript

import (
	"fmt"
	"testing"
)

func TestScript(t *testing.T) {
	script := "1 + util.UUID()"
	context := GetContext()

	res, err := Run(script, context)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
