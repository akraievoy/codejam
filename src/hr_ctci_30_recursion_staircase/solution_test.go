package main

import (
	"testing"
	"reflect"
)

func TestSimple(t *testing.T) {
	out := solve(In{[]int16{1,2,3,4,5,6,7}})

	countsExpected := []uint64{1, 2, 4, 7, 13, 24, 44}
	if !reflect.DeepEqual(out.counts, countsExpected) {
		t.Errorf("sum is not correct, expected\n%v, returned\n%v", countsExpected, out.counts)
	}
}
