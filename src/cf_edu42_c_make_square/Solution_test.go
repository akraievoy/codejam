package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	lcsTo81 := lcs("8314", "81")

	{
		actual := lcsTo81
		expected := "81"

		if actual != expected {
			t.Errorf("lcsTo81 is not correct: expected %v, actual %v", expected, actual)
		}
	}

}

func TestSimple2(t *testing.T) {
	lcsTo81 := lcs("333", "81")

	{
		actual := lcsTo81
		expected := ""

		if actual != expected {
			t.Errorf("lcsTo81 is not correct: expected %v, actual %v", expected, actual)
		}
	}

}
