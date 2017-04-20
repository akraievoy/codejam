package main

import (
	"testing"
	"reflect"
)

func TestSimple(t *testing.T) {
	out := solve(In{
		5,
		0,
		[]int{0, 1, 3, 4, 4},
		[]Link{{0, 1}, {1, 0}, {1, 2}, {2, 1}}})

	expected := []int{0, 6, 12, -1, -1}
	if !reflect.DeepEqual(out.distances, expected) {
		t.Errorf("distances is not correct, expected %v, returned %v", expected, out.distances)
	}
}

func TestPath10(t *testing.T) {
	fingers := []int{0, 1, 3, 5, 7, 9, 11, 13, 15, 17, 18}

	links := []Link{
		{0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 3}, {3, 2}, {3, 4}, {4, 3}, {4, 5},
		{5, 4}, {5, 6}, {6, 5}, {6, 7}, {7, 6}, {7, 8}, {8, 7}, {8, 9}, {9, 8}}

	out := solve(In{
		10,
		5,
		fingers,
		links})

	{
		expected := []int{30, 24, 18, 12, 6, 0, 6, 12, 18, 24};
		actual := out.distances;

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("out.distances is not correct: expected %v, returned %v", expected, actual)
		}
	}

}
