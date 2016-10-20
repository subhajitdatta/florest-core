package ratelimiter

import (
	"time"
)

type RateLimitResult struct {
	// Limit is the maximum number of requests that could be permitted
	// instantaneously for this key starting from an empty state. For
	// example, if a rate limiter allows 10 requests per second per
	// key, Limit would always be 10.
	Limit int

	// Remaining is the maximum number of requests that could be
	// permitted instantaneously for this key given the current
	// state. For example, if a rate limiter allows 10 requests per
	// second and has already received 6 requests for this key this
	// second, Remaining would be 4.
	Remaining int

	// ResetAfter is the time until the RateLimiter returns to its
	// initial state for a given key. For example, if a rate limiter
	// manages requests per second and received one request 200ms ago,
	// Reset would return 800ms. You can also think of this as the time
	// until Limit and Remaining will be equal.
	ResetAfter time.Duration

	// RetryAfter is the time until the next request will be permitted.
	// It should be -1 unless the rate limit has been exceeded.
	RetryAfter time.Duration
}
