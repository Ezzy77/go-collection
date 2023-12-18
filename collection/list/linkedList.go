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
	// list not empty
	if l.head != nil {
		nodeToAdd.next = l.head
		l.head = nodeToAdd
	}
	// when list is empty
	l.head = nodeToAdd
}

func (l *LinkedList[T]) InsertLast(v T) {
	nodeToAdd := &Node[T]{data: v, next: nil}
	// if the list is empty
	if l.head == nil {
		l.head = nodeToAdd
		return
	}
	// list not empty
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = nodeToAdd
}

func (l *LinkedList[T]) Find(v T) bool {
	if l.head == nil {
		return false
	}
	if v == l.head.data {
		return true
	}

	current := l.head
	for current.next != nil {
		if current.data == v {
			return true
		}
		current = current.next
	}
	return current.data == v
}

func (l *LinkedList[T]) DeleteFirst() bool {
	if l.head == nil {
		return false
	}

	return false
}

func (l *LinkedList[T]) DeleteLast() bool {
	if l.head == nil {
		return false
	}
	if l.head.next == nil {
		l.head = nil
		return true
	}
	current := l.head
	previous := l.head
	for current.next != nil {
		previous = current
		current = current.next
	}
	previous.next = nil
	return true
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
	fmt.Println("]")

}
