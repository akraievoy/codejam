package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{7, []int16{2, 3, 5}})

	if out.index1 != 0 {
		t.Errorf("index1 is not correct, expected %d, returned %d", 0, out.index1)
	}

	if out.index2 != 2 {
		t.Errorf("index2 is not correct, expected %d, returned %d", 2, out.index2)
	}
}

func TestSimpleUnsorted(t *testing.T) {
	out := solve(In{7, []int16{1, 5, 3, 2}})

	if out.index1 != 1 {
		t.Errorf("index1 is not correct, expected %d, returned %d", 1, out.index1)
	}

	if out.index2 != 3 {
		t.Errorf("index2 is not correct, expected %d, returned %d", 3, out.index2)
	}
}

func TestSimpleUnsortedTake2(t *testing.T) {
	out := solve(In{7, []int16{120, 5, 3, 2}})

	if out.index1 != 1 {
		t.Errorf("index1 is not correct, expected %d, returned %d", 1, out.index1)
	}

	if out.index2 != 3 {
		t.Errorf("index2 is not correct, expected %d, returned %d", 3, out.index2)
	}
}
