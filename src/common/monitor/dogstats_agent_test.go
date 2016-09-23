package monitor

import (
	"testing"
)

func getTestConfig() *MConf {
	c := new(MConf)
	c.APIKey = ""
	c.APPKey = ""
	c.APPName = "TestJadeGO"
	c.Platform = DatadogAgent
	c.AgentServer = "127.0.0.1:8125"
	c.Verbose = true
	c.Enabled = true
	return c
}

func getTestDatadogAgentClient() (MInterface, error) {
	c := getTestConfig()
	d, err := Get(c)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func TestGet(t *testing.T) {
	if _, err := getTestDatadogAgentClient(); err != nil {
		t.Errorf("Get DatadogAgentClient Failed %v", err)
	}
}

func TestInfo(t *testing.T) {
	d, err := getTestDatadogAgentClient()
	if err != nil {
		t.Errorf("Failed to get DatadogAgentClient for event info %v", err)
	}

	data := new(MData)
	data.Title = "Test Info"
	data.Body = "Hello World Info"
	data.Tags = map[string]string{"env": "local"}
	if err := d.Info(data); err != nil {
		t.Errorf("Error event failed %v", err)
	}
}

func TestGauges(t *testing.T) {
	d, err := getTestDatadogAgentClient()
	if err != nil {
		t.Errorf("Failed to get DatadogAgentClient for event Gauges %v", err)
	}
	if err := d.Gauge("Test_Gauge", 121.5, []string{"jadehol838"}, 1); err != nil {
		t.Errorf("Error event failed %v", err)
	}
}

func TestSet(t *testing.T) {
	d, err := getTestDatadogAgentClient()
	if err != nil {
		t.Errorf("Failed to get DatadogAgentClient for event Set %v", err)
	}
	if err := d.Set("Test_Set", "Hello World", []string{"jadehol838"}, 1); err != nil {
		t.Errorf("Error event failed %v", err)
	}
}
