package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic switch
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend!")
	case time.Monday:
		fmt.Println("Monday blues...")
	default:
		fmt.Println("Weekday")
	}

	// Switch with expressions
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 80:
		fmt.Println("Good")
	case score >= 70:
		fmt.Println("Average")
	default:
		fmt.Println("Needs improvement")
	}

	// Type switch
	var value interface{} = struct{ age int }{42}
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
