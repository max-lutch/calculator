package calculator_test

import (
	"calculator"
	"testing"
)

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 3
	got := calculator.Subtract(5, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 20, b: 2, want: 10},
		{a: -1, b: -1, want: 1},
		{a: 15, b: -3, want: -5},
	}
	for _, tc := range testCases {
		got := calculator.Divide(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}
