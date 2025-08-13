package mathutils

import "errors"

func Add(a, b int) int {
	return a + b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func IsEven(n int) bool {
	return n%2 == 0
}
