package versionmanager

import (
	"errors"
	"github.com/jabong/florest-core/src/common/ratelimiter"
)

/*
Instance of the Version Manager
*/
type versionManager struct {
	//Mapping of the resource, version, action and Executable type
	mapping VersionMap
}

var vmgr *versionManager

/*
Constructor for the version manager
*/
func Initialize(m VersionMap) {
	if vmgr != nil {
		return
	}
	vmgr = new(versionManager)
	vmgr.mapping = m
}

/*
Get the executable for the resource, version, action, bucketId
*/
func Get(resource string, version string, action string,
	bucketID string, pathParams string) (Versionable, *ratelimiter.RateLimiter, map[string]string, error) {
	if vmgr == nil {
		return nil, nil, nil, errors.New("Version manager not initialized")
	}

	if vmgr.mapping == nil {
		return nil, nil, nil, errors.New("No version mapping present")
	}

	ver := BasicVersion{
		Resource: resource,
		Version:  version,
		Action:   action,
		BucketID: bucketID,
	}

	param, ok := vmgr.mapping[ver]
	if !ok {
		return nil, nil, nil, errors.New("Versionable not found in version manager")
	}

	parameters := make(map[string]string)
	versionable, ratelimiter, err := param.GetVersionable(pathParams, &parameters)

	if err != nil {
		return nil, nil, nil, err
	}

	return versionable, ratelimiter, parameters, nil
}
