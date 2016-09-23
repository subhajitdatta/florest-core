package orchestrator

import (
	"fmt"
	"testing"
)

const (
	tESTEXECUTIONNODENAME = "Test Execution Node"
	tESTDECISIONNODENAME  = "Test Decision Node"

	dECISION    = "DECISION"
	yESNODENAME = "YESNODE"
	nONODENAME  = "NONODE"

	tESTWFFORKNODE = "TESTWFFORKNODE"
	tESTWFJOINNODE = "TESTWFJOINNODE"
)

/*
Test Execution Node
*/
type testExecNode struct {
	id string
}

func (n testExecNode) Name() string {
	return tESTEXECUTIONNODENAME
}

func (n *testExecNode) SetID(id string) {
	n.id = id
}

func (n testExecNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testExecNode) Execute(data WorkFlowData) (WorkFlowData, error) {
	data.setWorkflowState(n.Name(), fmt.Sprintln("Executing ", n.Name()))
	return data, nil
}

/*
Test Decision Node
*/
type testDecisionNode struct {
	id string
}

func (n testDecisionNode) Name() string {
	return tESTDECISIONNODENAME
}

func (n *testDecisionNode) SetID(id string) {
	n.id = id
}

func (n testDecisionNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testDecisionNode) GetDecision(data WorkFlowData) (bool, error) {
	data.setWorkflowState(n.Name(), fmt.Sprintln("Executing ", n.Name()))
	res, _ := data.IOData.Get(dECISION)
	dec, _ := res.(bool)
	return dec, nil
}

/*
Yes Node
*/
type testYesNode struct {
	id string
}

func (n testYesNode) Name() string {
	return yESNODENAME
}

func (n *testYesNode) SetID(id string) {
	n.id = id
}

func (n testYesNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testYesNode) Execute(data WorkFlowData) (WorkFlowData, error) {
	data.setWorkflowState(n.Name(), fmt.Sprintln("Executing ", n.Name()))
	return data, nil
}

/*
No Node
*/
type testNoNode struct {
	id string
}

func (n testNoNode) Name() string {
	return nONODENAME
}

func (n *testNoNode) SetID(id string) {
	n.id = id
}

func (n testNoNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testNoNode) Execute(data WorkFlowData) (WorkFlowData, error) {
	data.setWorkflowState(n.Name(), fmt.Sprintln("Executing ", n.Name()))
	return data, nil
}

/*
Test Workflow Fork Node
*/
type testForkNode struct {
	id string
}

func (n testForkNode) Name() string {
	return tESTWFFORKNODE
}

func (n *testForkNode) SetID(id string) {
	n.id = id
}

func (n testForkNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testForkNode) Fork(data WorkFlowData) (WorkFlowData, error) {
	fmt.Println("Inside Fork")
	data.setWorkflowState(n.Name(), fmt.Sprintln("Executing ", n.Name()))
	return data, nil
}

/*
Test Workflow Join Node
*/
type testJoinNode struct {
	id string
}

func (n testJoinNode) Name() string {
	return tESTWFJOINNODE
}

func (n *testJoinNode) SetID(id string) {
	n.id = id
}

func (n testJoinNode) GetID() (id string, err error) {
	return n.id, nil
}

func (n testJoinNode) Join(data []*WorkFlowData) (WorkFlowData, error) {
	fmt.Println(data)
	//data[0].setWorkflowState(n.Name(), fmt.Sprintf("Executing ", n.Name()))
	p := data[0]
	return *p, nil
}

/*
Helper to create a test workflow definition with one execution node
*/
func createTestWorkflowDefinition() *WorkFlowDefinition {
	testWorkflow := new(WorkFlowDefinition)
	testWorkflow.Create()

	testNode := new(testExecNode)
	testNode.SetID("1")
	testWorkflow.AddExecutionNode(testNode)

	testWorkflow.SetStartNode(testNode)

	return testWorkflow
}

/*
Helper to create a test workflow
*/
func createTestWorkflowData() *WorkFlowData {
	testInputOutput := new(WorkFlowIOInMemoryImpl)
	testEcContext := new(WorkFlowECInMemoryImpl)

	testWorkFlowData := new(WorkFlowData)
	testWorkFlowData.Create(testInputOutput, testEcContext)

	return testWorkFlowData
}

/*
Test for Orchestrator Creation
*/
func TestOrchestratorCreate(t *testing.T) {
	testOrchestrator := new(Orchestrator)
	testWorkflowDefinition := createTestWorkflowDefinition()
	testWorkflowDefinition.AddExecutionNode(new(testExecNode))
	testOrchestrator.Create(testWorkflowDefinition)

	if testOrchestrator.workflow == nil {
		t.Error("Failed to create orchestrator workflow")
	}
}

/*
Test for Executing workflow with single Execution Node
*/
func TestExecutionNodeRun(t *testing.T) {
	testOrchestrator := new(Orchestrator)
	testWorkflowDefinition := createTestWorkflowDefinition()
	testWorkflowDefinition.AddExecutionNode(new(testExecNode))
	testOrchestrator.Create(testWorkflowDefinition)

	testWorkFlowData := createTestWorkflowData()
	outputData := testOrchestrator.Start(testWorkFlowData)

	wfSate := outputData.GetWorkflowState()

	_, found := wfSate[tESTEXECUTIONNODENAME]
	if !found {
		t.Error("Failed execute Orchestrator with Execution Node")
	}

}

