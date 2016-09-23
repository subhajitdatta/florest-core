package orchestrator

/*
Fork Node Interface
*/
type WorkFlowForkNodeInterface interface {
	//Inherits from the WorkFlow Node Interface
	WorkFlowNodeInterface

	//Fork method
	Fork(WorkFlowData) (WorkFlowData, error)
}
