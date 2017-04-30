package main

import (
	"testing"
	"reflect"
	"fmt"
)

func Test_XX_XX(t *testing.T) {
	_run(t, []int{-1, -1}, []int{-1, -1}, []int{0, 0}, []int{0, 0})
}

func Test_999_000(t *testing.T) {
	_run(t, []int{9, 9, 9}, []int{0, 0, 0}, []int{9, 9, 9}, []int{0, 0, 0})
}

func Test_12X_1X3(t *testing.T) {
	_run(t, []int{1, 2, -1}, []int{1, -1, 3}, []int{1, 2, 3}, []int{1, 2, 3})
}

func Test_1X_2X(t *testing.T) {
	_run(t, []int{1, -1}, []int{2, -1}, []int{1, 9}, []int{2, 0})
}

func Test_2X_1X(t *testing.T) {
	_run(t, []int{2, -1}, []int{1, -1}, []int{2, 0}, []int{1, 9})
}

func Test_1X9_2X9(t *testing.T) {
	_run(t, []int{1, -1, 9}, []int{2, -1, 9}, []int{1, 9, 9}, []int{2, 0, 9})
}

func Test_2X5_1X5(t *testing.T) {
	_run(t, []int{2, -1, 1}, []int{1, -1, 1}, []int{2, 0, 1}, []int{1, 9, 1})
}

func Test_X1X_X2X(t *testing.T) {
	_run(t, []int{-1, 1, -1}, []int{-1, 2, -1}, []int{0, 1, 9}, []int{0, 2, 0})
}

func Test_XX9_X20(t *testing.T) {
	_run(t, []int{-1, -1, 9}, []int{-1, 2, 0}, []int{0, 1, 9}, []int{0, 2, 0})
}

func Test_XX9_2X0(t *testing.T) {
	_run(t, []int{-1, -1, 9}, []int{2, -1, 0}, []int{1, 9, 9}, []int{2, 0, 0})
}

func Test_XX2_X28(t *testing.T) {
	_run(t, []int{-1, -1, 2}, []int{-1, 2, 8}, []int{0, 3, 2}, []int{0, 2, 8})
}

func Test_XX3_X28(t *testing.T) {
	_run(t, []int{-1, -1, 3}, []int{-1, 2, 8}, []int{0, 2, 3}, []int{0, 2, 8})
}

func Test_XX9_X2X(t *testing.T) {
	_run(t, []int{-1, -1, 9}, []int{-1, 2, -1}, []int{0, 2, 9}, []int{0, 2, 9})
}

func Test_X99_XX0(t *testing.T) {
	_run(t, []int{-1, 9, 9}, []int{-1, -1, 0}, []int{0, 9, 9}, []int{1, 0, 0})
}

func Test_XX6_X00(t *testing.T) {
	_run(t, []int{-1,-1, 6}, []int{-1, 0, 0}, []int{0, 9, 6}, []int{1, 0, 0})
}

func TestVersusBrute(t *testing.T) {
	input := []int{-1, -1, -1, -1, -1, -1}
	cIn, jIn := input[0:len(input)/2], input[len(input)/2:]
	errored := false
	for !errored {
		in := In{3, cIn, jIn}
		outBrute := solveBrute(in)
		out := solve(in)
		
		{
			expected := outBrute.c;
			actual := out.c;

			if !reflect.DeepEqual(actual, expected) {
				errored = true
				t.Errorf("out.c is not correct for %v:\nexpected %v,\nreturned %v", in, expected, actual)
			}
		}
		{
			expected := outBrute.j;
			actual := out.j;

			if !reflect.DeepEqual(actual, expected) {
				errored = true
				t.Errorf("out.j is not correct for %v:\nexpected %v,\nreturned %v", in, expected, actual)
			}
		}

		rollPos := len(input) - 1
		for rollPos >= 0 && input[rollPos] == 9 {
			rollPos --
		}
		if (rollPos < 0) {
			fmt.Println(fmt.Sprintf("breaking at %v", rollPos))
			break
		}
		input[rollPos]++
		rollPos++
		for rollPos < len(input) {
			input[rollPos] = -1
			rollPos++
		}
	}
}

func _run(t *testing.T, c, j, cOpt, jOpt []int) {
	out := solve(In{3, c, j})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	{
		expected := cOpt;
		actual := out.c;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("out.c is not correct: expected %v, returned %v", expected, actual)
		}
	}
	{
		expected := jOpt;
		actual := out.j;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("out.j is not correct: expected %v, returned %v", expected, actual)
		}
	}

}
