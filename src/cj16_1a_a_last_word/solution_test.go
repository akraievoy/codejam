package main

import (
	"testing"
	"time"
	"math/rand"
	"reflect"
)

func TestSimple(t *testing.T) {
	var rnd = rand.NewSource(time.Now().UnixNano())
	if false {
		rnd.Int63()
	}
	// FIXME simple test case
	out := solve(In{3, []rune("ABAAB")})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	{
		expected := []rune("BBAAA");
		actual := out.lw;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("out.lw is not correct: expected %v, returned %v", expected, actual)
		}
	}
}
