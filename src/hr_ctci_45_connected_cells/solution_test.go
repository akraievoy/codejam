package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{2, 2, [][]bool{{true, false}, {false, true}}})

	if out.largest != 2 {
		t.Errorf("sum is not correct, expected %d, returned %d", 2, out.largest)
	}
}
