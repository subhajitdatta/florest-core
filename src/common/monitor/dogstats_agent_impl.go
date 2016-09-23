package monitor

import (
	_ "expvar"
	"fmt"
	"github.com/ooyala/go-dogstatsd"
	"net"
	"net/http"
)

// DataDogAgent client is the implementation of a client to talk with
// the datadog agent
type DataDogAgentClient struct {
	client *dogstatsd.Client
	conf   *MConf
}

// Init initialises a datadog client for agent
func (d *DataDogAgentClient) Init(conf *MConf) (err error) {
	defer recoverFromPanic(&err)
	d.conf = conf
	if !d.conf.Enabled {
		return nil
	}
	d.client, err = dogstatsd.New(conf.AgentServer)
	if err != nil {
		logMsg(errMsgTag, err, d.conf)
		return err
	}
	d.client.Namespace = conf.APPName
	return nil
}

// Info logs an info alert type event to datadog event stream
func (d *DataDogAgentClient) Info(data *MData) (err error) {
	defer recoverFromPanic(&err)

	if !d.conf.Enabled {
		return nil
	}
	title := d.getNamespacedEventTitle(data.Title)
	err = d.client.Info(title, data.Body, getTagsArray(data.Tags))
	if err != nil {
		logMsg(errMsgTag, err, d.conf)
		return err
	}
	return nil
}

// Success logs  an success alert type event to datadog event stream
func (d *DataDogAgentClient) Success(data *MData) (err error) {
	defer recoverFromPanic(&err)

	if !d.conf.Enabled {
		return nil
	}
	title := d.getNamespacedEventTitle(data.Title)
	err = d.client.Success(title, data.Body, getTagsArray(data.Tags))
	if err != nil {
		logMsg(errMsgTag, err, d.conf)
		return err
	}
	return nil
}

// Warning logs an warning alert type event to datadog event stream
func (d *DataDogAgentClient) Warning(data *MData) (err error) {
	defer recoverFromPanic(&err)

	if !d.conf.Enabled {
		return nil
	}
	title := d.getNamespacedEventTitle(data.Title)
	err = d.client.Warning(title, data.Body, getTagsArray(data.Tags))
	if err != nil {
		logMsg(errMsgTag, err, d.conf)
		return err
	}
	return nil
}

// Error logs an error alert type event to datadog event stream
func (d *DataDogAgentClient) Error(data *MData) (err error) {
	defer recoverFromPanic(&err)

	if !d.conf.Enabled {
		return nil
	}
	title := d.getNamespacedEventTitle(data.Title)
	err = d.client.Error(title, data.Body, getTagsArray(data.Tags))
	if err != nil {
		logMsg(errMsgTag, err, d.conf)
		return err
	}
	return nil
}

// Gauge records a gauge metric in datadog server. Gauges measure the value of a particular thing over
// time.
func (d *DataDogAgentClient) Gauge(name string, value float64, tags []string, rate float64) (err error) {
	defer recoverFromPanic(&err)

	if !d.conf.Enabled {
		return nil
	}
	err = d.client.Gauge(name, value, tags, rate)
	if err != nil {
		logMsg(errMsgTag, err, d.conf)
		return err
	}
	return nil
}

// Count records a count metric in datadog server. Counters are used to (ahem) count things.
func (d *DataDogAgentClient) Count(name string, value int64, tags []string, rate float64) (err error) {
	defer recoverFromPanic(&err)

	if !d.conf.Enabled {
		return nil
	}
	err = d.client.Count(name, value, tags, rate)
	if err != nil {
		logMsg(errMsgTag, err, d.conf)
		return err
	}
	return nil
}

// Histogram records a histrogram metric in datadog server. Histograms measure the statistical
// distribution of a set of values.
func (d *DataDogAgentClient) Histogram(name string, value float64, tags []string, rate float64) (err error) {
	defer recoverFromPanic(&err)

	if !d.conf.Enabled {
		return nil
	}
	err = d.client.Histogram(name, value, tags, rate)
	if err != nil {
		logMsg(errMsgTag, err, d.conf)
		return err
	}
	return nil
}

// Set records a set metric in datadog server. Sets are used to count the number of unique elements
// in a group. If you want to track the number of unique visitor to your site, sets are a great way
// to do that.
func (d *DataDogAgentClient) Set(name string, value string, tags []string, rate float64) (err error) {
	defer recoverFromPanic(&err)

	if !d.conf.Enabled {
		return nil
	}
	err = d.client.Set(name, value, tags, rate)
	if err != nil {
		logMsg(errMsgTag, err, d.conf)
		return err
	}
	return nil
}

// SendAppMetrics starts sending AppMetrics exposed through expvar. By default go exposes the
// Memstats. For seeing the expvar stats in datadog follow the steps below:-
// i)   Modify go_expvar.yaml placed in conf.d directory of dd-agent. For ubuntu this is
//      etc/conf.d/dd-agent. For other OS please refer the datadog doc. If go_expvar.yaml does not exist
//      ,copy it from go_expvar.yaml.example
// ii)  Modify the expvar_url to match the serverIP. ServerIP should be of the form "serverAddr:port"
// iii) Restart the dd-agent
func (d *DataDogAgentClient) SendAppMetrics(serverIP string) (err error) {
	defer recoverFromPanic(&err)
	if !d.conf.Enabled {
		return nil
	}
	sock, err := net.Listen("tcp", serverIP)
	if err != nil {
		return err
	}
	go func() {
		logMsg(infoMsgtag, fmt.Sprintf("App Metrics available at %s", serverIP), d.conf)
		http.Serve(sock, nil)
	}()
	return nil
}

// getNamespacedEventName prefixes the appName with  each event name
func (d *DataDogAgentClient) getNamespacedEventTitle(n string) string {
	return d.client.Namespace + n
}

// newDogstatsdAgent creates a new instance of DataDogAgentClient
func newDogstatsdAgent(conf *MConf) (d *DataDogAgentClient, err error) {
	defer recoverFromPanic(&err)

	d = new(DataDogAgentClient)
	err = d.Init(conf)
	if err != nil {
		logMsg(errMsgTag, err, conf)
		return nil, err
	}
	return d, err
}

// getTagsArray returns a string array from a map by joining each key, value in a
// map with ':'
func getTagsArray(tags map[string]string) []string {
	arr := make([]string, len(tags))
	i := 0
	for k, v := range tags {
		arr[i] = k + ":" + v
		i++
	}
	return arr
}
