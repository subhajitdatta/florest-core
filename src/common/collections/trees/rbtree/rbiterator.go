package rbtree

import (
	"github.com/jabong/florest-core/src/common/collections"
)

// Iterator - A stateful iterator for red black tree
type Iterator struct {
	tree       *Tree
	node       *Node
	state      state
	nextCalled bool
	hasNext    bool
}

type state byte

const (
	begin, between, end state = 0, 1, 2
)

// HasNext method moves the iterator to the next element and returns true if there was a next
// element in the tree.
func (iterator *Iterator) HasNext() bool {
	if !iterator.nextCalled {
		return iterator.hasNext
	}
	if iterator.state == end {
		return iterator.goToEnd()
	}
	if iterator.state == begin {
		left := iterator.tree.Left()
		if left == nil {
			return iterator.goToEnd()
		}
		iterator.node = left
		return iterator.goInBetween()
	}
	if iterator.node.Right != nil {
		iterator.node = iterator.node.Right
		for iterator.node.Left != nil {
			iterator.node = iterator.node.Left
		}
		return iterator.goInBetween()
	}
	node := iterator.node
	for iterator.node.Parent != nil {
		iterator.node = iterator.node.Parent
		if iterator.tree.Comparator.Compare(node.Key, iterator.node.Key) <= 0 {
			return iterator.goInBetween()
		}
	}
	return false
}

// Next method returns the next element entry if it exists
func (iterator *Iterator) Next() *collections.Entry {
	if iterator.HasNext() {
		iterator.nextCalled = true
		return collections.NewEntry(iterator.node.Key, iterator.node.Value)
	}
	return nil
}

// Reset method resets the iterator to its initial state
// Call Next() to fetch the first element if any.
func (iterator *Iterator) Reset() {
	iterator.node = nil
	iterator.state = begin
	iterator.nextCalled = true
}

func (iterator *Iterator) goInBetween() bool {
	iterator.state = between
	iterator.hasNext = true
	iterator.nextCalled = false
	return true
}

func (iterator *Iterator) goToEnd() bool {
	iterator.node = nil
	iterator.state = end
	iterator.hasNext = false
	return false
}
