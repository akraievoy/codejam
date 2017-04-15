package main

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)

func allFilled(cake [][]rune) bool {
	for _, r := range cake {
		for _, c := range r {
			if c == '?' {
				return false
			}
		}
	}
	return true
}

func TestSimple(t *testing.T) {
	out := solve(
		In{3, 3, 3, [][]rune{{'G', '?', '?'}, {'?', 'C', '?'}, {'?', '?', 'J'}}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if len(out.cake) != 3 {
		t.Errorf("rows removed, expected %d, returned %d", 3, len(out.cake))
	}

	if !allFilled(out.cake) {
		t.Error("cake NOT filled")
		for _, l := range out.cake {
			fmt.Println(fmt.Sprintf("%s", string(l)))
		}
	}
}

func TestFailing(t *testing.T) {
	out := solve(
		In{3, 3, 4,
			[][]rune{{'A', 'C', '?', '?'},
				{'?', '?', '?', 'E'},
				{'B', '?', 'D', 'F'}}})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if len(out.cake) != 3 {
		t.Errorf("rows removed, expected %d, returned %d", 3, len(out.cake))
	}

	if !allFilled(out.cake) {
		t.Error("cake NOT filled")
		for _, l := range out.cake {
			fmt.Println(fmt.Sprintf("%s", string(l)))
		}
	}

	if out.cake[0][3] == 'C' && out.cake[1][1] == 'C' {
		t.Error("cake NOT rectangular")
		for _, l := range out.cake {
			fmt.Println(fmt.Sprintf("%s", string(l)))
		}
	}
}

func TestIterated(t *testing.T) {
	for z := 0; z < 1000; z++ {
		cake := make([][]rune, 25)
		for r := range cake {
			cake[r] = make([]rune, 25)
			for c := range cake[r] {
				cake[r][c] = '?'
			}
		}
		var rnd = rand.NewSource(time.Now().UnixNano())
		cake[rnd.Int63() % 25][rnd.Int63() % 25] = 'A'
		for {
			r := rnd.Int63() % 25
			c := rnd.Int63() % 25
			if (cake[r][c] == '?') {
				cake[r][c] = 'Z'
				break;
			}
		}
		out := solve(In{3, 25, 25, cake})
		if out.index != 3 {
			t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
		}
		if len(out.cake) != 25 {
			t.Errorf("rows removed, expected %d, returned %d", 25, len(out.cake))
		}
		if !allFilled(out.cake) {
			t.Error("cake NOT filled")
			for _, l := range out.cake {
				fmt.Println(fmt.Sprintf("%s", string(l)))
			}
		}
	}
}
