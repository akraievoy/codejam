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

	// FIXME simple test case
	out := solveCase(caseInput{3, []int16{2, 3, 5}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if out.sum != 10 {
		t.Errorf("sum is not correct, expected %d, returned %d", 10, out.sum)
	}
}
