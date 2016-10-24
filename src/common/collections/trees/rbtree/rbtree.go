package rbtree

import (
	_ "fmt"
	"github.com/jabong/florest-core/src/common/collections"
	"github.com/jabong/florest-core/src/common/collections/utils"
)

type color bool

const (
	black, red color = true, false
)

// Node is a single element within the tree
type Node struct {
	Key    interface{}
	Value  interface{}
	color  color
	Left   *Node
	Right  *Node
	Parent *Node
}

func (node *Node) grandparent() *Node {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *Node) uncle() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (node *Node) rightmostChild() *Node {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (node *Node) isLeftChild() bool {
	if node == node.Parent.Left {
		return true
	}
	return false
}

func (node *Node) isRightChild() bool {
	if node == node.Parent.Right {
		return true
	}
	return false
}

// Tree represents the red-black tree
type Tree struct {
	Root       *Node
	size       int
	Comparator collections.Comparator
}

// New instantiates a red-black tree. Comparator passed is used to compare the elements of the tree
func New(comparator collections.Comparator) *Tree {
	return &Tree{Comparator: comparator}
}

// Put methods inserts a new node into the Tree.
// Key should adhere to the comparator's type assertion, otherwise it will panic.
func (t *Tree) Put(key interface{}, value interface{}) {
	newNode := &Node{Key: key, Value: value, color: red}
	if t.Root == nil {
		t.Root = newNode
	} else {
		node := t.Root
		loop := true
		for loop {
			compare := t.Comparator.Compare(key, node.Key)
			switch {
			// 0 means key already exists. So, just updating the value of it
			case compare == 0:
				node.Key = key
				node.Value = value
				return
			// -1 means key smaller than the root. Should go left.
			case compare < 0:
				if node.Left == nil {
					node.Left = newNode
					loop = false
				} else {
					node = node.Left
				}
			// +1 means key bigger than the root. Should go right.
			case compare > 0:
				if node.Right == nil {
					node.Right = newNode
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		newNode.Parent = node
	}
	t.insertRBCase1(newNode)
	t.size++
}

// Get method returns value of the key, true if key found. If key not found, returns nil, false
// Key should adhere to the comparator's type assertion, otherwise it will panic.
func (t *Tree) Get(key interface{}) (value interface{}, found bool) {
	node := t.search(key)
	if node != nil {
		return node.Value, true
	}
	return nil, false
}

// Remove method removes the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise it will panic.
func (t *Tree) Remove(key interface{}) {
	var child *Node
	// Get the node to be deleted
	node := t.search(key)
	if node == nil {
		return
	}
	// If node has both the children
	if node.Left != nil && node.Right != nil {
		predecessor := node.Left.rightmostChild()
		node.Key = predecessor.Key
		node.Value = predecessor.Value
		//  Now the node to be deleted becomes the predecessor
		node = predecessor
	}
	// If node has only one children
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == black {
			node.color = getNodeColor(child)
			// Handle delete cases otherwise violates black height rule
			t.deleteRBCase1(node)
		}
		// Delete the node by replacing it with it's child
		t.replaceNode(node, child)
		// If the node that was deleted is a root node
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
	t.size--
}

// Keys method returns all keys in-order (Ascending)
func (t *Tree) Keys() []interface{} {
	keys := make([]interface{}, t.size)
	it := t.Iterator()
	for i := 0; it.HasNext(); i++ {
		keys[i] = it.Next().GetKey()
	}
	return keys
}

// Values method returns all values in-order based on the key (Ascending).
func (t *Tree) Values() []interface{} {
	values := make([]interface{}, t.size)
	it := t.Iterator()
	for i := 0; it.HasNext(); i++ {
		values[i] = it.Next().GetValue()
	}
	return values
}

// Empty method returns true if the tree does not contain any nodes
func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

// Size method returns the number of nodes in the tree.
func (t *Tree) Size() int {
	return t.size
}

// Clear method removes all nodes from the tree.
func (t *Tree) Clear() {
	t.Root = nil
	t.size = 0
}

// Contains method returns true if all the given keys are present in the tree
func (t *Tree) Contains(keys ...interface{}) bool {
	keysInTree := t.Keys()
	keysMap := utils.ConvertArrayToMap(keysInTree)
	for _, key := range keys {
		if _, found := keysMap[key]; !found {
			return false
		}
	}
	return true
}

// Left method returns the left-most (min) node or nil if tree is empty.
func (t *Tree) Left() *Node {
	current := t.Root

	if current == nil {
		return nil
	}
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Right method returns the right-most (max) node or nil if tree is empty.
func (t *Tree) Right() *Node {
	current := t.Root

	if current == nil {
		return nil
	}
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// Iterator returns a stateful iterator used for iterating over all the entries of the tree
func (t *Tree) Iterator() collections.Iterator {
	return &Iterator{tree: t, node: nil, state: begin, nextCalled: true, hasNext: false}
}

// GetComparator returns the comparator associated with this tree
func (t *Tree) GetComparator() collections.Comparator {
	return t.Comparator
}

// Root node case -> Root node must be a black node
func (t *Tree) insertRBCase1(node *Node) {
	// If it's a root node
	if node.Parent == nil {
		node.color = black
	} else {
		t.insertRBCase2(node)
	}
}

// Node color case 1 -> Black node can have any colored children.
func (t *Tree) insertRBCase2(node *Node) {
	if getNodeColor(node.Parent) == black {
		return
	}
	t.insertRBCase3(node)
}

// Node color case 2 -> Red nodes must have only black children.
func (t *Tree) insertRBCase3(node *Node) {
	uncle := node.uncle()
	// If node uncle is present and if it's red, flip node color wherever required
	// traversing from the inserted node till the root node
	if getNodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandparent().color = red
		t.insertRBCase1(node.grandparent())
	} else {
		t.insertRBCase4(node)
	}
}

// Black height case -> Root to leaf must have same number of black nodes.
// Perform rotations to keep it balanced (For alternate/cross line nodes)
func (t *Tree) insertRBCase4(node *Node) {
	if node.isRightChild() && node.Parent.isLeftChild() {
		t.rotateLeft(node.Parent)
		node = node.Left
	} else if node.isLeftChild() && node.Parent.isRightChild() {
		t.rotateRight(node.Parent)
		node = node.Right
	}
	t.insertRBCase5(node)
}

// Black height case -> Root to leaf must have same number of black nodes.
// Perform rotations to keep it balanced (For Straight line nodes)
func (t *Tree) insertRBCase5(node *Node) {
	node.Parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node.isLeftChild() && node.Parent.isLeftChild() {
		t.rotateRight(grandparent)
	} else if node.isRightChild() && node.Parent.isRightChild() {
		t.rotateLeft(grandparent)
	}
}

// Root Node Case -> Just return.
func (t *Tree) deleteRBCase1(node *Node) {
	if node.Parent == nil {
		return
	}
	t.deleteRBCase2(node)
}

// If Sibling node is red
func (t *Tree) deleteRBCase2(node *Node) {
	sibling := node.sibling()
	if getNodeColor(sibling) == red {
		node.Parent.color = red
		sibling.color = black
		if node == node.Parent.Left {
			t.rotateLeft(node.Parent)
		} else {
			t.rotateRight(node.Parent)
		}
	}
	t.deleteRBCase3(node)
}

// If Parent. Sibling and its children are black
func (t *Tree) deleteRBCase3(node *Node) {
	sibling := node.sibling()
	if getNodeColor(node.Parent) == black &&
		getNodeColor(sibling) == black &&
		getNodeColor(sibling.Left) == black &&
		getNodeColor(sibling.Right) == black {
		sibling.color = red
		t.deleteRBCase1(node.Parent)
	} else {
		t.deleteRBCase4(node)
	}
}

//If Parent is red and Sibling & its children are black
func (t *Tree) deleteRBCase4(node *Node) {
	sibling := node.sibling()
	if getNodeColor(node.Parent) == red &&
		getNodeColor(sibling) == black &&
		getNodeColor(sibling.Left) == black &&
		getNodeColor(sibling.Right) == black {
		sibling.color = red
		node.Parent.color = black
	} else {
		t.deleteRBCase5(node)
	}
}

// If any one child of sibling is red
func (t *Tree) deleteRBCase5(node *Node) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		getNodeColor(sibling) == black &&
		getNodeColor(sibling.Left) == red &&
		getNodeColor(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		t.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		getNodeColor(sibling) == black &&
		getNodeColor(sibling.Right) == red &&
		getNodeColor(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		t.rotateLeft(sibling)
	}
	t.deleteRBCase6(node)
}

// If any one child of sibling is red
func (t *Tree) deleteRBCase6(node *Node) {
	sibling := node.sibling()
	sibling.color = getNodeColor(node.Parent)
	node.Parent.color = black
	if node.isLeftChild() && getNodeColor(sibling.Right) == red {
		sibling.Right.color = black
		t.rotateLeft(node.Parent)
	} else if getNodeColor(sibling.Left) == red {
		sibling.Left.color = black
		t.rotateRight(node.Parent)
	}
}

// Get the node color, return black if node is nil
func getNodeColor(node *Node) color {
	if node == nil {
		return black
	}
	return node.color
}

// Rotate the subtree towards left pointing at the given node
func (t *Tree) rotateLeft(node *Node) {
	right := node.Right
	t.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

// Rotate the subtree towards right pointing at the given node
func (t *Tree) rotateRight(node *Node) {
	left := node.Left
	t.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

// Searches the given key in the tree
func (t *Tree) search(key interface{}) *Node {
	node := t.Root
	for node != nil {
		compare := t.Comparator.Compare(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
}

// Replaces the oldNode with the newNode
func (t *Tree) replaceNode(oldNode *Node, newNode *Node) {
	if oldNode.Parent == nil {
		t.Root = newNode
	} else {
		if oldNode.isLeftChild() {
			oldNode.Parent.Left = newNode
		} else {
			oldNode.Parent.Right = newNode
		}
	}
	if newNode != nil {
		newNode.Parent = oldNode.Parent
	}
}
