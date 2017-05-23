package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{11})

	{
		expected := 1;
		actual := out.charAtX;

		if actual != expected {
			t.Errorf("out.charAtX is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
