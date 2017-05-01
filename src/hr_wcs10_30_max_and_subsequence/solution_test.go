package main

import (
	"testing"
	"math/big"
)

func TestSimple(t *testing.T) {
	out := solve(In{2, []int{2, 3, 5}})

	{
		expected := 2;
		actual := out.MaxAnd;

		if actual != expected {
			t.Errorf("out.MaxAnd is not correct: expected %v, returned %v", expected, actual)
		}
	}
	{
		expected := big.NewInt(1);
		actual := out.Count;

		if actual.Cmp(expected) != 0 {
			t.Errorf("out.Count is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
