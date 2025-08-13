package main

import "fmt"

// Basic struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Struct with embedded struct
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Employee struct {
	Person   // Embedded struct
	ID       int
	Position string
	Salary   float64
	Address  Address
}

// Method on struct
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

func main() {
	// Creating structs
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john.doe@example.com",
	}

	// Anonymous struct
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}

	// Struct with embedded struct
	employee := Employee{
		Person: Person{
			FirstName: "Jane",
			LastName:  "Smith",
			Age:       28,
			Email:     "jane.smith@company.com",
		},
		ID:       1001,
		Position: "Software Engineer",
		Salary:   75000.0,
		Address: Address{
			Street:  "123 Main St",
			City:    "Austin",
			State:   "TX",
			ZipCode: "78701",
		},
	}

	fmt.Printf("Person: %+v\n", person1)
	fmt.Printf("Config: %+v\n", config)
	// The Employee struct embeds the Person struct, which means Employee inherits all the fields and methods of Person.
	// Because of this embedding, you can call employee.FullName() directly, even though FullName is defined on Person. This is called method promotion in Go.
	fmt.Printf("%s works as %s\n", employee.FullName(), employee.Position)
	fmt.Printf("%s works as %s\n", employee.Person.FullName(), employee.Position)

	// Using methods
	fmt.Printf("Full name: %s\n", person1.FullName())
	person1.UpdateAge(31)
	fmt.Printf("Updated age: %d\n", person1.Age)
}
