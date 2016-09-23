package orchestrator

import (
	"testing"
)

/*
Test Workflow Execution Node
*/
type testWfExecNode struct {
	id string
}

func (n testWfExecNode) Name() string {
	return "testWfExecNode"
}

func (n *testWfExecNode) SetID(id string) {
	n.id = id
}

func (n testWfExecNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testWfExecNode) Execute(data WorkFlowData) (WorkFlowData, error) {
	return data, nil
}

/*
Test Workflow Decision Node
*/
type testWfDecisionNode struct {
	id string
}

func (n testWfDecisionNode) Name() string {
	return "testWfDecisionNode"
}

func (n *testWfDecisionNode) SetID(id string) {
	n.id = id
}

func (n testWfDecisionNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testWfDecisionNode) GetDecision(data WorkFlowData) (bool, error) {
	return true, nil
}

/*
Test Workflow Fork Node
*/
type testWfForkNode struct {
	id string
}

func (n testWfForkNode) Name() string {
	return "testWfForkNode"
}

func (n *testWfForkNode) SetID(id string) {
	n.id = id
}

func (n testWfForkNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testWfForkNode) Fork(data WorkFlowData) (WorkFlowData, error) {
	return data, nil
}

/*
Test Workflow Join Node
*/
type testWfJoinNode struct {
	id string
}

func (n testWfJoinNode) Name() string {
	return "testWfJoinNode"
}

func (n *testWfJoinNode) SetID(id string) {
	n.id = id
}

func (n testWfJoinNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testWfJoinNode) Join(data []*WorkFlowData) (WorkFlowData, error) {
	return *data[0], nil
}

/*
Test Workflow definition creation
*/
func TestWorkflowDefinitionCreate(t *testing.T) {
	testWfDefinition := new(WorkFlowDefinition)
	testWfDefinition.Create()

	if testWfDefinition.nodes == nil {
		t.Error("Workflow definition creation failed")
	}

	if testWfDefinition.edges == nil {
		t.Error("Workflow definition creation failed")
	}
}

/*
Test workflow definition set/get start node
*/
func TestWorkflowDefinitionSetGetStartNode(t *testing.T) {
	testWfDefinition := new(WorkFlowDefinition)
	testWfDefinition.Create()

	testWfNode := new(testWfExecNode)
	testWfNode.SetID("Execution Node 1")
	testWfDefinition.AddExecutionNode(testWfNode)

	serr := testWfDefinition.SetStartNode(testWfNode)
	if serr != nil {
		t.Error("Failed to set workflow start node")
	}

	wfNode, gerr := testWfDefinition.GetStartNode()
	if gerr != nil || wfNode == nil {
		t.Error("Failed to get workflow start node")
	}
}

/*
Test workflow definition add execution node
*/
func TestWorkflowDefinitionAddExecutionNode(t *testing.T) {
	testWfDefinition := new(WorkFlowDefinition)
	testWfDefinition.Create()

	testWfNode := new(testWfExecNode)
	testWfNode.SetID("Execution Node")
	testWfDefinition.AddExecutionNode(testWfNode)

	execNode, found := testWfDefinition.nodes["Execution Node"]

	if !found {
		t.Error("Failed to add Execution node to workflow definition")
	}

	_, isExecutionNodeType := execNode.(WorkFlowExecuteNodeInterface)
	if !isExecutionNodeType {
		t.Error("Mismatch in the data type of the workflow execution node")
	}
}

