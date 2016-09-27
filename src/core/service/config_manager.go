package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/core/common/env"
)

const DefaultConfFile = "conf/conf.json"

type ConfigManager struct {
}

func (cm *ConfigManager) InitializeGlobalConfig(confFile string) {
	log.Println("\nInitializing Config ")
	cm.Initialize(confFile, config.GlobalAppConfig)
	log.Printf("\nGlobal Config=%+v", config.GlobalAppConfig)
	log.Printf("\nApplication Config=%+v", config.GlobalAppConfig.ApplicationConfig)
}

func (cm *ConfigManager) Initialize(filePath string, conf interface{}) {

	fmt.Println(fmt.Sprintf("config %+v", conf))
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error loading App Config file %s \n %s", filePath, err))
	}
	err = json.Unmarshal(file, conf)
	if err != nil {
		panic(fmt.Sprintf("Incorrect Json in %s \n %s", filePath, err))
	}
	log.Println("Application Config Created")
}

// UpdateConfigFromEnv updates provided config from environment variables
func (cm *ConfigManager) UpdateConfigFromEnv(conf interface{}, ty string) {
	if conf == nil {
		return
	}
	localConfigMap := make(map[string]string)
	if ty == "global" {
		if globalEnvUpdateMap == nil {
			return
		}
		localConfigMap = globalEnvUpdateMap
	} else {
		if configEnvUpdateMap == nil {
			return
		}
		localConfigMap = configEnvUpdateMap
	}

	configEnvUpdateValuesMap := make(map[string]string)
	for k, v := range localConfigMap {
		updatedVal, envValfound := env.GetOsEnviron().Get(v)

		if !envValfound {
			panic(fmt.Errorf("Environment variable %s not found", v))
		}

		if strings.Trim(updatedVal, " ") == "" {
			panic(fmt.Errorf("Environment variable %s is empty", v))
		}

		configEnvUpdateValuesMap[k] = updatedVal
	}

	byt, _ := json.Marshal(conf)

	newbyt, juerr := cm.updateJSONPath(configEnvUpdateValuesMap, byt, ".")
	if juerr != nil {
		panic(juerr)
	}

	if uerr := json.Unmarshal(newbyt, &conf); uerr != nil {
		panic(uerr)
	}
	if ty == "global" {
		log.Printf("Updated config from environment variables: %+v\n", config.GlobalAppConfig)
	}
}

func (cm *ConfigManager) updateJSONPath(queries map[string]string, byt []byte, pathSep string) (newByt []byte, err error) {
	unMarshallObj := make(map[string]interface{})
	jerr := json.Unmarshal(byt, &unMarshallObj)
	if jerr != nil {
		return byt, jerr
	}

	for query, newNodeVal := range queries {
		path := strings.Split(query, pathSep)
		var v map[string]interface{}

		jsPath := unMarshallObj
		for _, node := range path {

			nextJsPath, found := jsPath[node]
			if !found {
				return byt, fmt.Errorf("Not found node %s in path", node)
			}
			v = jsPath
			jsPath, _ = nextJsPath.(map[string]interface{})

		}

		leafNode := path[len(path)-1]

		var newNodeValConv interface{}
		var convErr error

		switch v[leafNode].(type) {
		case float64:
			newNodeValConv, convErr = strconv.ParseFloat(newNodeVal, 64)
		case int64:
			newNodeValConv, convErr = strconv.ParseInt(newNodeVal, 10, 64)
		case uint64:
			newNodeValConv, convErr = strconv.ParseUint(newNodeVal, 10, 64)
		case string:
			newNodeValConv, convErr = newNodeVal, nil
		case bool:
			newNodeValConv, convErr = strconv.ParseBool(newNodeVal)
		default:
			newNodeValConv, convErr = nil, errors.New("Unsupported json value for json xpath update")
		}
		if convErr != nil {
			return byt, convErr
		}
		v[leafNode] = newNodeValConv
	}
	return json.Marshal(unMarshallObj)
}
