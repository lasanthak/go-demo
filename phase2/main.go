// File: main.go
package main

import (
	"fmt"

	"github.com/lasanthak/go-demo/phase2/mathutils"        // Import local package
	mutils "github.com/lasanthak/go-demo/phase2/mathutils" // Import with alias
)

func main() {
	// Using imported functions
	sum := mathutils.Add(5, 3)
	diff := mathutils.Subtract(10, 4)
	square := mathutils.Square(7)

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Square: %d\n", square)

	// Using imported constants and variables
	fmt.Printf("Pi: %.5f\n", mathutils.Pi)
	fmt.Printf("Max iterations: %d\n", mathutils.MaxIterations)

	// Using imported types
	calc := mathutils.Calculator{Name: "MyCalculator"}
	product := calc.Multiply(6, 7)
	fmt.Printf("Product: %d\n", product)

	// Error handling
	quotient, err := mathutils.Divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Quotient: %.2f\n", quotient)
	}

	// Using aliased import
	sqrt := mutils.Sqrt(118)
	power := mutils.Power(2, 8)
	fmt.Printf("Square root: %.2f\n", sqrt)
	fmt.Printf("Power: %.2f\n", power)
}
