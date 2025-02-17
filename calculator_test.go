package calculator_test

import (
	"calculator" // This should match the folder name where calculator.go is loc
	"errors"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 2, b: 1, want: 1},
		{a: 5, b: -4, want: 9},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: -1, b: -1, want: 1},
		{a: 5, b: 0, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

// Function sqrt takes one number and returns it's square root or error if invalid input (<0)
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot take square root of a negative number")
	}
	return math.Sqrt(a), nil
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Errorf("Divide error for invalid input, got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		input, want, tolerance float64
	}{
		{input: 4, want: 2, tolerance: 0.0001},   // sqrt(4) = 2
		{input: 9, want: 3, tolerance: 0.0001},   // sqrt(9) = 3
		{input: 2, want: 1.414, tolerance: 0.01}, // sqrt(2) ≈ 1.414
	}

	for _, tc := range testCases {
		got, err := Sqrt(tc.input)
		if err != nil {
			t.Fatalf("Sqrt(%f) returned error: %v", tc.input, err)
		}
		if !closeEnough(got, tc.want, tc.tolerance) {
			t.Errorf("Sqrt(%f) = %f, want %f (tolerance %f)",
				tc.input, got, tc.want, tc.tolerance)
		}
	}
}
