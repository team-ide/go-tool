package main

import (
	"fmt"
	"github.com/team-ide/go-tool/javascript"
	_ "github.com/team-ide/go-tool/javascript/context_service"
	"testing"
)

func TestJavascriptService(t *testing.T) {

	context := javascript.NewContext()
	script := `
let config = {
	address:"",
	auth:"",
};
let redisParam = newRedisParam()
redisParam.Database = 0
let redisService = newRedisService(config);
res = redisService.Get(redisParam, "online-statistics-key")
return res
`
	res, err := javascript.RunScript(script, context)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
