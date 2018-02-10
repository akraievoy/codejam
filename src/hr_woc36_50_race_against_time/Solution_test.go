package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{[]Runner{{5, 0},{2,2},{6,3},{2,2}}})

	{
		actual := out.MinTotalCost
		expected := int64(8)
		if actual != expected {
			t.Errorf("out.MinTotalCost is not correct: expected %v, actual %v", expected, actual)
		}
	}
}
