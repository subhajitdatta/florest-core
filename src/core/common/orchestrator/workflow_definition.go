package orchestrator

import (
	"errors"
	"fmt"
	"github.com/jabong/florest-core/src/common/collections/stack"
)

/*
Work flow definition storage
*/
type WorkFlowDefinition struct {

	//Node of the Work Flow Definition
	nodes map[string]WorkFlowNodeInterface

	//The connections of the Work flow Definition
	edges map[string][]string

	//Mapping of the join node for the corresponding fork node
	//Every Fork Node should have a Join node
	joinFork map[string]string

	startNodeID string
}

/*
TODO: Create the workflow definition from configuration file
*/
func (d *WorkFlowDefinition) CreateFromConfig(filename string) error {
	return nil
}

/*
TODO: Create the workflow definition from configuration file
*/
func (d *WorkFlowDefinition) Create() {
	d.nodes = make(map[string]WorkFlowNodeInterface)
	d.edges = make(map[string][]string)
	d.joinFork = make(map[string]string)
}

/*
Set the start Node
*/
func (d *WorkFlowDefinition) SetStartNode(node WorkFlowNodeInterface) error {
	id, iderr := node.GetID()
	if iderr != nil {
		return iderr
	}

	_, found := d.nodes[id]
	if !found {
		errString := fmt.Sprintln("Node with provide Id: ", id, " is not added")
		return errors.New(errString)
	}

	d.startNodeID = id
	return nil

}

/*
Get the start node
*/
func (d *WorkFlowDefinition) GetStartNode() (*WorkFlowNodeInterface, error) {
	if d.startNodeID == "" {
		errString := fmt.Sprintf("Start Node is not set")
		return nil, errors.New(errString)
	}

	startNode := d.nodes[d.startNodeID]

	return &startNode, nil
}

/*
Add a execution node
*/
func (d *WorkFlowDefinition) AddExecutionNode(execNode WorkFlowExecuteNodeInterface) error {

	id, iderr := execNode.GetID()
	if iderr != nil {
		return iderr
	}

	_, found := d.nodes[id]
	if found {
		errString := fmt.Sprintln("Node with provide Id: ", id, " is already added")
		return errors.New(errString)
	}

	d.nodes[id] = execNode

	return nil
}

/*
Add a decision node
*/

func (d *WorkFlowDefinition) AddDecisionNode(decisionNode WorkFlowDecisionNodeInterface,
	yesNode WorkFlowNodeInterface, noNode WorkFlowNodeInterface) error {

	id, iderr := decisionNode.GetID()
	if iderr != nil {
		return iderr
	}

	yesNodeID, yesNodeIderr := yesNode.GetID()
	if yesNodeIderr != nil {
		return yesNodeIderr
	}

	noNodeID, noNodeIderr := noNode.GetID()
	if noNodeIderr != nil {
		return noNodeIderr
	}

	_, found := d.nodes[id]
	if found {
		errString := fmt.Sprintln("Node with provide Id: ", id, " is already added")
		return errors.New(errString)
	}

	d.nodes[id] = decisionNode
	d.nodes[yesNodeID] = yesNode
	d.nodes[noNodeID] = noNode

	d.edges[id] = []string{yesNodeID, noNodeID}

	return nil
}

/*
Add a Fork Node
*/
func (d *WorkFlowDefinition) AddForkNode(forkNode WorkFlowForkNodeInterface,
	forkNodes []WorkFlowNodeInterface) error {

	id, iderr := forkNode.GetID()
	if iderr != nil {
		return iderr
	}

	forkNodesIds, forkNodeIdserr := d.getNodeIds(forkNodes)
	if forkNodeIdserr != nil {
		return forkNodeIdserr
	}

	_, found := d.nodes[id]
	if found {
		errString := fmt.Sprintln("Node with provide Id: ", id, " is already added")
		return errors.New(errString)
	}

	d.nodes[id] = forkNode
	for index, forkNodeID := range forkNodesIds {
		d.nodes[forkNodeID] = forkNodes[index]
	}
	d.edges[id] = forkNodesIds

	return nil
}

func (d *WorkFlowDefinition) getNodeIds(nodes []WorkFlowNodeInterface) ([]string, error) {
	var ids []string

	for _, node := range nodes {
		id, iderr := node.GetID()
		if iderr != nil {
			return []string{}, iderr
		}
		ids = append(ids, id)
	}

	return ids, nil
}

/*
Add a Join Node
*/
func (d *WorkFlowDefinition) AddJoinNode(joinNode WorkFlowJoinNodeInterface) error {
	id, iderr := joinNode.GetID()
	if iderr != nil {
		return iderr
	}

	_, found := d.nodes[id]
	if found {
		errString := fmt.Sprintln("Node with provide Id: ", id, " is already added")
		return errors.New(errString)
	}

	d.nodes[id] = joinNode

	return nil
}

/*
Add connection between 2 nodes
*/
func (d *WorkFlowDefinition) AddConnection(fromNode WorkFlowNodeInterface,
	toNode WorkFlowNodeInterface) error {

	fromNodeID, fromNodeIDErr := fromNode.GetID()
	if fromNodeIDErr != nil {
		return fromNodeIDErr
	}

	toNodeID, toNodeiderr := toNode.GetID()
	if toNodeiderr != nil {
		return toNodeiderr
	}

	_, fromNodefound := d.nodes[fromNodeID]
	if !fromNodefound {
		errString := fmt.Sprintln("Node with provide Id: ", fromNodeID, " is not present. Add the node")
		return errors.New(errString)
	}

	_, toNodefound := d.nodes[toNodeID]
	if !toNodefound {
		errString := fmt.Sprintln("Node with provide Id: ", toNodeID, " is not present. Add the node")
		return errors.New(errString)
	}

	edgeList, edgeListfound := d.edges[fromNodeID]
	if !edgeListfound {
		d.edges[fromNodeID] = []string{toNodeID}
		return nil
	}

	edgeList = append(edgeList, toNodeID)
	d.edges[fromNodeID] = edgeList

	return nil

}

