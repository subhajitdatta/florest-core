package ratelimiter

import ()

type RateLimiter interface {
	// Initialize the rate limiter
	Init(*Config) *Error

	// RateLimit checks whether a particular key has exceeded a rate
	// limit. It also returns a RateLimitResult to provide additional
	// information about the state of the RateLimiter.
	RateLimit(key string) (exceeded bool, res *RateLimitResult, err *Error)
}
