package orchestrator

/*
Execution context for the workflow is maintained here.
*/
type WorkFlowExecutionContextInterface interface {
	//Get the value for the key of the execution context
	Get(key string) (value interface{}, err error)

	//Set the key, value in execution context
	Set(key string, value interface{}) (err error)

	//Set Bucket Id list
	SetBuckets(bucketIDMap map[string]string) (err error)

	//Get the Bucket Id List
	GetBuckets() (bucketIDMap map[string]string, err error)

	//Get the current path execution thread id
	GetExecuteThreadID() (executeThreadID string, err error)

	//Set Debug Flag
	SetDebugFlag(flag bool) (err error)

	//Set Debug Message
	SetDebugMsg(msgkey string, msgData string) (err error)

	//Get Debug Message
	GetDebugMsg() (msg []interface{}, err error)
}
