package linkedlist

// Represents a singly linked node.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}
