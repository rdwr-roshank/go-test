package main

import (
	"testing"
)

// TestAdd tests the Add function with various inputs
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zero", 0, 0, 0},
		{"large numbers", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestGreet tests the Greet function
func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"basic greeting", "Alice", "Hello, Alice!"},
		{"another greeting", "Bob", "Hello, Bob!"},
		{"empty name", "", "Hello, !"},
		{"name with space", "John Doe", "Hello, John Doe!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet(tt.input)
			if got != tt.expected {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// BenchmarkAdd benchmarks the Add function
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

// BenchmarkGreet benchmarks the Greet function
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("World")
	}
}

// TestDivide tests the Divide function including error cases
func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a         int
		b         int
		expected  int
		wantError bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division result", 20, 4, 5, false},
		{"divide by zero", 10, 0, 0, true},
		{"negative division", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantError {
				t.Errorf("Divide(%d, %d) error = %v, wantError %v", tt.a, tt.b, err, tt.wantError)
				return
			}
			if err == nil && got != tt.expected {
				t.Errorf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
