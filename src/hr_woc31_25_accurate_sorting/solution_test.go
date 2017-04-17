package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{[]int32{2, 3, 5}})

	if !out.sortable {
		t.Errorf("sortable is not correct, expected %v, returned %v", true, out.sortable)
	}
}

func TestSimple2(t *testing.T) {
	out := solve(In{[]int32{3, 2, 5}})

	if !out.sortable {
		t.Errorf("sortable is not correct, expected %v, returned %v", true, out.sortable)
	}
}

func TestSimple3(t *testing.T) {
	out := solve(In{[]int32{1, 2, 0, 4, 3}})

	if out.sortable {
		t.Errorf("sortable is not correct, expected %v, returned %v", false, out.sortable)
	}
}

func TestLarge(t *testing.T) {
	nums := make([]int32, 100000)
	for i := range nums {
		nums[i] = int32(i);
	}
	out := solve(In{nums})

	if !out.sortable {
		t.Errorf("sortable is not correct, expected %v, returned %v", true, out.sortable)
	}
}
