package mathutils

import (
	"math"
	"testing"
)

const delta = 1e-9

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	// Test successful division
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	if math.Abs(result-5.0) > delta {
		t.Errorf("Divide(10, 2) = %.2f; want 5.00", result)
	}

	// Test division by zero
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) should return error")
	}
}

func TestIsEven(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-2, true},
		{-3, false},
	}

	for _, tc := range testCases {
		result := IsEven(tc.input)
		if result != tc.expected {
			t.Errorf("IsEven(%d) = %t; want %t", tc.input, result, tc.expected)
		}
	}
}
