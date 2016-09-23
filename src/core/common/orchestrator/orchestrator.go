package orchestrator

import (
	"fmt"
	"github.com/jabong/florest-core/src/common/logger"
)

/*
Instance of the workflow
*/
type Orchestrator struct {
	//Workflow definition
	workflow *WorkFlowDefinition
}

/*
TODO: This should read the workflow configuration file
and create the appropriate orchestrator
*/
func (o *Orchestrator) CreateFromConfig(fileName string) error {
	return nil
}

/*
Constructor function for the pipeline creation
create(workflow)
*/
func (o *Orchestrator) Create(workflowdefinition *WorkFlowDefinition) error {
	jfErr := workflowdefinition.createJoinForkMapping()
	if jfErr != nil {
		return jfErr
	}
	o.workflow = workflowdefinition
	return nil
}

//Helper function to execute the Execution Node
func execExecuteNode(execNodeID string,
	execNode WorkFlowExecuteNodeInterface,
	wfData *WorkFlowData,
	wfDefinition *WorkFlowDefinition) (nextNodeID string, nextwfData *WorkFlowData) {

	nextwfData = wfData

	outputData, err := execNode.Execute(*wfData)
	if err != nil {
		nextwfData.setWorkflowState(execNode.Name(), err)
		return "", nextwfData
	}
	nextwfData = &outputData
	nextNodeIDs, found := wfDefinition.edges[execNodeID]

	if !found {
		return "", nextwfData
	}

	nextNodeID = nextNodeIDs[0]
	return nextNodeID, nextwfData
}

//Helper function to execute the Decision Node
func execDecisionNode(decisionNodeID string,
	decisionNode WorkFlowDecisionNodeInterface,
	wfData *WorkFlowData,
	wfDefinition *WorkFlowDefinition) (nextNodeID string, nextwfData *WorkFlowData) {

	nextwfData = wfData

	yes, err := decisionNode.GetDecision(*wfData)
	if err != nil {
		nextwfData.setWorkflowState(decisionNode.Name(), err)
		return "", nextwfData
	}

	nextNodeIDs, found := wfDefinition.edges[decisionNodeID]

	if !found {
		return "", nextwfData
	}

	if yes {
		//Yes Node
		nextNodeID = nextNodeIDs[0]
	}

	if !yes {
		//No Node
		nextNodeID = nextNodeIDs[1]
	}

	return nextNodeID, nextwfData
}

func execForkWorkFlow(forkNodeID string,
	wfDefinition *WorkFlowDefinition,
	wfData *WorkFlowData,
	joinNodeID string,
	wfDataChannel chan *WorkFlowData) {

	//Fork Workflow path data passed into workflow data channel
	wfDataChannel <- run(forkNodeID, wfDefinition, wfData, joinNodeID)
}

//Helper function to execute the Fork Node
func execForkNode(forkNodeID string,
	forkNode WorkFlowForkNodeInterface,
	wfData *WorkFlowData,
	wfDefinition *WorkFlowDefinition) (nextNodeID string, nextwfData *WorkFlowData) {

	wfDataChannel := make(chan *WorkFlowData)

	joinNodeID := wfDefinition.joinFork[forkNodeID]
	forkNodesID := wfDefinition.edges[forkNodeID]

	//Execute concurrently the fork node paths
	for _, forkedNodeID := range forkNodesID {
		clonedWfData := wfData.Clone()
		go execForkWorkFlow(forkedNodeID, wfDefinition, &clonedWfData, joinNodeID, wfDataChannel)
	}

	var joinNodeWfData []*WorkFlowData
	for i := 0; i < len(forkNodesID); i++ {
		wfDataFromChannel := <-wfDataChannel
		joinNodeWfData = append(joinNodeWfData, wfDataFromChannel)
	}

	//Pass the data to Join Node
	jNode, found := wfDefinition.nodes[joinNodeID]
	if !found {
		errString := fmt.Sprintln("Node id ", joinNodeID, " not present for execution")
		wfData.setWorkflowState("WORKFLOW_ERROR", errString)
		return "", wfData
	}
	logger.Info("Current Node : " + jNode.Name())
	if joinNode, ok := jNode.(WorkFlowJoinNodeInterface); ok {
		logger.Info("Execute Node")
		return executeJoinNode(joinNodeID, joinNode, wfData, joinNodeWfData, wfDefinition)
	}

	return "", wfData
}

//Helper function to execute Join Node
func executeJoinNode(joinNodeID string,
	joinNode WorkFlowJoinNodeInterface,
	forkWfData *WorkFlowData,
	joinWfData []*WorkFlowData,
	wfDefinition *WorkFlowDefinition) (nextNodeID string, nextwfData *WorkFlowData) {

	nextwfData = forkWfData
	outputData, err := joinNode.Join(joinWfData)
	if err != nil {
		nextwfData.setWorkflowState(joinNode.Name(), err)
		return "", nextwfData
	}
	nextwfData = &outputData

	nextNodeIDs, found := wfDefinition.edges[joinNodeID]

	if !found {
		return "", nextwfData
	}

	nextNodeID = nextNodeIDs[0]
	return nextNodeID, nextwfData
}

//Helper function to run the pipeline
func run(currNodeID string,
	wfDefinition *WorkFlowDefinition,
	wfData *WorkFlowData,
	terminateNodeID string) *WorkFlowData {

	logger.Info("Current Node id: " + currNodeID)

	//No workflow definition
	if wfDefinition == nil {
		return wfData
	}

	node, found := wfDefinition.nodes[currNodeID]
	if !found {
		errString := fmt.Sprintln("Node id ", currNodeID, " not present for execution")
		wfData.setWorkflowState("WORKFLOW_ERROR", errString)
		return wfData
	}
	logger.Info("Current Node : " + node.Name())

	var nextNodeID string
	nextwfData := wfData

	if execNode, ok := node.(WorkFlowExecuteNodeInterface); ok {
		logger.Info("Execute Node")
		nextNodeID, nextwfData = execExecuteNode(currNodeID, execNode,
			wfData, wfDefinition)
	}

	if decisionNode, ok := node.(WorkFlowDecisionNodeInterface); ok {
		logger.Info("Decision Node")
		nextNodeID, nextwfData = execDecisionNode(currNodeID, decisionNode,
			wfData, wfDefinition)
	}

	if forkedNode, ok := node.(WorkFlowForkNodeInterface); ok {
		logger.Info("Fork Node")
		nextNodeID, nextwfData = execForkNode(currNodeID, forkedNode,
			wfData, wfDefinition)
	}

	if nextNodeID == terminateNodeID {
		//End of execution
		return nextwfData
	}
	logger.Info("Next Node id: " + nextNodeID)
	return run(nextNodeID, wfDefinition, nextwfData, terminateNodeID)
}

/*
Workflow execution begins here
the caller should create the Work Flow State
which has the InputOutput, ExecutionContext
*/
func (o *Orchestrator) Start(wfData *WorkFlowData) *WorkFlowData {

	if o.workflow == nil {
		logger.Error("Error Empty workflow definition passed for execution")
		return new(WorkFlowData)
	}
	return run(o.workflow.startNodeID, o.workflow, wfData, "")
}

func (o *Orchestrator) String() string {
	return o.workflow.String()
}

/*
Orchestrator implements the version manager GetInstance
*/
func (o Orchestrator) GetInstance() interface{} {
	return o
}
