package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 9
	got := Multiply(3, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
