package main

import (
	"fmt"
	"github.com/team-ide/go-tool/javascript"
	_ "github.com/team-ide/go-tool/javascript/context_service"
	"github.com/team-ide/go-tool/util"
	"testing"
)

func TestJavascriptService(t *testing.T) {

	context := javascript.NewContext()
	script := `
let config = {
	address:"127.0.0.1:6379",
	auth:"q7ZtCl^5S3",
};
start = getNowTime()
let redisService = newRedisService(config);
end = getNowTime()
console.log("new redisService use:",(end-start))

let key = "xx";

start = getNowTime()
res = redisService.get(key)
end = getNowTime()
console.log("redis get key:",key,",value:",res,",use:",(end-start))


start = getNowTime()
redisService.set(key,"这是一个UUID:"+getUUID())
end = getNowTime()
console.log("redis set key:",key,",use:",(end-start))

start = getNowTime()
res = redisService.get(key)
end = getNowTime()
console.log("redis get key:",key,",value:",res,",use:",(end-start))

config = {
	address:"127.0.0.1:2181",
};
start = getNowTime()
let zookeeperService = newZookeeperService(config);
end = getNowTime()
console.log("new zookeeperService use:",(end-start))


let path = "/xx";



start = getNowTime()
exists = zookeeperService.exists(path)
end = getNowTime()
console.log("zookeeper exists path:",path,",exists:",exists,",use:",(end-start))

if(!exists){
	start = getNowTime()
	zookeeperService.create(path, "这是一个UUID:"+getUUID())
	end = getNowTime()
	console.log("zookeeper create path:",path,",use:",(end-start))
}

start = getNowTime()
res = zookeeperService.get(path)
end = getNowTime()
console.log("zookeeper get path:",path,",value:",res,",use:",(end-start))

wait = newWaitGroup()
wait.Add(1)
zookeeperService.watchChildren(path, function(event){
	console.log("watchChildren path:",path,",event:", objToJson(event))
	return false;
})
console.log("zookeeper watchChildren path:",path,",use:",(end-start))
wait.Wait()

return res
`
	start := util.GetNowTime()
	res, err := javascript.RunScript(script, context)
	if err != nil {
		panic(err)
	}
	end := util.GetNowTime()

	fmt.Println("use:", end-start)
	fmt.Println(res)
}
