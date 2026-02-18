package calculator

import "testing"

// Example 1: Basic test
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

// Example 2: Multiple test cases in one function
func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Subtract(5, 3) = %d; expected 2", result)
	}

	result = Subtract(10, 15)
	if result != -5 {
		t.Errorf("Subtract(10, 15) = %d; expected -5", result)
	}
}

// Example 3: Table-driven tests (best practice!)
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 3, 4, 12},
		{"with zero", 5, 0, 0},
		{"negative numbers", -2, 3, -6},
		{"both negative", -2, -3, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; expected %d",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Example 4: Testing edge cases
func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"normal division", 10, 2, 5},
		{"divide by zero", 10, 0, 0},
		{"negative result", -10, 2, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Divide(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; expected %d",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Example 5: Using t.Fatal (stops test immediately)
func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Fatal("2 should be even") // Stops here if fails
	}

	if IsEven(3) {
		t.Error("3 should not be even") // Continues testing
	}

	if !IsEven(0) {
		t.Error("0 should be even")
	}
}

// Example 6: Benchmark test (measures performance)
func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

// Example 7: Parallel tests (runs concurrently)
func TestAddParallel(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"test1", 1, 1, 2},
		{"test2", 2, 2, 4},
		{"test3", 3, 3, 6},
	}

	for _, tt := range tests {
		tt := tt // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Run this test in parallel
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("got %d, expected %d", result, tt.expected)
			}
		})
	}
}

// Example 8: Helper function (marked with t.Helper())
func assertEqual(t *testing.T, got, expected int) {
	t.Helper() // Makes error messages point to the caller
	if got != expected {
		t.Errorf("got %d, expected %d", got, expected)
	}
}

func TestWithHelper(t *testing.T) {
	result := Add(5, 7)
	assertEqual(t, result, 12)
}
