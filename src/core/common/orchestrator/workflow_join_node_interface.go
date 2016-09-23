package orchestrator

/*
Join Node Interface
*/
type WorkFlowJoinNodeInterface interface {
	//Inherits from the WorkFlow Node Interface
	WorkFlowNodeInterface

	//Join Method
	Join([]*WorkFlowData) (WorkFlowData, error)
}
