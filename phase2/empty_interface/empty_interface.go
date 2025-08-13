package main

import (
	"fmt"
	"math"

	m "github.com/lasanthak/go-demo/phase2/model"
)

// Empty interface can hold any type
func PrintAnything(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// Type assertions and type switches
func ProcessValue(value interface{}) {
	// Type assertion
	if str, ok := value.(string); ok {
		fmt.Printf("String value: %s (length: %d)\n", str, len(str))
		return
	}

	// Type switch
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d (squared: %d)\n", v, v*v)
	case float64:
		fmt.Printf("Float: %.2f (sqrt: %.2f)\n", v, math.Sqrt(v))
	case bool:
		fmt.Printf("Boolean: %t (negated: %t)\n", v, !v)
	case []int:
		fmt.Printf("Int slice: %v (sum: %d)\n", v, sum(v))
	case m.Person:
		fmt.Printf("Person: %s is %d years old\n", v.Name, v.Age)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	values := []interface{}{
		42,
		3.14159,
		"Hello, Go!",
		true,
		[]int{1, 2, 3, 4, 5},
		m.Person{Name: "Alice", Age: 30},
		map[string]int{"a": 1, "b": 2},
	}

	for _, value := range values {
		PrintAnything(value)
		ProcessValue(value)
		fmt.Println("---")
	}
}
