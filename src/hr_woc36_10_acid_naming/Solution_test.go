package hr_woc36_10_acid_naming

import (
	"testing"
	"reflect"
)

func TestSimple(t *testing.T) {
	out := solve(In{[]string{"hydrochloric", "rainbowic", "idontevenknow"}})

	{
		actual := len(out.kinds)
		expected := 3
		if actual != expected {
			t.Errorf("len(out.kinds) is not correct: expected %v, actual %v", expected, actual)
		}
	}

	{
		actual := out.kinds
		expected := []string{"non-metal acid", "polyatomic acid", "not an acid"}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("out.kinds is not correct: expected %v, actual %v", expected, actual)
		}
	}

}