func createDecisionWorkflow() *WorkFlowDefinition {
	decisionWorkflow := new(WorkFlowDefinition)
	decisionWorkflow.Create()

	decisionNode := new(testDecisionNode)
	decisionNode.SetID("1")

	yesNode := new(testYesNode)
	yesNode.SetID("2")

	noNode := new(testNoNode)
	noNode.SetID("3")

	decisionWorkflow.AddDecisionNode(decisionNode, yesNode, noNode)
	decisionWorkflow.SetStartNode(decisionNode)

	return decisionWorkflow
}

/*
Test for Executing workflow with single decision node with yes condition
*/
func TestDecisionNodeYesCondRun(t *testing.T) {
	testOrchestrator := new(Orchestrator)
	testWorkflowDefinition := createDecisionWorkflow()
	testOrchestrator.Create(testWorkflowDefinition)

	testWorkFlowData := createTestWorkflowData()
	testWorkFlowData.IOData.Set(dECISION, true)

	outputData := testOrchestrator.Start(testWorkFlowData)

	wfSate := outputData.GetWorkflowState()

	_, found := wfSate[yESNODENAME]
	if !found {
		t.Error("Failed execute Orchestrator with Decision Node on YES condition")
	}
}

/*
Test for Executing workflow with single decision node with no condition
*/
func TestDecisionNodeNoCondRun(t *testing.T) {
	testOrchestrator := new(Orchestrator)
	testWorkflowDefinition := createDecisionWorkflow()
	testOrchestrator.Create(testWorkflowDefinition)

	testWorkFlowData := createTestWorkflowData()
	testWorkFlowData.IOData.Set(dECISION, false)

	outputData := testOrchestrator.Start(testWorkFlowData)

	wfSate := outputData.GetWorkflowState()

	_, found := wfSate[nONODENAME]
	if !found {
		t.Error("Failed execute Orchestrator with Decision Node on NO condition")
	}
}

func createForkJoinTestWorkflowDefinition() *WorkFlowDefinition {

	testWfDefinition := new(WorkFlowDefinition)
	testWfDefinition.Create()

	testWfNode1 := new(testExecNode)
	testWfNode1.SetID("E1")
	testWfDefinition.AddExecutionNode(testWfNode1)

	testWfNode2 := new(testExecNode)
	testWfNode2.SetID("E2")
	testWfDefinition.AddExecutionNode(testWfNode2)

	testWfNode3 := new(testExecNode)
	testWfNode3.SetID("E3")
	testWfDefinition.AddExecutionNode(testWfNode3)

	testWfFrkNode := new(testForkNode)
	testWfFrkNode.SetID("F1")
	testWfDefinition.AddForkNode(testWfFrkNode, []WorkFlowNodeInterface{testWfNode1, testWfNode2})

	testWfJnNode := new(testJoinNode)
	testWfJnNode.SetID("J1")
	testWfDefinition.AddJoinNode(testWfJnNode)

	testWfDefinition.AddConnection(testWfNode2, testWfNode3)
	testWfDefinition.AddConnection(testWfNode1, testWfJnNode)
	testWfDefinition.AddConnection(testWfNode3, testWfJnNode)

	testWfDefinition.SetStartNode(testWfFrkNode)

	return testWfDefinition
}

/*
Test for Executing workflow with fork and join node
*/
func TestForkJoinNodeRun(t *testing.T) {
	testOrchestrator := new(Orchestrator)
	testWorkflowDefinition := createForkJoinTestWorkflowDefinition()
	testOrchestrator.Create(testWorkflowDefinition)

	testWorkFlowData := createTestWorkflowData()
	outputData := testOrchestrator.Start(testWorkFlowData)

	wfSate := outputData.GetWorkflowState()
	t.Log(wfSate)

	_, fFound := wfSate[tESTEXECUTIONNODENAME]
	if !fFound {
		t.Error("Failed execute Orchestrator with Fork Join Node")
	}
}

/*
Test for Orchestrator tostring
*/
func TestOrchestratorToString(t *testing.T) {
	testOrchestrator := new(Orchestrator)
	testWorkflowDefinition := createTestWorkflowDefinition()
	testWorkflowDefinition.AddExecutionNode(new(testExecNode))
	testOrchestrator.Create(testWorkflowDefinition)

	if testOrchestrator.String() == "" {
		t.Error("Empty string from orchestrator toString")
	}
}

/*
Test that Orchestrator implements Versionable GetInstance
*/
func TestOrchestratorIsVersionable(t *testing.T) {
	testOrchestrator := new(Orchestrator)
	testWorkflowDefinition := createTestWorkflowDefinition()
	testOrchestrator.Create(testWorkflowDefinition)

	if testOrchestrator.GetInstance() == nil {
		t.Error("Orchestrator is not versionable")
	}
}
