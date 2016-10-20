package service

import (
	"fmt"

	"encoding/json"
	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/monitor"
	utilhttp "github.com/jabong/florest-core/src/common/utils/http"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

type HTTPResponseCreator struct {
	id string
}

func (n HTTPResponseCreator) Name() string {
	return "Http Response Creator"
}

func (n *HTTPResponseCreator) SetID(id string) {
	n.id = id
}

func (n HTTPResponseCreator) GetID() (id string, err error) {
	return n.id, nil
}

func (n *HTTPResponseCreator) Execute(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {

	rc, _ := data.ExecContext.Get(constants.RequestContext)
	logger.Info(fmt.Sprintln("entered ", n.Name()), rc)

	resStatus, _ := data.IOData.Get(constants.APPError)
	resData, _ := data.IOData.Get(constants.ResponseData)

	appError := new(constants.AppErrors)

	if resStatus != nil {
		if v, ok := resStatus.(*constants.AppError); ok {
			if v != nil { //if v is of type *AppError and is not nil
				appError.Errors = []constants.AppError{*v}
			}
		} else if v, ok := resStatus.(*constants.AppErrors); ok {
			if v != nil { //v is of type *AppErrors and is not nil
				appError = v
			}
		} else {
			appError.Errors = []constants.AppError{constants.AppError{Code: constants.InvalidErrorCode,
				Message: "Invalid App error"}}
		}
	}
	status := constants.GetAppHTTPError(*appError)
	debugData, _ := data.ExecContext.GetDebugMsg()

	resource, version, action, orchBucket, pathParams := getServiceVersion(data)

	serviceStatusKey := fmt.Sprintf("%v_%v_%v_%v_%v_%vHttp_%v", action,
		version, resource, orchBucket, pathParams, getCustomMetricPrefix(data), status.HTTPStatusCode)

	if status.HTTPStatusCode != constants.HTTPStatusSuccessCode {
		logger.Error(fmt.Sprintf("%s_%v Application Errors : %v", resource, status.HTTPStatusCode, appError), rc)
	}

	dderr := monitor.GetInstance().Count(serviceStatusKey, 1, nil, 1)
	if dderr != nil {
		logger.Error(fmt.Sprintln("Monitoring Error ", dderr.Error()), rc)
	}

	var appDebugData []utilhttp.Debug
	for _, d := range debugData {
		if v, ok := d.(workflow.WorkflowDebugDataInMemory); ok {
			appDebugData = append(appDebugData, utilhttp.Debug{Key: v.Key, Value: v.Value})
		}
	}

	m, _ := data.IOData.Get(constants.ResponseMetaData)
	md, _ := m.(*utilhttp.ResponseMetaData)
	appResponse := utilhttp.Response{Status: *status, Data: resData, DebugData: appDebugData, MetaData: md}
	data.IOData.Set(constants.Response, appResponse)
	jsonBody, err := json.Marshal(appResponse)
	if err != nil {
		return data, err
	}
	r, _ := data.IOData.Get(constants.APIResponse)
	apiResponse, _ := r.(utilhttp.APIResponse)
	apiResponse.HTTPStatus = appResponse.Status.HTTPStatusCode
	apiResponse.Body = jsonBody
	data.IOData.Set(constants.APIResponse, apiResponse)

	logger.Info(fmt.Sprintln("exiting ", n.Name()), rc)

	return data, nil
}
