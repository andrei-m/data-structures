package list

import "errors"

type LinkedList struct {
	head *linkedListNode
}

// Push adds an element to the front of the list
func (l *LinkedList) Push(element int) {
	node := linkedListNode{val: element, next: l.head}
	l.head = &node
}

var errPopFromEmptyList = errors.New("cannot pop from an empty list")

// Pop removes the first element from the list and returns it
func (l *LinkedList) Pop() (int, error) {
	if l.head == nil {
		return 0, errPopFromEmptyList
	}
	toReturn := l.head.val
	l.head = l.head.next
	return toReturn, nil
}

type linkedListNode struct {
	val  int
	next *linkedListNode
}
