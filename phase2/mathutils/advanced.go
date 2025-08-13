package mathutils

import "math"

// Exported function
func Sqrt(n float64) float64 {
	return math.Sqrt(n)
}

// Exported function
func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// Package-level initialization
func init() {
	// This runs when the package is imported
	println("mathutils package initialized")
}
