package main

import (
	"testing"
	"reflect"
)

func TestSimple(t *testing.T) {
	out := solve(In{3, []int8{1, 2, 2, 2, 0}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	expected := []int8{1, 1, 9, 9, 9}
	if !reflect.DeepEqual(out.digits, expected) {
		t.Errorf("digits are not correct, expected %+v, returned %+v", expected, out.digits)
	}
}

func TestSimpleOneDigit8(t *testing.T) {
	out := solve(In{0, []int8{8}})

	expected := []int8{8}
	if !reflect.DeepEqual(out.digits, expected) {
		t.Errorf("digits are not correct, expected %+v, returned %+v", expected, out.digits)
	}
}

func TestSimpleOneDigit0(t *testing.T) {
	out := solve(In{0, []int8{0}})

	expected := []int8{0}
	if !reflect.DeepEqual(out.digits, expected) {
		t.Errorf("digits are not correct, expected %+v, returned %+v", expected, out.digits)
	}
}

func TestSimpleNoOp(t *testing.T) {
	out := solve(In{0, []int8{1, 3, 5, 7, 9}})

	expected := []int8{1, 3, 5, 7, 9}
	if !reflect.DeepEqual(out.digits, expected) {
		t.Errorf("digits are not correct, expected %+v, returned %+v", expected, out.digits)
	}
}

func TestSimpleRolldown(t *testing.T) {
	out := solve(In{0, []int8{1, 3, 5, 7, 9, 8}})

	expected := []int8{1, 3, 5, 7, 8, 9}
	if !reflect.DeepEqual(out.digits, expected) {
		t.Errorf("digits are not correct, expected %+v, returned %+v", expected, out.digits)
	}
}

func TestLongRolldown(t *testing.T) {
	out := solve(In{0, []int8{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5, 2, 1, 0}})

	expected := []int8{1, 2, 3, 4, 4, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	if !reflect.DeepEqual(out.digits, expected) {
		t.Errorf("digits are not correct, expected %+v, returned %+v", expected, out.digits)
	}
}

func TestZeroingRolldown(t *testing.T) {
	out := solve(In{0, []int8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}})

	expected := []int8{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	if !reflect.DeepEqual(out.digits, expected) {
		t.Errorf("digits are not correct, expected %+v, returned %+v", expected, out.digits)
	}
}
