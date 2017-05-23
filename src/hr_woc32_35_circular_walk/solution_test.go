package main

import (
	"testing"
)

func TestWithin(t *testing.T) {
	{
		expected := true;
		actual := within(10, 90, 50);

		if actual != expected {
			t.Errorf("within(10, 90, 50) is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := true;
		actual := within(10, 90, 10);

		if actual != expected {
			t.Errorf("within(10, 90, 89) is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := true;
		actual := within(10, 90, 89);

		if actual != expected {
			t.Errorf("within(10, 90, 89) is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := false;
		actual := within(10, 90, 90);

		if actual != expected {
			t.Errorf("within(10, 90, 90) is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := true;
		actual := within(90, 10, 90);

		if actual != expected {
			t.Errorf("within(10, 90, 90) is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := false;
		actual := within(90, 10, 10);

		if actual != expected {
			t.Errorf("within(10, 90, 10) is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := false;
		actual := within(90, 10, 50);

		if actual != expected {
			t.Errorf("within(10, 90, 50) is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := true;
		actual := within(90, 10, 91);

		if actual != expected {
			t.Errorf("within(10, 90, 91) is not correct: expected %v, returned %v", expected, actual)
		}
	}

	{
		expected := true;
		actual := within(90, 10, 9);

		if actual != expected {
			t.Errorf("within(10, 90, 9) is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestSimple(t *testing.T) {
	out := solve(In{9, 0, 2, 1, 3, 4, 7})

	{
		expected := 2;
		actual := out.seconds;

		if actual != expected {
			t.Errorf("out.seconds is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestSimple2(t *testing.T) {
	out := solve(In{9, 0, 5, 1, 3, 4, 3})

	{
		expected := 5;
		actual := out.seconds;

		if actual != expected {
			t.Errorf("out.seconds is not correct: expected %v, returned %v", expected, actual)
		}
	}

	out2 := solve(In{9, 5, 0, 1, 3, 4, 3})

	{
		expected := 5;
		actual := out2.seconds;

		if actual != expected {
			t.Errorf("out2.seconds is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestSimple3(t *testing.T) {
	out := solve(In{9, 8, 4, 1, 3, 4, 3})

	{
		expected := 5;
		actual := out.seconds;

		if actual != expected {
			t.Errorf("out.seconds is not correct: expected %v, returned %v", expected, actual)
		}
	}

	out2 := solve(In{9, 4, 8, 1, 3, 4, 3})

	{
		expected := 5;
		actual := out2.seconds;

		if actual != expected {
			t.Errorf("out2.seconds is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
