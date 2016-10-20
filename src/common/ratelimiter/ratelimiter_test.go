package ratelimiter

import (
	"fmt"
	"testing"
	"time"
)

type testResult struct {
	Iteration  int
	Bucket     string
	Exceeded   bool
	Limit      int
	Remaining  int
	RetryAfter time.Duration
}

func (o1 testResult) equals_helper(o2 testResult) bool {
	return o1.Iteration == o2.Iteration &&
		o1.Bucket == o2.Bucket &&
		o1.Exceeded == o2.Exceeded &&
		o1.Limit == o2.Limit &&
		o1.Remaining == o2.Remaining
}

func (o1 testResult) Equals(o2 testResult) bool {
	if !o1.Exceeded && !o2.Exceeded {
		return o1.equals_helper(o2) &&
			o1.RetryAfter == -1 &&
			o2.RetryAfter == -1
	} else if o1.Exceeded && o2.Exceeded {
		return o1.equals_helper(o2) &&
			o1.RetryAfter != -1 &&
			o2.RetryAfter != -1 &&
			o1.RetryAfter < o2.RetryAfter
	} else {
		return o1.equals_helper(o2) &&
			o1.RetryAfter == o2.RetryAfter
	}
}

func TestGCRARateSetting(t *testing.T) {
	conf := new(Config)

	conf.Type = GCRA
	conf.MaxRate = 0

	_, err := New(conf)
	if err == nil {
		t.Fatal("Should throw init error")
	}

}

func TestGCRABurstSetting(t *testing.T) {
	conf := new(Config)

	conf.Type = GCRA
	conf.MaxRate = 1
	conf.MaxBurst = -1

	_, err := New(conf)
	if err == nil {
		t.Fatal("Should throw init error")
	}

}

func TestGCRA(t *testing.T) {
	conf := new(Config)

	conf.Type = GCRA
	// 2 request per second
	conf.MaxRate = 2
	conf.MaxBurst = 5

	rateLimiter, err := New(conf)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Bucket according to the number i / 10 (so 1 falls into the bucket 0
	// while 11 falls into the bucket 1). This has the effect of allowing a
	// burst of 5 plus 1 (a single emission interval) on every ten iterations
	// of the loop. See the output for better clarity here.
	var res = []testResult{
		testResult{Iteration: 0, Bucket: "bucket:0", Exceeded: false, Limit: 6, Remaining: 5, RetryAfter: -1},
		testResult{Iteration: 1, Bucket: "bucket:0", Exceeded: false, Limit: 6, Remaining: 4, RetryAfter: -1},
		testResult{Iteration: 2, Bucket: "bucket:0", Exceeded: false, Limit: 6, Remaining: 3, RetryAfter: -1},
		testResult{Iteration: 3, Bucket: "bucket:0", Exceeded: false, Limit: 6, Remaining: 2, RetryAfter: -1},
		testResult{Iteration: 4, Bucket: "bucket:0", Exceeded: false, Limit: 6, Remaining: 1, RetryAfter: -1},
		testResult{Iteration: 5, Bucket: "bucket:0", Exceeded: false, Limit: 6, Remaining: 0, RetryAfter: -1},
		testResult{Iteration: 6, Bucket: "bucket:0", Exceeded: true, Limit: 6, Remaining: 0, RetryAfter: time.Second},
		testResult{Iteration: 7, Bucket: "bucket:0", Exceeded: true, Limit: 6, Remaining: 0, RetryAfter: time.Second},
		testResult{Iteration: 8, Bucket: "bucket:0", Exceeded: true, Limit: 6, Remaining: 0, RetryAfter: time.Second},
		testResult{Iteration: 9, Bucket: "bucket:0", Exceeded: true, Limit: 6, Remaining: 0, RetryAfter: time.Second},
		testResult{Iteration: 10, Bucket: "bucket:1", Exceeded: false, Limit: 6, Remaining: 5, RetryAfter: -1},
		testResult{Iteration: 11, Bucket: "bucket:1", Exceeded: false, Limit: 6, Remaining: 4, RetryAfter: -1},
		testResult{Iteration: 12, Bucket: "bucket:1", Exceeded: false, Limit: 6, Remaining: 3, RetryAfter: -1},
		testResult{Iteration: 13, Bucket: "bucket:1", Exceeded: false, Limit: 6, Remaining: 2, RetryAfter: -1},
		testResult{Iteration: 14, Bucket: "bucket:1", Exceeded: false, Limit: 6, Remaining: 1, RetryAfter: -1},
		testResult{Iteration: 15, Bucket: "bucket:1", Exceeded: false, Limit: 6, Remaining: 0, RetryAfter: -1},
		testResult{Iteration: 16, Bucket: "bucket:1", Exceeded: true, Limit: 6, Remaining: 0, RetryAfter: time.Second},
		testResult{Iteration: 17, Bucket: "bucket:1", Exceeded: true, Limit: 6, Remaining: 0, RetryAfter: time.Second},
		testResult{Iteration: 18, Bucket: "bucket:1", Exceeded: true, Limit: 6, Remaining: 0, RetryAfter: time.Second},
		testResult{Iteration: 19, Bucket: "bucket:1", Exceeded: true, Limit: 6, Remaining: 0, RetryAfter: time.Second},
	}

	for i := 0; i < 20; i++ {
		bucket := fmt.Sprintf("bucket:%v", i/10)

		exceeded, result, err := rateLimiter.RateLimit(bucket)
		if err != nil {
			t.Fatal(err.Error())
		}

		iterRes := testResult{Iteration: i,
			Bucket:     bucket,
			Exceeded:   exceeded,
			Limit:      result.Limit,
			Remaining:  result.Remaining,
			RetryAfter: result.RetryAfter,
		}
		if !iterRes.Equals(res[i]) {
			t.Fatalf("Test GCRA failed at iteration %v: %v expected to be %v", i, iterRes, res[i])
		}
	}

}
