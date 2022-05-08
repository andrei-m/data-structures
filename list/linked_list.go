package list

import "errors"

type LinkedList[T any] struct {
	head *linkedListNode[T]
	size int
}

// Push adds an element to the front of the list
func (l *LinkedList[T]) Push(element T) {
	node := linkedListNode[T]{val: element, next: l.head}
	l.head = &node
	l.size++
}

var errPopFromEmptyList = errors.New("cannot pop from an empty list")

// Pop removes the first element from the list and returns it
func (l *LinkedList[T]) Pop() (T, error) {
	if l.head == nil {
		var empty T
		return empty, errPopFromEmptyList
	}
	toReturn := l.head.val
	l.head = l.head.next
	l.size--
	return toReturn, nil
}

// Size returns the number of elements in the list
func (l *LinkedList[T]) Size() int {
	return l.size
}

type linkedListNode[T any] struct {
	val  T
	next *linkedListNode[T]
}
