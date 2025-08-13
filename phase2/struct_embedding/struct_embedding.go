package main

import "fmt"

// Base types
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, I'm %s\n", p.Name)
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

type Address struct {
	Street  string
	City    string
	Country string
}

func (a Address) GetLocation() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

// Embedding examples
type Employee struct {
	Person     // Embedded struct - promoted fields and methods
	EmployeeID string
	Department string
	Salary     float64
}

func (e Employee) GetDetails() string {
	return fmt.Sprintf("Employee %s: %s in %s department",
		e.EmployeeID, e.GetInfo(), e.Department)
}

type Customer struct {
	Person     // Embedded Person
	Address    // Embedded Address
	CustomerID string
	Orders     []Order
}

type Order struct {
	ID    string
	Total float64
}

// Multiple embedding
type Manager struct {
	Employee // Embedded Employee (which embeds Person)
	TeamSize int
	Reports  []Employee
}

func (m Manager) GetTeamInfo() string {
	return fmt.Sprintf("Manager %s leads a team of %d people", m.Name, m.TeamSize)
}

// Interface embedding
type Animal interface {
	Speak() string
	Move() string
}

type Mammal interface {
	Animal // Embedded interface
	GiveBirth()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running"
}

func (d Dog) GiveBirth() {
	fmt.Printf("%s is giving birth to puppies\n", d.Name)
}

func main() {
	// Basic embedding
	emp := Employee{
		Person: Person{
			Name: "Alice Johnson",
			Age:  30,
		},
		EmployeeID: "EMP001",
		Department: "Engineering",
		Salary:     75000,
	}

	// Accessing embedded fields and methods
	fmt.Printf("Name: %s\n", emp.Name) // Promoted field
	fmt.Printf("Age: %d\n", emp.Age)   // Promoted field
	emp.Greet()                        // Promoted method
	fmt.Println(emp.GetDetails())      // Own method

	// Multiple embedding
	customer := Customer{
		Person: Person{
			Name: "Bob Smith",
			Age:  45,
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "Austin",
			Country: "USA",
		},
		CustomerID: "CUST001",
		Orders: []Order{
			{ID: "ORD001", Total: 99.99},
			{ID: "ORD002", Total: 149.50},
		},
	}

	fmt.Printf("Customer: %s\n", customer.Name)          // From Person
	fmt.Printf("Location: %s\n", customer.GetLocation()) // From Address
	customer.Greet()                                     // From Person

	// Nested embedding
	manager := Manager{
		Employee: Employee{
			Person: Person{
				Name: "Charlie Brown",
				Age:  40,
			},
			EmployeeID: "MGR001",
			Department: "Engineering",
			Salary:     95000,
		},
		TeamSize: 8,
	}

	fmt.Printf("Manager name: %s\n", manager.Name) // From Person via Employee
	fmt.Println(manager.GetTeamInfo())             // Own method
	fmt.Println(manager.GetDetails())              // From Employee

	// Interface embedding
	var mammal Mammal = Dog{Name: "Buddy"}
	fmt.Printf("Dog says: %s\n", mammal.Speak())
	fmt.Printf("Dog is: %s\n", mammal.Move())
	mammal.GiveBirth()
}
