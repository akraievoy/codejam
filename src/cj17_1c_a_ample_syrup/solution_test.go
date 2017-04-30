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
	out := solve(In{3, 2, 1, []PC{{100,20},{200, 10}}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}

	{
		expected := float64(200 * 200 + 2 * 10 * 200);
		actual := out.maxS;

		if actual != expected {
			t.Errorf("out.maxS is not correct: expected %v, returned %v", expected, actual)
		}
	}
}

func TestFailing(t *testing.T) {
	var rnd = rand.NewSource(time.Now().UnixNano())
	if false {
		rnd.Int63()
	}
	out := solve(In{3, 4, 2, []PC{{9,3},{7, 1},{10, 1},{8, 4}}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}

	{
		expected := float64(9*9+2*3*9+2*8*4);
		actual := out.maxS;

		if actual != expected {
			t.Errorf("out.maxS is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
