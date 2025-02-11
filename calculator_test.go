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
