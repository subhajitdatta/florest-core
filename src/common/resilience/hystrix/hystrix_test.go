package hystrix

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

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

func TestCmdConfigure(t *testing.T) {
	tD := map[string]HCommandConf{
		"test1": HCommandConf{
			Timeout:                100,
			MaxConcurrentRequests:  10,
			RequestVolumeThreshold: 5,
			SleepWindow:            40,
			ErrorPercentThreshold:  50,
		},
		"test2": HCommandConf{
			Timeout:                200,
			MaxConcurrentRequests:  20,
			RequestVolumeThreshold: 25,
			SleepWindow:            20,
			ErrorPercentThreshold:  20,
		},
	}
	ConfigureCommands(tD)
	res := GetCommandConfigurations()

	test1, ok1 := res["test1"]
	if !ok1 {
		t.Errorf("\n Not able to find test1 in command configuration")
	}
	if !reflect.DeepEqual(test1, res["test1"]) {
		t.Errorf("\ntest1 Failed Expected: %+v\nGot: %+v", res["test1"], test1)
	}

	test2, ok2 := res["test2"]
	if !ok2 {
		t.Errorf("\n Not able to find test2 in command configuration")
	}
	if !reflect.DeepEqual(test2, res["test2"]) {
		t.Errorf("\ntest2 Failed Expected: %+v\nGot: %+v", res["test2"], test2)
	}
}

func TestHystrixGoSucces(t *testing.T) {
	RegisterDatadogCollector(dummyTestDatadogClient{})
	output := make(chan bool, 1)
	errors := Go("testSuccess", func() error {
		output <- true
		return nil
	}, nil)
	select {
	case out := <-output:
		if !out {
			t.Error("\nFailed")
		}
	case err := <-errors:
		t.Errorf("\nFailed. Error %v", err)
	}
}

func TestHystrixGoFailure(t *testing.T) {
	RegisterDatadogCollector(dummyTestDatadogClient{})
	errMsg := "Testing Failure"
	failExec := make(chan bool, 1)
	errors := Go("testFailure", func() error {
		return errors.New(errMsg)
	}, func(err error) error {
		return err
	})
	select {
	case _ = <-failExec:
		t.Error("\nTest Hystrix Go Failure Failed")
	case act := <-errors:
		exp := fmt.Sprintf("fallback failed with '%v'. run error was '%v'", errMsg, errMsg)
		if exp != act.Error() {
			t.Errorf("\nTest Hystrix Go Failure.\nExpected: %v\nGot: %v", exp, act)
		}
	}
}
