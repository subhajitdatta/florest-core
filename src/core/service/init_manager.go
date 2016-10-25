package service

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/jabong/florest-core/src/common/config"
	"github.com/jabong/florest-core/src/common/constants"
	"github.com/jabong/florest-core/src/common/logger"
	"github.com/jabong/florest-core/src/common/monitor"
	"github.com/jabong/florest-core/src/common/profiler"
	"github.com/jabong/florest-core/src/common/utils/http"
	"github.com/jabong/florest-core/src/core/common/env"
	"github.com/jabong/florest-core/src/core/common/orchestrator"
	"github.com/jabong/florest-core/src/core/common/utils/healthcheck"
	"github.com/jabong/florest-core/src/core/common/utils/responseheaders"
	"github.com/jabong/florest-core/src/core/common/versionmanager"
)

type InitManager struct {
}

var ResourceToThreshold = map[string]int64{}

/*
Entry point for the init module.
Creates the following:
1> Application Configuration Object
2> Logger
3> Application Orchestrator pipelines
4> Data Access Objects
*/
func (im InitManager) Execute() {

	//Intialise OsEnvVariables
	env.GetOsEnviron()

	//Create the Global Application Config
	initConfig()

	//Sets the Application Performance Params
	setAppPerfParams()

	//Initialises Logger
	initLogger()

	// Initilalize Monitor
	InitMonitor()

	// initialize profiler
	initProfiler()

	//Create the WorkFlows
	InitVersionManager()

	// Initialise http pooling
	InitHTTPPool()

	//Initializes custom api init functionality
	InitCustomAPIInit()

	//Initialize Apis
	InitApis()

	//Initialise the Health Checks
	InitHealthCheck()
}

//initConfig initialises the Global Application Config
func initConfig() {
	cm := new(ConfigManager)
	cm.InitializeGlobalConfig(DefaultConfFile)
	cm.UpdateConfigFromEnv(config.GlobalAppConfig, "global")
}

//initLogger initialises the logger
func initLogger() {
	err := logger.Initialise(config.GlobalAppConfig.LogConfFile)
	if err != nil {
		panic(err)
	}
}

//initVersionManager create the WorkFlows
func InitVersionManager() {
	serviceParam := versionmanager.NewParam()
	serviceParam.Update("", createServiceOrchestrator(), nil)
	healthCheckParam := versionmanager.NewParam()
	healthCheckParam.Update("", createHealthCheckOrchestrator(), nil)
	vmap := versionmanager.VersionMap{
		versionmanager.BasicVersion{
			Resource: "SERVICE",
			Version:  "V1",
			Action:   "GET",
			BucketID: constants.OrchestratorBucketDefaultValue,
		}: serviceParam,
		versionmanager.BasicVersion{
			Resource: constants.HealthCheckAPI,
			Version:  "",
			Action:   "GET",
			BucketID: constants.OrchestratorBucketDefaultValue,
		}: healthCheckParam,
	}

	addAPIVersions(vmap)
	versionmanager.Initialize(vmap)
}

//Calls the custom api init function
func InitCustomAPIInit() {
	//If the apiCustomInitFunc is defined then execute it
	if apiCustomInitFunc != nil {
		apiCustomInitFunc()
	}
}

func addAPIVersions(vmap versionmanager.VersionMap) {
	for _, apiInstance := range apiList {
		version := apiInstance.GetVersion()
		param := vmap[version.GetBasicVersion()]
		if param == nil {
			param = versionmanager.NewParam()
			vmap[version.GetBasicVersion()] = param
		}
		rl := apiInstance.GetRateLimiter()
		err := param.Update(version.Path, apiInstance.GetOrchestrator(), &rl)
		if err != nil {
			logger.Error("Path - " + version.Path + " is not valid. Err : " + err.Error())
		}
	}
}

