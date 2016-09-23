package constants

const (

	// Request
	Request = "REQUEST"
	// RequestContext
	RequestContext = "REQUEST_CONTEXT"
	// RequestBodyParam
	RequestBodyParam = "REQUEST_BODY_PARAM"
	// RequestPathParam
	RequestPathParam = "REQUEST_PATH_PARAMETER"

	// Response
	Response = "RESPONSE"
	// ResponseMetaData
	ResponseMetaData = "RESPONSE_META_DATA"
	// ResponseData
	ResponseData          = "RESPONSE_DATA"
	ResponseStatus        = "RESPONSE_STATUS"
	ResponseHeadersConfig = "RESPONSE_HEADERS_CONFIG"
	APIResponse           = "API_RESPONSE"

	APPID    = "APP-ID"
	APPError = "APPERROR"

	HTTPVerb = "HTTPVERB"
	URI      = "URI"

	BucketID = "BUCKETID"

	Resource   = "RESOURCE"
	Version    = "VERSION"
	Action     = "ACTION"
	PathParams = "PATH_PARAMS"

	Result = "RESULT"

	UserID        = "USER_ID"
	SessionID     = "SESSION_ID"
	AuthToken     = "AUTHTOKEN"
	UserAgent     = "USER_AGENT"
	HTTPReferrer  = "HTTP_REFERRER"
	BucketsList   = "BUCKETSLIST"
	RequestID     = "REQUEST_ID"
	TransactionID = "TRANSACTION_ID"
	TokenID       = "TOKEN_ID"
	Debug         = "DEBUG"

	FieldSeperator    = ","
	KeyValueSeperator = ":"

	HealthCheckAPI  = "HEALTHCHECK"
	HealthCheckList = "HEALTH_CHECK_LIST"
)

const (
	SortAsc  = "asc"
	SortDesc = "desc"
)

//constant for buckets and their values
const (
	OrchestratorBucketKey          = "Algo"
	OrchestratorBucketDefaultValue = "Old"
	OrchestratorBucketNewAlgo      = "New"
)

// constant for monitor
const (
	MonitorCustomMetric = "monitor_custom_metric"
)
