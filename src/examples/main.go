package main

import (
	"fmt"

	"github.com/jabong/florest-core/src/core/service"
	appConf "github.com/jabong/florest-core/src/examples/config"
	"github.com/jabong/florest-core/src/examples/hellomongo"
	"github.com/jabong/florest-core/src/examples/hellomysql"
	"github.com/jabong/florest-core/src/examples/helloredis"
	"github.com/jabong/florest-core/src/examples/hellorediscluster"
	"github.com/jabong/florest-core/src/examples/helloworld"
)

//main is the entry point of the florest web service
func main() {
	fmt.Println("APPLICATION BEGIN")
	webserver := new(service.Webserver)
	registerConfig()
	registerErrors()
	registerBuckets()
	registerAllApis()
	webserver.Start()
}

func registerAllApis() {
	service.RegisterAPI(new(helloworld.HelloAPI))
	service.RegisterAPI(new(hellomongo.HelloMongoAPI))
	service.RegisterAPI(new(helloredis.HelloRedisAPI))
	service.RegisterAPI(new(hellomysql.HelloMySQLAPI))
	service.RegisterAPI(new(hellorediscluster.HelloRedisClusterAPI))
}

func registerConfig() {
	service.RegisterConfig(new(appConf.ExampleAppConfig))
}

func registerErrors() {
}

func registerBuckets() {
}
