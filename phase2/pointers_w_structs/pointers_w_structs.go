package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Birthday() {
	p.Age++
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s @ %p", p.Name, &p)
}

func updatePersonValue(p Person) {
	p.Age = 100 // Modifies copy
}

func updatePersonPointer(p *Person) {
	p.Age = 100 // Modifies original
}

func main() {
	// Creating struct instances
	person1 := Person{Name: "Alice", Age: 30}
	person2 := &Person{Name: "Bob", Age: 25} // Pointer to struct

	fmt.Printf("Person1: %+v\n", person1)
	fmt.Printf("Person2: %+v\n", *person2)

	fmt.Println(person1.Greet())
	fmt.Println(person1.Greet())

	// Method calls (Go handles pointer conversion automatically)
	person1.Birthday() // Converted to (&person1).Birthday()
	person2.Birthday() // Already a pointer

	fmt.Printf("After birthday - Person1: %+v\n", person1)
	fmt.Printf("After birthday - Person2: %+v\n", *person2)

	// Function calls with value vs pointer
	fmt.Printf("Before updates - Person1.Age: %d\n", person1.Age)

	updatePersonValue(person1)
	fmt.Printf("After updatePersonValue - Person1.Age: %d\n", person1.Age) // Unchanged

	updatePersonPointer(&person1)
	fmt.Printf("After updatePersonPointer - Person1.Age: %d\n", person1.Age) // Changed
}
