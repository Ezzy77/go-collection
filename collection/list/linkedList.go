package list

import "fmt"

// Node represent a single item in the list
type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

// insert item in the beginning of the list
func (l *LinkedList[T]) InsertFirst(v T) {
	nodeToAdd := &Node[T]{data: v, next: nil}
	if l.head != nil {
		nodeToAdd.next = l.head
		l.head = nodeToAdd
	}
	l.head = nodeToAdd
}

// display all items in the list
func (l *LinkedList[T]) Print() {
	current := l.head
	fmt.Print("[")
	for current != nil {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}
	fmt.Print("nil")
	fmt.Print("]")
}