/*
Test workflow definition add execution node
*/
func TestWorkflowDefinitionAddDecisionNode(t *testing.T) {
	testWfDefinition := new(WorkFlowDefinition)
	testWfDefinition.Create()

	testWfDecNode := new(testWfDecisionNode)
	testWfDecNode.SetID("Decision Node")

	testWfYesNode := new(testWfExecNode)
	testWfYesNode.SetID("Yes Execution Node")

	testWfNoNode := new(testWfExecNode)
	testWfNoNode.SetID("No Execution Node")

	testWfDefinition.AddDecisionNode(testWfDecNode, testWfYesNode, testWfNoNode)

	decisionNode, found := testWfDefinition.nodes["Decision Node"]

	if !found {
		t.Error("Failed to add Execution node to workflow definition")
	}

	_, isDecisionNodeType := decisionNode.(WorkFlowDecisionNodeInterface)
	if !isDecisionNodeType {
		t.Error("Mismatch in the data type of the workflow decision node")
	}
}

/*
Test workflow definition add connection
*/
func TestWorkflowDefinitionAddConnection(t *testing.T) {
	testWfDefinition := new(WorkFlowDefinition)
	testWfDefinition.Create()

	testWfFromNode := new(testWfExecNode)
	testWfFromNode.SetID("From Execution Node")
	testWfDefinition.AddExecutionNode(testWfFromNode)

	testWfNoNode := new(testWfExecNode)
	testWfNoNode.SetID("To Execution Node")
	testWfDefinition.AddExecutionNode(testWfNoNode)

	testWfDefinition.AddConnection(testWfFromNode, testWfNoNode)

	edgeList, found := testWfDefinition.edges["From Execution Node"]

	if !found || edgeList == nil {
		t.Error("Failed to add connection between execution nodes in workflow definition")
	}

	if len(edgeList) != 1 {
		t.Error("Incorrect list of edges returned from workflow definition")
	}

	if edgeList[0] != "To Execution Node" {
		t.Error("Failed to retrieve edge added in workflow definition")
	}
}

/*
Test workflow definition join fork node creation
*/
func TestWorkflowDefinitionAddForkJoinNodes(t *testing.T) {
	testWfDefinition := new(WorkFlowDefinition)
	testWfDefinition.Create()

	testWfNode1 := new(testWfExecNode)
	testWfNode1.SetID("Execution Node 1")
	testWfDefinition.AddExecutionNode(testWfNode1)

	testWfNode2 := new(testWfExecNode)
	testWfNode2.SetID("Execution Node 2")
	testWfDefinition.AddExecutionNode(testWfNode2)

	testWfNode3 := new(testWfExecNode)
	testWfNode3.SetID("Execution Node 3")
	testWfDefinition.AddExecutionNode(testWfNode3)

	testWfFrkNode := new(testWfForkNode)
	testWfFrkNode.SetID("Fork Node")
	testWfDefinition.AddForkNode(testWfFrkNode, []WorkFlowNodeInterface{testWfNode1, testWfNode2})

	testWfJnNode := new(testWfJoinNode)
	testWfJnNode.SetID("Join Node")
	testWfDefinition.AddJoinNode(testWfJnNode)

	testWfDefinition.AddConnection(testWfNode2, testWfNode3)
	testWfDefinition.AddConnection(testWfNode1, testWfJnNode)
	testWfDefinition.AddConnection(testWfNode3, testWfJnNode)

	testWfDefinition.SetStartNode(testWfFrkNode)
	jferr := testWfDefinition.createJoinForkMapping()
	if jferr != nil {
		t.Error("Failed to create the join fork mapping")
	}

	joinNode, found := testWfDefinition.joinFork["Fork Node"]
	t.Log(testWfDefinition.nodes)
	t.Log(testWfDefinition.edges)
	t.Log(testWfDefinition.joinFork)

	if !found {
		t.Error("Failed to retrieve join node added in workflow definition")
	}

	if joinNode != "Join Node" {
		t.Error("Incorrect join node returned from workflow definition")
	}

}

/*
Test workflow definition toString
*/
func TestWorkflowDefinitionToString(t *testing.T) {
	testWfDefinition := new(WorkFlowDefinition)
	testWfDefinition.Create()

	if testWfDefinition.String() == "" {
		t.Error("Empty tostring in workflow definition")
	}
}
