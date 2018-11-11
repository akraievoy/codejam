package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solveCase(caseInput{3, 7, []string{"HELPIAM","TRAPPED","INSIDEA","CODEJAM","FACTORY"}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	{
		actual := out.word
		expected := "HELPIAD"

		if actual != expected {
			t.Errorf("out.word is not correct: expected %v, actual %v", expected, actual)
		}
	}

}
