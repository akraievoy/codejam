package main

import (
	"testing"
	"time"
	"math/rand"
)

func TestSimple(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if false {
		rnd.Int63()
	}

	// FIXME simple test case
	out := solveCase(caseInput{[]int16{2, 3, 5}})

	{
		actual := out.equat
		expected := 2

		if actual != expected {
			t.Errorf("out.equat is not correct: expected %v, actual %v", expected, actual)
		}
	}

}
