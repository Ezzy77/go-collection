package list

import "fmt"

type DNode[T comparable] struct {
	data T
	next *DNode[T]
	prev *DNode[T]
}

type DoublyLinkedList[T comparable] struct {
	head *DNode[T]
	tail *DNode[T]
}

func (dl *DoublyLinkedList[T]) InsertFirst(v T) {
	nodeToAdd := &DNode[T]{data: v, prev: nil, next: nil}
	if dl.head == nil {
		dl.head = nodeToAdd
		dl.tail = nodeToAdd
		return
	}
	nodeToAdd.next = dl.head
	dl.head.prev = nodeToAdd
	dl.head = nodeToAdd
}

func (dl *DoublyLinkedList[T]) InsertLast(v T) {
	nodeToAdd := &DNode[T]{data: v, prev: nil, next: nil}
	if dl.head == nil {
		dl.head = nodeToAdd
		dl.tail = nodeToAdd
		return
	}
	nodeToAdd.prev = dl.tail
	dl.tail.next = nodeToAdd
	dl.tail = nodeToAdd
}

func (dl DoublyLinkedList[T]) Find(v T) bool {
	if dl.head == nil {
		return false
	}
	if v == dl.head.data || v == dl.tail.data {
		return true
	}
	current := dl.head
	for current.next != nil {
		if current.data == v {
			return true
		}
		current = current.next
	}
	return current.data == v
}

func (dl *DoublyLinkedList[T]) DeleteFirst() bool {
	if dl.head == nil {
		return false
	}
	nodeToDelete := dl.head
	if dl.head.next != nil {
		dl.head.next.prev = nil
	}
	dl.head = dl.head.next

	nodeToDelete.next = nil
	nodeToDelete.prev = nil
	return true
}

func (dl *DoublyLinkedList[T]) DeleteLast() bool {
	if dl.head == nil {
		return false
	}
	if dl.head == dl.tail {
		dl.head = nil
		dl.tail = nil
		return true
	}
	dl.tail = dl.tail.prev
	dl.tail.next = nil
	return true
}

// display all items in the list
func (dl *DoublyLinkedList[T]) Print() {
	current := dl.head
	fmt.Print("[")
	for current != nil {
		fmt.Printf("%v = ", current.data)
		current = current.next

	}
	fmt.Print("nil")
	fmt.Println("]")

}
