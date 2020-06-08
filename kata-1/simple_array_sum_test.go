package main

import "testing"

func TestSimpleArraySum(t *testing.T) {
	sum := SimpleArraySum([]int32{1, 2, 3, 4, 5})
	if sum != 15 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 15)
	}
}
