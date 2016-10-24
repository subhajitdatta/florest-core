// Hash table and linked list implementation of the Set interface, with predictable iteration order.
//
// This implementation differs from HashSet in that it maintains a doubly-linked list running through
// all of its entries. This linked list defines the iteration ordering, which is the order in
// which elements were inserted into the set (insertion-order).
//
// Note that insertion order is not affected if an element is re-inserted into the set
//
// Structure is not thread safe.
package linkedhashset
