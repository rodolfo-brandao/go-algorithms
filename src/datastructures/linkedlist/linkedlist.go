package linkedlist

import (
	"fmt"
	"strings"
)

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

// Initializes a new instance of the list.
//
// Returns:
//   - Its pointer.
func (l *LinkedList[T]) New() {
	l.Head = nil
	l.Tail = nil
}

// Adds an item to the beginning of the list.
//
// Parameters:
//   - value[any]: The item to be added to the beginning of the list.
func (l *LinkedList[T]) Prepend(value T) {
	node := &Node[T]{Value: value, Next: l.Head}
	l.Head = node

	if l.Tail == nil {
		l.Tail = node
	}
}

// Adds an item to the end of the list.
//
// Parameters:
//   - value[any]: The item to be added to the end of the list.
func (l *LinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = node
}

// Prints every item on the list.
func (l *LinkedList[T]) String() string {
	current := l.Head
	items := []string{}

	for current != nil {
		itemStr := fmt.Sprintf("%v", current.Value)
		items = append(items, itemStr)
		current = current.Next
	}

	return strings.Join(items, ", ")
}
