package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solveCase(caseInput{2, []int32{1, 2, 5}})

	{
		actual := out.minutes
		expected := uint32(1)

		if actual != expected {
			t.Errorf("out.minutes is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}

}
