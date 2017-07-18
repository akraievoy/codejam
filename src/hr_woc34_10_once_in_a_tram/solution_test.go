package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{555555})

	{
		expected := 555564;
		actual := out.nextLucky;

		if actual != expected {
			t.Errorf("out.nextLucky is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestRollUp(t *testing.T) {
	out := solve(In{900900})

	{
		expected := 901019;
		actual := out.nextLucky;

		if actual != expected {
			t.Errorf("out.nextLucky is not correct: expected %v, returned %v", expected, actual)
		}
	}
}


func TestMaxIn(t *testing.T) {
	out := solve(In{1000000 - 2})

	{
		expected := 999999;
		actual := out.nextLucky;

		if actual != expected {
			t.Errorf("out.nextLucky is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestMinIn(t *testing.T) {
	out := solve(In{100000})

	{
		expected := 100001;
		actual := out.nextLucky;

		if actual != expected {
			t.Errorf("out.nextLucky is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestSimple2(t *testing.T) {
	out := solve(In{165901})

	{
		expected := 165903;
		actual := out.nextLucky;

		if actual != expected {
			t.Errorf("out.nextLucky is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