//Create mapping of the fork and corresponding join nodes
func (d *WorkFlowDefinition) createJoinForkMapping() error {

	sNode, serr := d.GetStartNode()
	if serr != nil {
		return serr
	}

	return d.createJoinForkMappingHelper(*sNode, d.startNodeID, &stack.Stack{})
}

func (d *WorkFlowDefinition) String() string {
	var res string
	res = fmt.Sprintf("nodes %v: \n", d.nodes)
	res = res + fmt.Sprintf("edges  %v: \n", d.edges)
	res = res + fmt.Sprintf("join-fork  %v: \n", d.joinFork)
	res = res + fmt.Sprintf("startNodeID  %v: \n", d.startNodeID)
	return res
}

func (d *WorkFlowDefinition) forkTypeNodeHelper(forkNode WorkFlowForkNodeInterface,
	forkNodeID string,
	stck *stack.Stack) error {

	stck.Push(forkNode)
	forkedEdges, found := d.edges[forkNodeID]
	if !found {
		return errors.New("No outgoing edges from the fork node")
	}
	for _, forkedNodeID := range forkedEdges {
		forkedNode, fNodeFound := d.nodes[forkedNodeID]
		if !fNodeFound {
			return errors.New("Forked Node with the node Id not found")
		}
		//clone the stack and send to all the forked nodes for this fork node
		fStck := stck.Clone()

		ferr := d.createJoinForkMappingHelper(forkedNode, forkedNodeID, fStck)
		if ferr != nil {
			return ferr
		}
	}
	return nil
}

func (d *WorkFlowDefinition) joinTypeNodeHelper(joinNode WorkFlowJoinNodeInterface,
	joinNodeID string,
	stck *stack.Stack) error {

	if stck.IsEmpty() {
		return errors.New("Join node does not have a corresponding fork node")
	}
	stckValue := stck.Pop()

	forkNode, ok := stckValue.(WorkFlowForkNodeInterface)
	if !ok {
		return errors.New("Incorrect Node type popped from stack")
	}
	forkNodeID, ferr := forkNode.GetID()
	if ferr != nil {
		return errors.New("Fork Node does not have ID set")
	}
	d.joinFork[forkNodeID] = joinNodeID
	edges, found := d.edges[joinNodeID]
	if !found || len(edges) == 0 {
		return nil
	}
	nextNodeID := edges[0]
	nextNode := d.nodes[nextNodeID]
	return d.createJoinForkMappingHelper(nextNode, nextNodeID, stck)
}

func (d *WorkFlowDefinition) executeTypeNodeHelper(execNode WorkFlowExecuteNodeInterface,
	execNodeID string,
	stck *stack.Stack) error {

	edges, found := d.edges[execNodeID]
	if !found || len(edges) == 0 {
		return nil
	}
	nextNodeID := edges[0]
	nextNode := d.nodes[nextNodeID]
	return d.createJoinForkMappingHelper(nextNode, nextNodeID, stck)
}

func (d *WorkFlowDefinition) decisionTypeNodeHelper(decisionNode WorkFlowDecisionNodeInterface,
	decisionNodeID string,
	stck *stack.Stack) error {

	edges, found := d.edges[decisionNodeID]
	if !found || len(edges) == 0 {
		return nil
	}

	for _, nextNodeID := range edges {
		nextNode, nextNodeFound := d.nodes[nextNodeID]
		if !nextNodeFound {
			return errors.New("Node with the node Id not found")
		}
		//clone the stack and send to all the nodes for this decision node
		dStck := stck.Clone()

		derr := d.createJoinForkMappingHelper(nextNode, nextNodeID, dStck)
		if derr != nil {
			return derr
		}
	}
	return nil
}

func (d *WorkFlowDefinition) createJoinForkMappingHelper(currNode WorkFlowNodeInterface,
	currNodeID string,
	stck *stack.Stack) error {

	//Error condition
	if (currNode == nil || currNodeID == "") && !stck.IsEmpty() {
		return errors.New("Error in creating Fork Join Node mapping")
	}

	//Recursion termination
	if currNode == nil || currNodeID == "" {
		return nil
	}

	if forkNode, ok := currNode.(WorkFlowForkNodeInterface); ok {
		//Current Node is a forked node
		return d.forkTypeNodeHelper(forkNode, currNodeID, stck)
	}

	if joinNode, ok := currNode.(WorkFlowJoinNodeInterface); ok {
		//Current Node is a join node
		return d.joinTypeNodeHelper(joinNode, currNodeID, stck)
	}

	if execNode, ok := currNode.(WorkFlowExecuteNodeInterface); ok {
		//Current Node is an execution node
		return d.executeTypeNodeHelper(execNode, currNodeID, stck)
	}

	if decisionNode, ok := currNode.(WorkFlowDecisionNodeInterface); ok {
		//Current Node is a decision node
		return d.decisionTypeNodeHelper(decisionNode, currNodeID, stck)
	}

	return errors.New("Unknown Node type")
}
