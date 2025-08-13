// Package mathutils provides basic and advanced mathematical operations.
//
// This package includes functions for arithmetic operations, mathematical
// calculations, and utility functions for numerical computations.
//
// Example usage:
//
//	result := mathutils.Add(5, 3)
//	fmt.Println(result) // Output: 8
//
//	calc := mathutils.Calculator{Name: "MyCalc"}
//	product := calc.Multiply(4, 5)
//	fmt.Println(product) // Output: 20
package mathutils

import "errors"

// Exported function (starts with capital letter)
func Add(a, b int) int {
	return a + b
}

// Exported function
func Subtract(a, b int) int {
	return a - b
}

// Private function (starts with lowercase letter)
func multiply(a, b int) int {
	return a * b
}

// Exported function that uses private function
func Square(n int) int {
	return multiply(n, n)
}

// Exported variable
var Pi = 3.14159

// Private variable
var version = "1.0.0"

// Exported constant
const MaxIterations = 1000

// Exported type
type Calculator struct {
	Name string
}

// Method on exported type
func (c Calculator) Multiply(a, b int) int {
	return multiply(a, b) // Can access private function within package
}

// Exported function with error
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
