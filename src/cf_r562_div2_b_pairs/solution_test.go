package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	t.Skip()

	for z := 0; z < 1000; z++ {
		seed := time.Now().UnixNano()
		t.Logf("my seed is %d\n", seed)
		var src = rand.NewSource(seed)
		var rnd = rand.New(src)
		N, M := int64(6), int64(12)
		b := new(bytes.Buffer)
		var writer = bufio.NewWriterSize(b, 1024*1024)
		fmt.Fprintf(writer, "%d %d\n", N, M)
		pairs := make([][2]int64, M, M)
		v := make(map[int64]bool)
		for i := int64(0); i < M; i++ {
			pairs[i][0] = rnd.Int63n(N) + 1
			pairs[i][1] = (pairs[i][0] - 1 + rnd.Int63n(N-1) + 1) % N + 1
			if pairs[i][0] == pairs[i][1] {
				t.Fatal("ouch! random bites my shiny metal ass")
			}
			fmt.Fprintf(writer, "%d %d\n", pairs[i][0], pairs[i][1])
			v[pairs[i][0]] = true
			v[pairs[i][1]] = true
		}
		writer.Flush()

		possible, posV0, posV1 := false, int64(-1), int64(-1)
		for v0 := range v {
			for v1 := range v {
				if v0 >= v1 {
					continue
				}

				possibleNow := true
				for i := int64(0); i < M; i++ {
					if pairs[i][0] == v0 || pairs[i][0] == v1 || pairs[i][1] == v0 || pairs[i][1] == v1 {
						//	ooookay
					} else {
						possibleNow = false
						break
					}
				}
				if possibleNow {
					possible = true
					posV0 = v0
					posV1 = v1
					break
				}
			}
			if possible {
				break
			}
		}

		outputExpected := fmt.Sprintf(ts(possible, "YES\n", "NO\n"))
		j, closeResFunc := JamNewMock(b.String())
		solveAll(j)
		res := closeResFunc()
		if res != outputExpected {
			t.Errorf("INPUT:\n%s\nOUTPUT:\n%s", b.String(), outputExpected)
			t.Fatalf("test %d (ans = %d,%d) returned:\n %s", z, posV0, posV1, res)
		}
	}
}

func ts(b bool, t, f string) string { if b {return t};return f }

func JamNewMock(input string) (Jam, func() string) {
	var scanner = bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	b := new(bytes.Buffer)
	var writer = bufio.NewWriterSize(b, 1024*1024)
	jam := &jam{scanner, writer}
	return jam, func() string {
		jam.Close()
		return b.String()
	}
}
