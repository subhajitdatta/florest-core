package orchestrator

/*
Data structure to store the input and output for each workflow execution node
*/
type WorkFlowIOInterface interface {
	Get(key string) (value interface{}, err error)
	Set(key string, value interface{}) (err error)
	Clone() WorkFlowIOInterface
}
