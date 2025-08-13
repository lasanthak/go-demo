package main

import "fmt"

// Function with parameters and return value
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Function with named return values
func calculate(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return // naked return
}

// Variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// Basic function call
	fmt.Printf("5 + 3 = %d\n", add(5, 3))

	// Function with error handling
	//quotient, err := divide(10, 0)
	quotient, err := divide(10, 7)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	// Multiple return values
	s, p := calculate(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", s, p)

	// Variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
}
