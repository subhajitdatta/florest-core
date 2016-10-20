package ratelimiter

import (
	"strings"
)

func New(conf *Config) (res RateLimiter, err *Error) {
	if conf == nil {
		err = getErrObj(ErrInitialization, "Rate limiter empty config passed")
		return nil, err
	}

	if conf.MaxRate <= 0 {
		return nil, getErrObj(ErrInitialization, "Max rate must be greater than 0")
	}

	if conf.MaxBurst < 0 {
		return nil, getErrObj(ErrInitialization, "Max burst must be greater than or equal to 0")
	}

	switch strings.ToUpper(conf.Type) {
	case GCRA:
		res = new(GCRARateLimiter)
		err = res.Init(conf)
	default:
		res = new(GCRARateLimiter)
		err = res.Init(conf)
	}

	return res, err
}
