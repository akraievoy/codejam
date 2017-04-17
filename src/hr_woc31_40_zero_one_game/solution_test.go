package main


import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{[]bool{false, false, false}})

	if out.deletions != 1 {
		t.Errorf("deletions is not correct, expected %d, returned %d", 1, out.deletions)
	}
}

func TestSimpleSingleOne(t *testing.T) {
	out := solve(In{[]bool{false, true, false}})

	if out.deletions != 1 {
		t.Errorf("deletions is not correct, expected %d, returned %d", 1, out.deletions)
	}
}

func TestSimpleSingleEdgeyOne(t *testing.T) {
	out := solve(In{[]bool{true, false, false}})

	if out.deletions != 0 {
		t.Errorf("deletions is not correct, expected %d, returned %d", 0, out.deletions)
	}
}

func TestLarge(t *testing.T) {
	out := solve(In{make([]bool, 100*1000)})

	deletionsExpected := int32(100 * 1000 - 2)
	if out.deletions != deletionsExpected {
		t.Errorf("deletions is not correct, expected %d, returned %d", deletionsExpected, out.deletions)
	}
}

func TestLargeWithOnes(t *testing.T) {
	flags := make([]bool, 100 * 1000)

	flags[10000] = true
	flags[20000] = true
	flags[30000] = true
	flags[40000] = true
	flags[50000] = true
	flags[60000] = true
	flags[70000] = true
	flags[80000] = true
	flags[90000] = true
	
	out := solve(In{flags})

	deletionsExpected := int32(100 * 1000 - 2)
	if out.deletions != deletionsExpected {
		t.Errorf("deletions is not correct, expected %d, returned %d", deletionsExpected, out.deletions)
	}
}
func TestLargeWithOneAreas(t *testing.T) {
	flags := make([]bool, 100 * 1000)

	flags[10000] = true

	flags[20000] = true
	flags[20001] = true
	flags[20002] = true

	flags[30000] = true
	flags[40000] = true

	flags[50000] = true
	flags[50001] = true

	flags[60000] = true
	flags[70000] = true
	flags[80000] = true
	flags[90000] = true

	out := solve(In{flags})

	deletionsExpected := int32(100 * 1000 - 2 - 5 - 4)
	if out.deletions != deletionsExpected {
		t.Errorf("deletions is not correct, expected %d, returned %d", deletionsExpected, out.deletions)
	}
}
