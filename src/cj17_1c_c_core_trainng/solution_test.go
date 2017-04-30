package main

import (
	"testing"
	"time"
	"math/rand"
)

func TestSimple(t *testing.T) {
	var rnd = rand.NewSource(time.Now().UnixNano())
	if false {
		rnd.Int63()
	}
	out := solve(In{3, 2, 2, 1.0, []float64{0.0, 0.0}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	{
		expected := 0.25;
		actual := out.maxProb;

		if actual != expected {
			t.Errorf("out.maxProb is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestFailing(t *testing.T) {
	var rnd = rand.NewSource(time.Now().UnixNano())
	if false {
		rnd.Int63()
	}
	out := solve(In{3, 4, 4, 1.4, []float64{0.5, 0.7, 0.8, 0.6}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	{
		expected := 1.00;
		actual := out.maxProb;

		if actual != expected {
			t.Errorf("out.maxProb is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
