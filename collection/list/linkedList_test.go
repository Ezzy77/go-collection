package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string
	Age  int
}

func TestInsertFirstInt(t *testing.T) {
	testingList := LinkedList[int]{}

	testingList.InsertFirst(8)
	testingList.InsertFirst(9)
	testingList.InsertFirst(1)

	assert.Equal(t, 1, testingList.head.data, "First element should be 1")
	assert.Equal(t, 9, testingList.head.next.data, "First element should be 9")
	assert.Equal(t, 8, testingList.head.next.next.data, "First element should be 8")
	assert.Nil(t, testingList.head.next.next.next, "Next of the third element should be nil")

}

func TestInsertFirstStruct(t *testing.T) {
	person1 := Person{Name: "Bill", Age: 70}
	person2 := Person{Name: "David", Age: 17}
	person3 := Person{Name: "John", Age: 40}

	testingList := LinkedList[Person]{}

	testingList.InsertFirst(person1)
	testingList.InsertFirst(person2)
	testingList.InsertFirst(person3)

	assert.Equal(t, person3, testingList.head.data, "First element should be %v", person3)
	assert.Equal(t, person2, testingList.head.next.data, "Second element should be %v", person2)
	assert.Equal(t, person1, testingList.head.next.next.data, "Third element should be %v", person1)
	assert.Nil(t, testingList.head.next.next.next, "Next of the third element should be nil")

}
