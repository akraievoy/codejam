package algo

import (
	"testing"
	"reflect"
)

func TestPermsOfThree(t *testing.T) {
	permNew := PermNew(3)

	{
		actual := permNew
		expected := Perm([]uint32{0, 1, 2})

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("permNew is not correct: expected %v, actual %v", expected, actual)
		}
	}

	for _, expectedPerm := range [][]uint32{{0, 2, 1}, {1, 0 ,2}, {1, 2 ,0}, {2, 0 ,1}, {2, 1 ,0}} {
		{
			actual := permNew.Next()
			expected := true

			if actual != expected {
				t.Errorf("permNew.Next() is not correct: expected %v, actual %v", expected, actual)
			}
		}

		{
			actual := permNew
			expected := Perm(expectedPerm)

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("permNew is not correct: expected %v, actual %v", expected, actual)
			}
		}
	}

	{
		actual := permNew.Next()
		expected := false

		if actual != expected {
			t.Errorf("permNew.Next() is not correct: expected %v, actual %v", expected, actual)
		}
	}
}

func TestPermsOfFive(t *testing.T) {
	perm := PermNew(5)

	perms := 1
	for perm.Next() {
		perms++
	}

	{
		actual := perms
		expected := 120

		if actual != expected {
			t.Errorf("perms is not correct: expected %v, actual %v", expected, actual)
		}
	}
}
