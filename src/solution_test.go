package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{[]int{7, 2, 3, 5, 1}})

	{
		expected := 1;
		actual := out.removed;

		if actual != expected {
			t.Errorf("out.removed is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestSimple01(t *testing.T) {
	out := solve(In{[]int{6, 5, 4, 9, 1}})

	{
		expected := 1;
		actual := out.removed;

		if actual != expected {
			t.Errorf("out.removed is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestSimple02(t *testing.T) {
	out := solve(In{[]int{9, 5, 7, 8, 1, 2}})

	{
		expected := 1;
		actual := out.removed;

		if actual != expected {
			t.Errorf("out.removed is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestSimple03(t *testing.T) {
	out := solve(In{[]int{1, 2, 3, 4}})

	{
		expected := 2;
		actual := out.removed;

		if actual != expected {
			t.Errorf("out.removed is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
