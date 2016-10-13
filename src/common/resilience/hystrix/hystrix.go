package hystrix

import (
	"time"

	hys "github.com/afex/hystrix-go/hystrix"
)

func Go(name string, runFunc func() error, fallbackFunc func(error) error) chan error {
	return hys.Go(name, runFunc, fallbackFunc)
}

// ConfigureCommand changes the default command configuration of the specified command
func ConfigureCommand(cmdName string, conf HCommandConf) {
	hc := hys.CommandConfig{
		Timeout:                conf.Timeout,
		MaxConcurrentRequests:  conf.MaxConcurrentRequests,
		RequestVolumeThreshold: conf.RequestVolumeThreshold,
		SleepWindow:            conf.SleepWindow,
		ErrorPercentThreshold:  conf.ErrorPercentThreshold,
	}
	hys.ConfigureCommand(cmdName, hc)
}

// ConfigureCommands changes the default command configurations
func ConfigureCommands(cmds map[string]HCommandConf) {
	for k, v := range cmds {
		ConfigureCommand(k, v)
	}
}

// GetCommandConfigurations returns a map containing all command configrations indexed by command name
func GetCommandConfigurations() map[string]HCommandConf {
	cmdConf := hys.GetCircuitSettings()
	if cmdConf == nil {
		return nil
	}
	res := make(map[string]HCommandConf, len(cmdConf))
	for k, v := range cmdConf {
		res[k] = HCommandConf{
			Timeout:                int(v.Timeout / time.Millisecond),
			MaxConcurrentRequests:  v.MaxConcurrentRequests,
			RequestVolumeThreshold: int(v.RequestVolumeThreshold),
			SleepWindow:            int(v.SleepWindow / time.Millisecond),
			ErrorPercentThreshold:  v.ErrorPercentThreshold,
		}
	}
	return res
}
