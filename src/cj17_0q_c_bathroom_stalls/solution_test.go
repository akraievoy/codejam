package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{9, 7, 1})

	if out.index != 9 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if out.maxLR != 3 {
		t.Errorf("maxLR is not correct, expected %d, returned %d", 3, out.maxLR)
	}
	if out.minLR != 3 {
		t.Errorf("minLR is not correct, expected %d, returned %d", 3, out.minLR)
	}
}

func TestSimple1(t *testing.T) {
	out := solve(In{3, 7, 3})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if out.maxLR != 1 {
		t.Errorf("maxLR is not correct, expected %d, returned %d", 1, out.maxLR)
	}
	if out.minLR != 1 {
		t.Errorf("minLR is not correct, expected %d, returned %d", 1, out.minLR)
	}
}

func TestSimpleSingleSlots(t *testing.T) {
	for k := int64(4); k <= int64(7); k++ {
		out := solve(In{3, 7, k})
		if out.maxLR != 0 {
			t.Errorf("maxLR is not correct for k = %d, expected %d, returned %d", k, 0, out.maxLR)
		}
		if out.minLR != 0 {
			t.Errorf("minLR is not correct for k = %d, expected %d, returned %d", k, 0, out.minLR)
		}
	}
}

func TestOdd(t *testing.T) {
	out := solve(In{9, 8, 1})

	if out.maxLR != 4 {
		t.Errorf("maxLR is not correct, expected %d, returned %d", 4, out.maxLR)
	}
	if out.minLR != 3 {
		t.Errorf("minLR is not correct, expected %d, returned %d", 3, out.minLR)
	}
}

func TestOddNext(t *testing.T) {
	out := solve(In{9, 8, 2})

	if out.maxLR != 2 {
		t.Errorf("maxLR is not correct, expected %d, returned %d", 2, out.maxLR)
	}
	if out.minLR != 1 {
		t.Errorf("minLR is not correct, expected %d, returned %d", 1, out.minLR)
	}
}

//	G.*.0.0..G
func TestOddNextNext(t *testing.T) {
	out := solve(In{9, 8, 3})

	if out.maxLR != 1 {
		t.Errorf("maxLR is not correct, expected %d, returned %d", 1, out.maxLR)
	}
	if out.minLR != 1 {
		t.Errorf("minLR is not correct, expected %d, returned %d", 1, out.minLR)
	}
}

//	G.0.0.0*.G
func TestOddNextNextNext(t *testing.T) {
	out := solve(In{9, 8, 4})

	if out.maxLR != 1 {
		t.Errorf("maxLR is not correct, expected %d, returned %d", 1, out.maxLR)
	}
	if out.minLR != 0 {
		t.Errorf("minLR is not correct, expected %d, returned %d", 0, out.minLR)
	}
}

func TestOddSingleSlots(t *testing.T) {
	for k := int64(5); k <= int64(7); k++ {
		out := solve(In{3, 7, k})
		if out.maxLR != 0 {
			t.Errorf("maxLR is not correct for k = %d, expected %d, returned %d", k, 0, out.maxLR)
		}
		if out.minLR != 0 {
			t.Errorf("minLR is not correct for k = %d, expected %d, returned %d", k, 0, out.minLR)
		}
	}
}

func TestLarge(t *testing.T) {
	out := solve(In{3, 1000*1000*1000*1000*1000*1000, 1000*1000*1000*1000*1000*1000})
	if out.maxLR != 0 {
		t.Errorf("maxLR is not correct, expected %d, returned %d", 0, out.maxLR)
	}
	if out.minLR != 0 {
		t.Errorf("minLR is not correct, expected %d, returned %d", 0, out.minLR)
	}
}
