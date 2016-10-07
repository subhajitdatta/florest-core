package config

import (
	"github.com/jabong/florest-core/src/common/monitor"
	"github.com/jabong/florest-core/src/common/utils/http"
)

// AppConfig will contain all the app related config data which should be provided at the start of the app
type AppConfig struct {
	AppName           string `json:"AppName"`
	AppVersion        string `json:"AppVersion"`
	ServerPort        string
	LogConfFile       string
	MonitorConfig     monitor.MConf
	Performance       PerformanceConfigs
	DynamicConfig     DynamicConfigInfo
	HTTPConfig        http.Config `json:"HttpConfig"`
	Profiler          ProfilerConfig
	ResponseHeaders   ResponseHeaderFields
	ApplicationConfig interface{}
}

// PerformanceConfigs contains Garbage Collector detials, which will determine when the GC will kick
type PerformanceConfigs struct {
	UseCorePercentage float64
	GCPercentage      float64
}

// ResponseHeaderFields
type ResponseHeaderFields struct {
	CacheControl CacheControlHeaders
}

// CacheControlHeaders helps in telling the caller whether to cache the response and up to what time etc.
type CacheControlHeaders struct {
	ResponseType    string
	NoCache         bool
	NoStore         bool
	MaxAgeInSeconds int
}

// Application
type Application struct {
	ResponseHeaders ResponseHeaderFields
}

// DynamicConfigInfo
type DynamicConfigInfo struct {
	Active          bool
	RefreshInterval int
	ConfigKey       string
	CacheKey        string
}

// ProfilerConfig is used to profile the application, like the time taken for a request etc.
type ProfilerConfig struct {
	Enable       bool
	SamplingRate float64
}

// GlobalAppConfig is applicationconfig Singleton
var GlobalAppConfig = new(AppConfig)
