package service

import (
	"github.com/jabong/florest-core/src/common/constants"
	workflow "github.com/jabong/florest-core/src/core/common/orchestrator"
)

func getServiceVersion(data workflow.WorkFlowData) (resource string,
	version string, action string, orchBucket string, pathParams string) {
	resourceData, _ := data.IOData.Get(constants.Resource)
	versionData, _ := data.IOData.Get(constants.Version)
	actionData, _ := data.IOData.Get(constants.Action)
	pathParamsData, _ := data.IOData.Get(constants.PathParams)
	bucketsMap, _ := data.ExecContext.GetBuckets()

	resource = ""
	if v, ok := resourceData.(string); ok {
		resource = v
	}

	version = ""
	if v, ok := versionData.(string); ok {
		version = v
	}

	action = ""
	if v, ok := actionData.(string); ok {
		action = v
	}

	pathParams = ""
	if v, ok := pathParamsData.(string); ok {
		pathParams = v
	}

	orchBucket = getBucketValue(resource, bucketsMap)
	return resource, version, action, orchBucket, pathParams
}

func getBucketValue(resource string, bucketsMap map[string]string) string {
	orchBucket := constants.OrchestratorBucketDefaultValue
	resourceBucketKey, present := resourceBucketMapping[resource]
	if !present {
		return orchBucket
	}

	orchbucketID, found := bucketsMap[resourceBucketKey]
	if found {
		orchBucket = orchbucketID
	}
	return orchBucket
}
