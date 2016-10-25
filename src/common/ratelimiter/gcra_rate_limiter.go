package ratelimiter

import (
	"gopkg.in/throttled/throttled.v2"
	"gopkg.in/throttled/throttled.v2/store/memstore"
)

// GCRARateLimiter is a RateLimiter that uses the Generic Cell Rate Algorithm
// Reference https://en.wikipedia.org/wiki/Generic_cell_rate_algorithm
type GCRARateLimiter struct {
	// Number of requests per second
	// Must be greater than 0
	MaxRate int

	// Number of requests that will be allowed to exceed the rate
	// in a single burst
	// Must be greater than or equal to 0
	MaxBurst int

	// Limiter worker
	rateLimiter *throttled.GCRARateLimiter

	// Stores the state of the GCRA Rate limiter
	store *memstore.MemStore

	// Quota stores the maxr rate per second and max burst
	quota *throttled.RateQuota
}

// Init initializes the GCRA rate limiter
func (o *GCRARateLimiter) Init(conf *Config) *Error {

	o.MaxRate = conf.MaxRate
	o.MaxBurst = conf.MaxBurst

	var serr error
	o.store, serr = memstore.New(65536)
	if serr != nil {
		return getErrObj(ErrInitialization, serr.Error()+" GCRA memstore error")
	}

	o.quota = &throttled.RateQuota{MaxRate: throttled.PerSec(o.MaxRate),
		MaxBurst: o.MaxBurst}

	var rerr error
	o.rateLimiter, rerr = throttled.NewGCRARateLimiter(o.store, *o.quota)
	if rerr != nil {
		return getErrObj(ErrInitialization, rerr.Error()+" GCRA rate limiter not initialized")
	}

	return nil
}

// RateLimit checks whether a particular key has exceeded a rate
// limit. It also returns a RateLimitResult to provide additional
// information about the state of the RateLimiter.
func (o *GCRARateLimiter) RateLimit(key string) (exceeded bool,
	res *RateLimitResult, err *Error) {

	res = new(RateLimitResult)
	exceeded, rres, rerr := o.rateLimiter.RateLimit(key, 1)

	if rerr != nil {
		return exceeded, nil, getErrObj(ErrFatal, rerr.Error())
	}

	res.Limit = rres.Limit
	res.Remaining = rres.Remaining
	res.ResetAfter = rres.ResetAfter
	res.RetryAfter = rres.RetryAfter

	return exceeded, res, nil

}
