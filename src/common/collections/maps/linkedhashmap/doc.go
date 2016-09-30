// Package linkedhashmap - Hash table and linked list implementation of the Map interface, with predictable iteration order.
//
// This implementation differs from HashMap in that it maintains a doubly-linked list running through
// all of its entries. This linked list defines the iteration ordering, which is normally the order
// in which keys were inserted into the map (insertion-order).
//
// Note that insertion order is not affected if a key is re-inserted into the map
//
// Structure is not thread safe.
package linkedhashmap
