package mathutils

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

// Benchmark test
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Add(1000, 2000)
	}
}

// Table-driven tests
func TestAddTable(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"with zero", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Setup and teardown
func TestMain(m *testing.M) {
	// Setup
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Int()

	// Run tests
	code := m.Run()

	// Teardown
	// cleanup code here

	os.Exit(code)
}

// Helper functions
func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %v; want %v", actual, expected)
	}
}

func TestWithHelper(t *testing.T) {
	result := Add(2, 3)
	assertEqual(t, result, 5)
}
