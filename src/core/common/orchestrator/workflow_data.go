package orchestrator

/*
Data structure to hold the workflow data passed from one node to the other
*/
type WorkFlowData struct {
	IOData      WorkFlowIOInterface
	ExecContext WorkFlowExecutionContextInterface
	state       workFlowState
}

/*
Set the workflow state for a workflow node
*/
func (d *WorkFlowData) setWorkflowState(nodeName string, nodeState interface{}) error {
	return d.state.set(nodeName, nodeState)
}

/*
Get the states for all the work flow nodes
*/
func (d *WorkFlowData) GetWorkflowState() map[string]interface{} {
	return d.state.GetAll()
}

/*
Initialize the workflow data
*/
func (d *WorkFlowData) Create(wfio WorkFlowIOInterface, ec WorkFlowExecutionContextInterface) {
	d.IOData = wfio
	d.ExecContext = ec

	d.state = *new(workFlowState)
	d.state.create()
}

/*
Create clone of the workflow data
*/
func (d *WorkFlowData) Clone() WorkFlowData {
	//Only clone the io data
	ioClone := d.IOData.Clone()
	ioCloneData, _ := ioClone.(WorkFlowIOInterface)
	return WorkFlowData{IOData: ioCloneData,
		ExecContext: d.ExecContext,
		state:       d.state}
}
