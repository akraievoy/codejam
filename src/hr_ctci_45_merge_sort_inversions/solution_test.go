package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{[]int32{2, 3, 5}})

	if out.inversions != 0 {
		t.Errorf("sum is not correct, expected %d, returned %d", 0, out.inversions)
	}
}

func TestSimpleInverted(t *testing.T) {
	out := solve(In{[]int32{5, 3, 2}})

	if out.inversions != 3 {
		t.Errorf("sum is not correct, expected %d, returned %d", 3, out.inversions)
	}
}

func TestSimpleInverted2(t *testing.T) {
	out := solve(In{[]int32{5, 5, 5, 5, 1}})

	if out.inversions != 4 {
		t.Errorf("sum is not correct, expected %d, returned %d", 4, out.inversions)
	}
}

func TestLargeFullyInverted(t *testing.T) {
	n := 1 * 100 * 1000
	nums := make([]int32, n)
	val := int32(10*1000*1000)
	for i := range nums {
		nums[i] = val
		val--
	}
	out := solve(In{nums})

	if out.inversions != n*(n-1)/2 {
		t.Errorf("sum is not correct, expected %d, returned %d", n*(n-1)/2, out.inversions)
	}
}

func TestLargeContantFill(t *testing.T) {
	n := 1 * 100 * 1000
	nums := make([]int32, n)
	val := int32(10*1000*1000)
	for i := range nums {
		nums[i] = val
	}
	out := solve(In{nums})

	if out.inversions != 0 {
		t.Errorf("sum is not correct, expected %d, returned %d", 0, out.inversions)
	}
}
