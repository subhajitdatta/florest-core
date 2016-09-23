package orchestrator

/*
Execution Node Interface
*/
type WorkFlowExecuteNodeInterface interface {
	//Inherits from the WorkFlow Node Interface
	WorkFlowNodeInterface

	//Execution method
	Execute(WorkFlowData) (WorkFlowData, error)
}
