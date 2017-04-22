package main

import (
	"testing"
)

func TestKitRange(t *testing.T) {
	kMin, kMax := kitRange(900, 500)

	{
		expected := 2;
		actual := kMin;

		if actual != expected {
			t.Errorf("kMin is not correct: expected %v, returned %v", expected, actual)
		}
	}
	{
		expected := 2;
		actual := kMax;

		if actual != expected {
			t.Errorf("kMax is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestKitRange2(t *testing.T) {
	kMin, kMax := kitRange(660, 300)

	{
		expected := 2;
		actual := kMin;

		if actual != expected {
			t.Errorf("kMin is not correct: expected %v, returned %v", expected, actual)
		}
	}
	{
		expected := 2;
		actual := kMax;

		if actual != expected {
			t.Errorf("kMax is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestKitRange3(t *testing.T) {
	kMin, kMax := kitRange(809, 300)

	{
		expected := -1;
		actual := kMin;

		if actual != expected {
			t.Errorf("kMin is not correct: expected %v, returned %v", expected, actual)
		}
	}
	{
		expected := -1;
		actual := kMax;

		if actual != expected {
			t.Errorf("kMax is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestForErrorBreakingMySmall(t *testing.T) {
	kMin, kMax := kitRange(1350, 3)

	{
		expected := 410;
		actual := kMin;

		if actual != expected {
			t.Errorf("kMin is not correct: expected %v, returned %v", expected, actual)
		}
	}
	{
		expected := 500;
		actual := kMax;

		if actual != expected {
			t.Errorf("kMax is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
