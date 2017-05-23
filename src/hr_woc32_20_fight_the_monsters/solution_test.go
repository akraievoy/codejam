package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{2, 100, []int{2, 3, 5}})

	{
		expected := 3;
		actual := out.maxKilledZombies;

		if actual != expected {
			t.Errorf("out.maxKilledZombies is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestOther(t *testing.T) {
	out := solve(In{100, 2, []int{2, 3, 5}})

	{
		expected := 2;
		actual := out.maxKilledZombies;

		if actual != expected {
			t.Errorf("out.maxKilledZombies is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
