package main

import (
	"fmt"

	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/components/cache"
	"github.com/jabong/florest-core/src/components/mongodb"
	"github.com/jabong/florest-core/src/components/sqldb"
	"github.com/jabong/florest-core/src/core/service"
	appConf "github.com/jabong/florest-core/src/examples/config"
	expConf "github.com/jabong/florest-core/src/examples/config"
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
	registerCustomInitFunc()
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

func registerCustomInitFunc() {
	service.RegisterCustomAPIInitFunc(func() {
		appConfig, err := expConf.GetExampleAppConfig()
		if err != nil {
			logger.Error(err)
			return
		}
		// initialize mongo
		if err = mongodb.Set("mymongo", appConfig.Mongo, new(mongodb.MongoDriver)); err != nil {
			logger.Error(err)
		}
		// initialize sqldb
		if err = sqldb.Set("mysdb", appConfig.MySQL, new(sqldb.MysqlDriver)); err != nil {
			logger.Error(err)
		}
		// initialize redis
		if err = cache.Set("myredis", appConfig.Cache.Redis, new(cache.RedisClientAdapter)); err != nil {
			logger.Error(err)
		}
		// initialize redis cluster
		if err = cache.Set("myRedisCluster", appConfig.Cache.RedisCluster, new(cache.RedisClientAdapter)); err != nil {
			logger.Error(err)
		}
	})
}

func registerErrors() {
}

func registerBuckets() {
}
