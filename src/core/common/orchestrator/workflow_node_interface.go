package orchestrator

/*
Interface which has to be implemented by the workflow nodes
*/
type WorkFlowNodeInterface interface {
	//Static name for the implementing class
	Name() string

	//Set Identifier for the Node
	SetID(id string)

	//Get Identifier for the Node
	GetID() (id string, err error)
}
