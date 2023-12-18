package main

import "github.com/Ezzy77/go-collection/collection/list"

func main() {
	person1 := Person{Name: "Elisio", Age: 24}
	person2 := Person{Name: "Dionisio", Age: 30}
	person3 := Person{Name: "Abel", Age: 10}

	lists := list.LinkedList[int]{}
	personList := list.LinkedList[Person]{}

	lists.InsertFirst(6)
	lists.InsertFirst(3)
	lists.InsertFirst(7)

	personList.InsertFirst(person1)
	personList.InsertFirst(person2)
	personList.InsertFirst(person3)

	lists.Print()
	personList.Print()

}

type Person struct {
	Name string
	Age  int
}
