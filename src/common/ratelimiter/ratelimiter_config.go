package ratelimiter

import ()

type Config struct {
	// The rate limiting algorithm
	// Default is GCRA
	Type string

	// Number of requests per second
	// Must be greater than 0
	MaxRate int

	// Number of requests that will be allowed to exceed the rate
	// in a single burst
	// Must be greater than or equal to 0
	MaxBurst int
}
