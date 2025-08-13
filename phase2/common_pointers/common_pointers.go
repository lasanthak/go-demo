package main

import (
	"fmt"

	m "github.com/lasanthak/go-demo/phase2/model"
)

// Optional values using pointers (like nullable types)
func findPerson(id int) *m.Person {
	if id == 1 {
		return &m.Person{Name: "Alice", Age: 30}
	}
	return nil // Not found
}

// Builder pattern with method chaining
type PersonBuilder struct {
	person *m.Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &m.Person{}}
}

func (pb *PersonBuilder) Name(name string) *PersonBuilder {
	pb.person.Name = name
	return pb
}

func (pb *PersonBuilder) Age(age int) *PersonBuilder {
	pb.person.Age = age
	return pb
}

func (pb *PersonBuilder) Build() *m.Person {
	return pb.person
}

// Linked list example
type Node struct {
	Value int
	Next  *Node
}

func (n *Node) Append(value int) *Node {
	current := n
	for current.Next != nil {
		current = current.Next
	}
	newNode := Node{Value: value}
	current.Next = &newNode
	return &newNode
}

func (n *Node) Print() {
	current := n
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Optional values
	person := findPerson(1)
	if person != nil {
		fmt.Printf("Found: %+v\n", *person)
	} else {
		fmt.Println("Person not found")
	}

	// Builder pattern
	builtPerson := NewPersonBuilder().
		Name("Charlie").
		Age(28).
		Build()
	fmt.Printf("Built person: %+v\n", *builtPerson)

	// Linked list
	head := &Node{Value: 1}
	head.
		Append(2).
		Append(3).
		Append(4)

	fmt.Print("Linked list: ")
	head.Print()
}
