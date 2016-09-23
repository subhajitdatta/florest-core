package config

import (
	"errors"
	"fmt"

	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/components/cache"
	"github.com/jabong/florest-core/src/components/mongodb"
	"github.com/jabong/florest-core/src/components/sqldb"
)

type ExampleAppConfig struct {
	ResponseHeaders config.ResponseHeaderFields `json:"ResponseHeaders,omitempty"`
	Mongo           *mongodb.MDBConfig          `json:"Mongo,omitempty"`
	Cache           *CacheConf                  `json:"Cache,omitempty"`
	MySQL           *sqldb.SDBConfig            `json:"MySQL,omitempty"`
}

type CacheConf struct {
	Redis        *cache.Config `json:"Redis,omitempty"`
	RedisCluster *cache.Config `json:"RedisCluster,omitempty"`
}

func GetExampleAppConfig() (*ExampleAppConfig, error) {
	c := config.GlobalAppConfig.ApplicationConfig
	appConfig, ok := c.(*ExampleAppConfig)
	if !ok {
		msg := fmt.Sprintf("Example APP Config Not correct %+v", c)
		logger.Error(msg)
		return nil, errors.New(msg)
	}
	return appConfig, nil
}
