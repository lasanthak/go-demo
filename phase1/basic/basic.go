package main

import "fmt"

func main() {
	// Numeric types
	var smallInt int8 = 127
	var regularInt int = 42
	var bigInt int64 = 9223372036854775807

	var smallFloat float32 = 3.14
	var bigFloat float64 = 3.141592653589793

	// String and character types
	var message string = "Hello, Go!"
	var character rune = 'A' // rune is alias for int32
	var byteValue byte = 65  // byte is alias for uint8

	// Boolean
	var isReady bool = true

	// Complex numbers
	var complexNum complex64 = 1 + 2i
	var complexSqr complex64 = complexNum * complexNum

	fmt.Printf("Integers: %d, %d, %d\n", smallInt, regularInt, bigInt)
	fmt.Printf("Floats: %.2f, %.15f\n", smallFloat, bigFloat)
	fmt.Printf("String: %s, Char: %c, Byte: %c\n", message, character, byteValue)
	fmt.Printf("Boolean: %t, Complex: %v\n", isReady, complexNum)
	fmt.Printf("complexSqr: %T, %+v\n", complexSqr, complexSqr)
}
