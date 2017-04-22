package main

import (
	"testing"
	"math/big"
)

func TestSimple(t *testing.T) {
	out := solve(In{12, 3,[]int{2, 3, 5}})

	{
		expected := big.NewInt(5);
		actual := out.waysToChange;

		if actual.Cmp(expected) != 0 {
			t.Errorf("out.waysToChange is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
