package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	// FIXME simple test case
	out := solve(In{[]int16{2, 3, 5}})

	if out.sum != 10 {
		t.Errorf("sum is not correct, expected %d, returned %d", 10, out.sum)
	}
}
