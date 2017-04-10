package main

import (
	"testing"
	"reflect"
)

func TestSimple(t *testing.T) {
	
	out := solve(In{[]uint32{1, 2, 3, 4, 5, 6, 7, 8, 13, 15, 2*1000*1000*1000}})

	primesExpected := []bool{false, true, true, false, true, false, true, false, true, false, false}
	if !reflect.DeepEqual(out.primes, primesExpected) {
		t.Errorf("primes is not correct, expected\n\t%+v,\nreturned\n\t%+v", primesExpected, out.primes)
	}
}
