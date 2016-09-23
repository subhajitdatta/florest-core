package service

import (
	"fmt"
	"strings"

	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	utilhttp "github.com/jabong/florest-core/src/common/utils/http"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

type URIInterpreter struct {
	id string
}

func (u URIInterpreter) Name() string {
	return "URL Interpreter"
}

func (u *URIInterpreter) SetID(id string) {
	u.id = id
}

func (u URIInterpreter) GetID() (id string, err error) {
	return u.id, nil
}

func (u URIInterpreter) Execute(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {
	rc, _ := data.ExecContext.Get(constants.RequestContext)

	logger.Info(fmt.Sprintln("Entered ", u.Name()), rc)

	resource, version, action, pathParams := u.getResource(data)
	data.IOData.Set(constants.Resource, resource)
	data.IOData.Set(constants.Version, version)
	data.IOData.Set(constants.Action, action)
	data.IOData.Set(constants.PathParams, pathParams)
	data.IOData.Set(constants.ResponseMetaData, utilhttp.NewResponseMetaData())

	logger.Info(fmt.Sprintln("exiting ", u.Name()), rc)
	return data, nil
}

func (u URIInterpreter) getResource(data workflow.WorkFlowData) (resource string,
	version string,
	action string, pathParams string) {

	rc, _ := data.ExecContext.Get(constants.RequestContext)
	uridata, _ := data.IOData.Get(constants.URI)
	actiondata, _ := data.IOData.Get(constants.HTTPVerb)

	var uri string
	if v, ok := uridata.(string); ok {
		uri = v
	}
	logger.Info(fmt.Sprintln("uri is ", uri), rc)

	uriArr := strings.Split(uri[1:], "/")

	if len(uriArr) >= 2 &&
		uriArr[0] == config.GlobalAppConfig.AppName &&
		strings.ToUpper(uriArr[1]) == constants.HealthCheckAPI {
		resource = constants.HealthCheckAPI
		version = ""
		pathParams = ""
	} else if len(uriArr) >= 3 && uriArr[0] == config.GlobalAppConfig.AppName {
		resource = strings.ToUpper(uriArr[2])
		version = strings.ToUpper(uriArr[1])
		if len(uriArr) > 3 {
			pathParams = strings.Join(uriArr[3:], "/")
		}
	} else {
		//Badly formed URI
		resource = ""
		version = ""
		pathParams = ""
	}

	if v, ok := actiondata.(utilhttp.Method); ok {
		action = string(v)
	}

	return resource, version, action, pathParams
}
