package algo

import (
	"testing"
	"reflect"
)

func TestSubsetFlagsOfThree(t *testing.T) {
	sf := SubsetFlagsNew(3)

	{
		actual := sf
		expected := SubsetFlags([]bool{false, false, false})

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("sf is not correct: expected %v, actual %v", expected, actual)
		}
	}

	f := false
	tt := true
	flags := [][]bool{{f, f, tt}, {f, tt, f}, {f, tt, tt}, {tt, f, f}, {tt, f, tt}, {tt, tt, f}, {tt, tt, tt}}
	for _, expectedFlags := range flags {
		{
			actual := sf.Next()
			expected := true

			if actual != expected {
				t.Errorf("sf.Next() is not correct: expected %v, actual %v", expected, actual)
			}
		}

		{
			actual := sf
			expected := SubsetFlags(expectedFlags)

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("sf is not correct: expected %v, actual %v", expected, actual)
			}
		}
	}

	{
		actual := sf.Next()
		expected := false

		if actual != expected {
			t.Errorf("sf.Next() is not correct: expected %v, actual %v", expected, actual)
		}
	}
}
