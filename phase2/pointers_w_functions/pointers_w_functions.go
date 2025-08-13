package main

import "fmt"

// Function that takes a value (copy)
func incrementValue(x int) {
	x++
	fmt.Printf("Inside incrementValue: %d\n", x)
}

// Function that takes a pointer (reference)
func incrementPointer(x *int) {
	*x++
	fmt.Printf("Inside incrementPointer: %d\n", *x)
}

// Function that returns a pointer
func createInt(value int) *int {
	x := value // Local variable
	fmt.Printf("Inside createInt, &x: %p\n", &x)
	return &x // Safe to return pointer to local variable in Go
}

// Pointer receiver method
type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c Counter) GetValue() int {
	return c.value
}

func main() {
	num := 10
	fmt.Printf("Original: %d, at: %p\n", num, &num)
	num++
	fmt.Printf("Original: %d, at: %p\n", num, &num)

	incrementValue(num)
	fmt.Printf("After incrementValue: %d\n", num) // Still 11

	incrementPointer(&num)
	fmt.Printf("After incrementPointer: %d\n", num) // Now 12

	// Using function that returns pointer
	intPtr := createInt(42)
	fmt.Printf("Created int: %d, at: %p\n", *intPtr, intPtr)

	// Pointer receiver methods
	counter := Counter{value: 0}
	fmt.Printf("Initial counter: %d\n", counter.GetValue())

	counter.Increment() // Go automatically converts to (&counter).Increment()
	fmt.Printf("After increment: %d\n", counter.GetValue())
}
