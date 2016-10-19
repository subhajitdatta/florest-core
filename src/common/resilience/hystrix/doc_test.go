package hystrix_test

import (
	"errors"
	"fmt"

	"github.com/jabong/florest-core/src/common/resilience/hystrix"
)

// dummyTestDatadogClient is a dummy datadog metrics collector implementing IDatadogClient
// interface
type dummyTestDatadogClient struct {
}

func (d dummyTestDatadogClient) Gauge(name string, value float64, tags []string, rate float64) error {
	return nil
}
func (d dummyTestDatadogClient) Count(name string, value int64, tags []string, rate float64) error {
	return nil
}
func (d dummyTestDatadogClient) Histogram(name string, value float64, tags []string, rate float64) error {
	return nil
}

func workFunc(arg bool) error {
	if !arg {
		return errors.New("Testing Failure")
	}
	fmt.Println("Hello World")
	return nil
}

func Example() {
	// Register a metrics collector
	hystrix.RegisterDatadogCollector(dummyTestDatadogClient{})
	failExec := make(chan bool, 1)
	errors := hystrix.Go("Example", func() error {
		// Execute the code which has external dependency here
		if err := workFunc(true); err != nil {
			return err
		}
		failExec <- true
		return nil
	}, func(err error) error {
		// Execute the fallback logic in case of failure
		return err
	})
	select {
	case _ = <-failExec:
		fmt.Println("Function Successful")
	case act := <-errors:
		fmt.Printf("\nFailback function called %v\n", act)
	}
}