func createServiceOrchestrator() orchestrator.Orchestrator {
	logger.Info("Service Pipeline Creation begin")

	serviceOrchestrator := new(orchestrator.Orchestrator)
	serviceWorkflow := new(orchestrator.WorkFlowDefinition)
	serviceWorkflow.Create()

	//Create and add execution node UriInterpreter
	uriInterpreter := new(URIInterpreter)
	uriInterpreter.SetID("1")
	uerr := serviceWorkflow.AddExecutionNode(uriInterpreter)
	if uerr != nil {
		logger.Error(fmt.Sprintln(uerr))
	}

	//Create and add execution node BusinessLogicExecutor
	businessLogicExecutor := new(BusinessLogicExecutor)
	businessLogicExecutor.SetID("3")
	berr := serviceWorkflow.AddExecutionNode(businessLogicExecutor)
	if berr != nil {
		logger.Error(fmt.Sprintln(berr))
	}
	responseHeaderWriter := new(responseheaders.Writer)
	responseHeaderWriter.SetID("4")
	rhErr := serviceWorkflow.AddExecutionNode(responseHeaderWriter)
	if rhErr != nil {
		logger.Error(fmt.Sprintln(rhErr))
	}

	//Create and add execution node HTTPResponseCreator
	httpResponseCreator := new(HTTPResponseCreator)
	httpResponseCreator.SetID("5")
	herr := serviceWorkflow.AddExecutionNode(httpResponseCreator)
	if herr != nil {
		logger.Error(fmt.Sprintln(herr))
	}

	//Add the connections between the nodes
	c1err := serviceWorkflow.AddConnection(uriInterpreter, businessLogicExecutor)
	if c1err != nil {
		logger.Error(fmt.Sprintln(c1err))
	}
	c2err := serviceWorkflow.AddConnection(businessLogicExecutor, responseHeaderWriter)
	if c2err != nil {
		logger.Error(fmt.Sprintln(c2err))
	}
	c3err := serviceWorkflow.AddConnection(responseHeaderWriter, httpResponseCreator)
	if c3err != nil {
		logger.Error(fmt.Sprintln(c3err))
	}

	//Set the Workflow Start Node
	serviceWorkflow.SetStartNode(uriInterpreter)

	//Assign the workflow definition to the Orchestrator
	serviceOrchestrator.Create(serviceWorkflow)

	logger.Info(serviceOrchestrator.String())
	logger.Info("Service Pipeline Created")

	return *serviceOrchestrator
}

func createHealthCheckOrchestrator() orchestrator.Orchestrator {
	logger.Info("Health Check Pipeline Creation begin")

	healthCheckOrchestrator := new(orchestrator.Orchestrator)
	healthCheckOrchestratorWorkflow := new(orchestrator.WorkFlowDefinition)
	healthCheckOrchestratorWorkflow.Create()

	healthCheckExecutor := new(healthcheck.HCExecutor)
	healthCheckExecutor.SetID("1")
	hcerr := healthCheckOrchestratorWorkflow.AddExecutionNode(healthCheckExecutor)
	if hcerr != nil {
		logger.Error(fmt.Sprintln(hcerr))
	}

	healthCheckOrchestratorWorkflow.SetStartNode(healthCheckExecutor)
	healthCheckOrchestrator.Create(healthCheckOrchestratorWorkflow)

	logger.Info(healthCheckOrchestrator.String())
	logger.Info("Health Check Pipeline Created")

	return *healthCheckOrchestrator
}

//initApis initializes all apis
func InitApis() {
	for _, apiInstance := range apiList {
		apiInstance.Init()
	}
	logger.Info("Initialized apis")
}

//setAppPerfParams sets the Application's performance parameters
func setAppPerfParams() {
	perf := config.GlobalAppConfig.Performance
	setGCPercentage(perf.GCPercentage)
	setNoOfCPUCores(perf.UseCorePercentage)
}

//setGCPercentage sets when to trigger the garbage collection
func setGCPercentage(gcPer float64) {
	debug.SetGCPercent(int(gcPer))
}

//setNoOfCPUCores sets number of CPU cores the app should use
func setNoOfCPUCores(cpuCorePer float64) {
	corePer := config.GlobalAppConfig.Performance.UseCorePercentage
	if corePer <= 0 {
		fmt.Printf("No of Cpu Core to be Used = 1\n")
		return
	}
	totalCpus := float64(runtime.NumCPU())
	cpuCore := totalCpus * (corePer / 100)
	if cpuCore <= 0 {
		fmt.Printf("No of Cpu Core to be Used = 1\n")
		return
	}

	if cpuCore > totalCpus {
		cpuCore = totalCpus
	}
	fmt.Printf("No of Cpu Core to be Used = %d\n", int(cpuCore))
	runtime.GOMAXPROCS(int(cpuCore))
}

func InitHealthCheck() {
	healthCheckArray := make([]healthcheck.HCInterface, len(apiList))
	//get all Healthcheck instances
	for i, apiInstance := range apiList {
		healthCheckArray[i] = apiInstance.GetHealthCheck()
	}
	healthcheck.Initialise(healthCheckArray)
}

// InitMonitor initializes the monitor
func InitMonitor() {
	if err := monitor.Initialize(&config.GlobalAppConfig.MonitorConfig); err != nil {
		logger.Error(fmt.Sprintln(err))
	}
}

// InitHTTPPool: initialize http pool
func InitHTTPPool() {
	http.InitConnPool(&config.GlobalAppConfig.HTTPConfig)
}

func initProfiler() {
	profiler.InitProfiler(config.GlobalAppConfig.Profiler.SamplingRate)
}
