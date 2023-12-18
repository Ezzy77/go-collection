package main

import (
	"fmt"

	"github.com/Ezzy77/go-collection/collection/list"
)

func main() {
	// person1 := Person{Name: "Elisio", Age: 24}
	// person2 := Person{Name: "Dionisio", Age: 30}
	// person3 := Person{Name: "Abel", Age: 10}

	lists := list.LinkedList[int]{}
	//personList := list.LinkedList[Person]{}

	lists.InsertLast(6)
	lists.InsertLast(3)
	lists.InsertLast(7)
	lists.InsertLast(9)

	fmt.Println(lists.Find(6))

	// personList.InsertFirst(person1)
	// personList.InsertFirst(person2)
	// personList.InsertFirst(person3)

	lists.Print()
	//personList.Print()

}

type Person struct {
	Name string
	Age  int
}
