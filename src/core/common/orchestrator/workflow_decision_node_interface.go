package orchestrator

/*
Decision Node Interface
*/
type WorkFlowDecisionNodeInterface interface {
	//Inherits from the WorkFlow Node Interface
	WorkFlowNodeInterface

	//Decision method
	GetDecision(WorkFlowData) (bool, error)
}
