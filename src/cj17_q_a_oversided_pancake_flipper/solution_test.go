package main

import (
	"testing"
	"os"
	"fmt"
	"time"
	"math/rand"
)

func TestNoOp(t *testing.T) {
	out := solve(In{3, []bool{false, false, false}, 2})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if out.minFlips != 0 {
		t.Errorf("minFlips is not correct, expected %d, returned %d", 0, out.minFlips)
	}
}

func TestSimple(t *testing.T) {
	out := solve(In{3, []bool{false, true, true}, 2})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if out.minFlips != 1 {
		t.Errorf("minFlips is not correct, expected %d, returned %d", 1, out.minFlips)
	}
}

func TestOverlapping(t *testing.T) {
	out := solve(In{3, []bool{true, false, true}, 2})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if out.minFlips != 2 {
		t.Errorf("minFlips is not correct, expected %d, returned %d", 2, out.minFlips)
	}
}

func TestImpossible(t *testing.T) {
	out := solve(In{3, []bool{true, false, false}, 2})

	if out.index != 3 {
		t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
	}
	if out.minFlips != -1 {
		t.Errorf("minFlips is not correct, expected %d, returned %d", -1, out.minFlips)
	}
}

func TestImpossible21(t *testing.T) {
	for pos := 0; pos < 21; pos++ {
		faces := make([]bool, 21)
		faces[pos] = true;
		out := solve(In{3, faces, 10})

		if out.index != 3 {
			t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
		}
		if out.minFlips != -1 {
			t.Errorf("minFlips for pos %d is not correct, expected %d, returned %d", pos, -1, out.minFlips)
		}
	}
}

func TestOneFlip21(t *testing.T) {
	for pos := 0; pos < 11; pos++ {
		faces := make([]bool, 21)
		for offs := 0; offs < 10; offs++ {
			faces[pos + offs] = true
		}
		out := solve(In{3, faces, 10})

		if out.index != 3 {
			t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
		}
		if out.minFlips != 1 {
			t.Errorf("minFlips for pos %d is not correct, expected %d, returned %d", pos, 1, out.minFlips)
		}
	}
}

func TestGenLarge(t *testing.T) {
	t.SkipNow()
	
	var rnd = rand.NewSource(time.Now().UnixNano())

	for f := 0; f < 10; f++ {
		inFile, err := os.Create(fmt.Sprintf("inout/large_%d.in", f))
		if err != nil {
			panic(err)
		}
		defer inFile.Close()

		outFile, err := os.Create(fmt.Sprintf("inout/large_%d.out", f))
		if err != nil {
			panic(err)
		}
		defer outFile.Close()

		inFile.WriteString("100\n")
		for c := 0; c < 100; c++ {
			facesLen := 1000
			faces := make([]bool, facesLen, facesLen)

			K := 2 + int(rnd.Int63() % int64(facesLen - 1))
			flipped := make(map[int]bool)
			flips := 0
			for i := 0; i < 100; i++ {
				start := int(rnd.Int63() % int64(facesLen - K + 1))
				if (flipped[start]) {
					continue
				}
				flipped[start] = true
				flips += 1
				for o := 0; o < K; o++ {
					faces[start + o] = !faces[start + o]
				}
			}

			for _, face := range faces {
				if face {
					inFile.WriteString("-")
				} else {
					inFile.WriteString("+")
				}
			}
			inFile.WriteString(fmt.Sprintf(" %d\n", K))

			out := solve(In{3, faces, K})

			if out.index != 3 {
				t.Errorf("index is not retained, expected %d, returned %d", 3, out.index)
			}
			if out.minFlips != flips {
				t.Errorf("minFlips for pos %d is not correct, expected %d, returned %d", t, 1, out.minFlips)
			}


			outFile.WriteString(fmt.Sprintf("Case #%d: %d\n", c+1, flips))
		}

	}
}
