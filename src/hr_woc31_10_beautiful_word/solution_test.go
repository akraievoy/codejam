package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	out := solve(In{"bsdartesicapetatatat"})

	if !out.beautiful {
		t.Errorf("beatiful is not correct, expected %v, returned %v", true, out.beautiful)
	}
}

func TestSimpleNeg(t *testing.T) {
	out := solve(In{"bsdartessicapetatatat"})

	if out.beautiful {
		t.Errorf("beatiful is not correct, expected %v, returned %v", false, out.beautiful)
	}
}

func TestSimpleNegVowel(t *testing.T) {
	out := solve(In{"bsdartesicapetaotatat"})

	if out.beautiful {
		t.Errorf("beatiful is not correct, expected %v, returned %v", false, out.beautiful)
	}
}
