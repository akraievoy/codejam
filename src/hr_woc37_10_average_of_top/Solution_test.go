package main

import (
	"testing"
)

func TestRoundUp(t *testing.T) {
	ratings := make([]int, 200)
	for i := 0; i < 69; i++ {
		ratings[i] = 96
	}
	for i := 69; i < 200; i++ {
		ratings[i] = 95
	}
	out := solveCase(caseInput{ratings})

	{
		actual := out.averageOfTop
		expected := "95.35"

		if actual != expected {
			t.Errorf("out.averageOfTop is not correct: expected %v, actual %v", expected, actual)
		}
	}

}

func TestRoundDown(t *testing.T) {
	ratings := make([]int, 200)
	for i := 0; i < 68; i++ {
		ratings[i] = 96
	}
	for i := 68; i < 199; i++ {
		ratings[i] = 95
	}
	out := solveCase(caseInput{ratings})

	{
		actual := out.averageOfTop
		expected := "95.34"

		if actual != expected {
			t.Errorf("out.averageOfTop is not correct: expected %v, actual %v", expected, actual)
		}
	}

}
