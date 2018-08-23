package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	// FIXME simple test case
	out := solve(In{[]bool{false, true, true, false, true, true, true, true, false, false}})

	{
		actual := out.movesMin
		expected := 3
		if actual != expected {
			t.Errorf("out.movesMin is not correct: expected %v, actual %v", expected, actual)
		}
	}

	{
		actual := out.movesMax
		expected := 6
		if actual != expected {
			t.Errorf("out.movesMax is not correct: expected %v, actual %v", expected, actual)
		}
	}

}
