package javascript

import (
	"fmt"
	"testing"
)

func TestScript(t *testing.T) {
	script := "1 + UUID()"
	context := NewContext()

	res, err := Run(script, context)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
