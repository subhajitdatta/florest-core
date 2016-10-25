package versionmanager

import (
	"errors"
	"github.com/jabong/florest-core/src/common/ratelimiter"
	"strings"
)

/*
Data structure to store the API that is versioned
*/
type BasicVersion struct {
	Resource string
	Version  string
	Action   string
	BucketID string
}

type Version struct {
	Resource string
	Version  string
	Action   string
	BucketID string
	Path     string
}

func (v *Version) GetBasicVersion() BasicVersion {
	return BasicVersion{
		v.Resource,
		v.Version,
		v.Action,
		v.BucketID,
	}
}

type Param struct {
	versionable Versionable
	pathParams  map[string]*Param
	namedParams map[string]*Param
	rateLimiter *ratelimiter.RateLimiter
}

func NewParam() *Param {
	p := Param{
		pathParams:  make(map[string]*Param),
		namedParams: make(map[string]*Param),
	}

	return &p
}

func (param *Param) Update(path string, versionable Versionable,
	rl *ratelimiter.RateLimiter) error {
	if path == "" {
		param.versionable = versionable
		param.rateLimiter = rl
		return nil
	}
	pathArr := strings.Split(path, "/")
	for _, pathParamKey := range pathArr {
		if strings.Index(pathParamKey, "{") == 0 {
			lastIndex := len(pathParamKey) - 1
			if strings.Index(pathParamKey, "}") != lastIndex {
				return errors.New("Invalid named param. Named param should be in between { and }")
			}
			namedParamKey := pathParamKey[1:lastIndex]
			namedParam := param.namedParams[namedParamKey]
			if namedParam == nil {
				namedParam = NewParam()
				param.namedParams[namedParamKey] = namedParam
			}
			param = namedParam
		} else {
			pathParam := param.pathParams[pathParamKey]
			if pathParam == nil {
				pathParam = NewParam()
				param.pathParams[pathParamKey] = pathParam
			}
			param = pathParam
		}
	}
	param.versionable = versionable
	param.rateLimiter = rl
	return nil
}

func (param *Param) GetVersionable(pathParams string,
	parameters *map[string]string) (Versionable, *ratelimiter.RateLimiter, error) {
	if pathParams == "" {
		if param.versionable == nil {
			return nil, nil, errors.New("Versionable not found in version manager")
		}
		return param.versionable, param.rateLimiter, nil
	}
	pathParamsArr := strings.Split(pathParams, "/")
	return param.getVersionableObj(pathParamsArr, parameters)
}

func (param *Param) getVersionableObj(pathParams []string,
	parameters *map[string]string) (Versionable, *ratelimiter.RateLimiter, error) {
	tempParams := pathParams
	notFoundError := errors.New("Versionable not found in version manager")
	for _, pathParam := range pathParams {
		tempParams = tempParams[1:]
		if param.pathParams[pathParam] != nil {
			param = param.pathParams[pathParam]
		} else {
			for key, namedParam := range param.namedParams {
				versionable, rl, err := namedParam.getVersionableObj(tempParams, parameters)
				if err == nil {
					(*parameters)[key] = pathParam
					return versionable, rl, nil
				}
			}
			return nil, nil, notFoundError
		}
	}
	if param.versionable == nil {
		return nil, nil, notFoundError
	}
	return param.versionable, param.rateLimiter, nil
}

type VersionMap map[BasicVersion]*Param
