package api

import (
	"encoding/json"
	"fmt"
	"github.com/jabong/florest-core/src/common/constants"
	validator "github.com/jabong/florest-core/src/common/validator"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
	"io/ioutil"
)

type sample struct {
}

type TestData struct {
	User                     User     `json: "User"`
	Address                  Address  `json: "Address"`
	RequestValidationMessage []string `json: "RequestValidationMessage"`
}

type User struct {
	FirstName      string `validate:"required"`
	LastName       string `validate:"required"`
	Age            uint16 `validate:"gte=0,lte=130"`
	Email          string `validate:"required,email"`
	FavouriteColor string `validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func (s sample) GetStructs(data orchestrator.WorkFlowData) validator.RequestStructs {
	var tstData *TestData
	tstData = new(TestData)
	bt, err := ioutil.ReadFile("testdata/requestValidationNode.json")

	if err != nil {
		fmt.Sprintf("Error loading Request Validation Node Data file %s \n %s", err)
	}
	err = json.Unmarshal(bt, tstData)
	if err != nil {
		fmt.Sprintf("Error Unmarshalling Request Validation Node Data %s \n %s", err)
	}
	var reqStructs *validator.RequestStructs = new(validator.RequestStructs)
	reqStructs.Body = tstData.Address
	reqStructs.Params = tstData.User
	return *reqStructs
}

type ReqVNAPI struct {
	v versionmanager.Version
}

func (a *ReqVNAPI) GetVersion() versionmanager.Version {
	return a.v
}

func (a *ReqVNAPI) SetVersion(action string, version string, resource string, path string) {
	a.v = versionmanager.Version{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: constants.OrchestratorBucketDefaultValue, //todo - should it be a constant
		Path:     path,
	}
}

func (a *ReqVNAPI) GetOrchestrator() orchestrator.Orchestrator {
	requestValidationOrchestrator := new(orchestrator.Orchestrator)
	requestValidationWorkflow := new(orchestrator.WorkFlowDefinition)
	requestValidationWorkflow.Create()
	//Creation of the nodes in the workflow definition
	s := new(sample)
	p := &(validator.RequestValidationExecutor{ValidateRequestInstance: s})
	p.SetID("request validation node")
	eerr := requestValidationWorkflow.AddExecutionNode(p)
	requestValidationWorkflow.SetStartNode(p)
	if eerr != nil {
		fmt.Println(eerr)
	}
	requestValidationOrchestrator.Create(requestValidationWorkflow)
	return *requestValidationOrchestrator
}

func (a *ReqVNAPI) GetHealthCheck() healthcheck.HCInterface {
	h := new(TestHealthCheck)
	h.S = "Request Validation Node"
	return h
}

func (a *ReqVNAPI) Init() {
	//api initialization should come here
}
