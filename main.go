package main

import (
	"fmt"

	"github.com/Ezzy77/go-collection/collection/list"
)

func main() {
	// person1 := Person{Name: "Elisio", Age: 24}
	// person2 := Person{Name: "Dionisio", Age: 30}
	// person3 := Person{Name: "Abel", Age: 10}

	//lists := list.LinkedList[int]{}
	dlist := list.DoublyLinkedList[int]{}
	dlist.InsertFirst(7)
	dlist.InsertFirst(9)
	dlist.InsertFirst(3)
	dlist.InsertFirst(1)

	fmt.Println(dlist.Find(7))

}

type Person struct {
	Name string
	Age  int
}
