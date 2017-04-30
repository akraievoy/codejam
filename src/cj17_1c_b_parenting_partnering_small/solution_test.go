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
	out := solve(In{3, []A{{540, 600}}, []A{{840, 900}}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if out.minExchanges != 2 {
		t.Errorf("sum is not correct, expected %d, returned %d", 2, out.minExchanges)
	}
}
