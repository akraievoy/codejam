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
	out := solve(In{3, []int{2, 3, 5}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
}
