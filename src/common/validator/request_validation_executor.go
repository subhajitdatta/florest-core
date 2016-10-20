package service

import (
	"bytes"
	"github.com/jabong/florest-core/src/common/constants"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type RequestValidationExecutor struct {
	id                      string
	ValidateRequestInstance ValidateRequest
}

func (n RequestValidationExecutor) Execute(data workflow.WorkFlowData) (workflow.WorkFlowData, error) {

	validate = validator.New()

	var reqStructs RequestStructs
	appErrs := new(constants.AppErrors)
	reqStructs = n.ValidateRequestInstance.GetStructs(data)

	if &reqStructs != nil {
		if reqStructs.Body != nil {
			bodyVldErr := validate.Struct(reqStructs.Body)
			n.createAppErr(bodyVldErr, appErrs)
		}
		if reqStructs.Headers != nil {
			hdrVldErr := validate.Struct(reqStructs.Headers)
			n.createAppErr(hdrVldErr, appErrs)
		}

		if reqStructs.Params != nil {
			prmVldErr := validate.Struct(reqStructs.Params)
			n.createAppErr(prmVldErr, appErrs)
		}
	}
	return data, appErrs
}

func (n RequestValidationExecutor) createAppErr(err error, appErrs *constants.AppErrors) {
	if err == nil {
		return
	}
	var appErr = new(constants.AppError)
	for _, err := range err.(validator.ValidationErrors) {
		var buffer bytes.Buffer
		buffer.WriteString("Validation Failed for Field = " + err.StructNamespace())
		buffer.WriteString("\n")
		buffer.WriteString("Validation Condition = " + err.ActualTag())
		buffer.WriteString("\n")
		appErr.Message = buffer.String()
		appErr.Code = constants.RequestValidationFailedCode
		appErr.DeveloperMessage = ""
		appErrs.Errors = append(appErrs.Errors, *appErr)
	}
}

func (n RequestValidationExecutor) Name() string {
	return "Request Validation Executor"
}

func (n *RequestValidationExecutor) SetID(id string) {
	n.id = id
}

func (n RequestValidationExecutor) GetID() (id string, err error) {
	return n.id, nil
}
