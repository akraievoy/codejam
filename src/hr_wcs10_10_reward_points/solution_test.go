package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{[]int{10, 20, 5}})

	{
		expected := 250;
		actual := out.reward;

		if actual != expected {
			t.Errorf("out.reward is not correct: expected %v, returned %v", expected, actual)
		}
	}}
