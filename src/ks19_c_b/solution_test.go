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
		R, C, K := int64(5), int64(5), 0
		b := new(bytes.Buffer)
		var writer = bufio.NewWriterSize(b, 1024*1024)
		fmt.Fprintf(writer, "1\n%d %d %d\n", R, C, K)
		mtx := make([][]int64, R, R)
		for r := int64(0); r < R; r++ {
			mtx[r] = make([]int64, C, C)
			for c := int64(0); c < C; c++ {
				mtx[r][c] = int64(rnd.Intn(2))
				fmt.Fprintf(writer, "%d ", mtx[r][c])
			}
			fmt.Fprintf(writer, "\n")
		}
		writer.Flush()

		best := int64(1)
		for c0 := int64(0); c0 < C; c0++ {
			for c1 := int64(0); c1 < C; c1++ {
				for r0 := int64(0); r0 < R; r0++ {
					for r1 := r0; r1 < R; r1++ {
						valid := true
						for r := r0; r <= r1 && valid; r++ {
							mi, ma := mtx[r][c0], mtx[r][c0]
							for c := c0; c <= c1; c++ {
								mi, ma := min64(mi, mtx[r][c]), max64(ma, mtx[r][c])
								if int(ma-mi) > K {
									valid = false
									break
								}
							}
						}
						if valid {
							if best < (1+r1-r0)*(1+c1-c0) {
								t.Logf("next best: %d %d %d %d\n", r0, c0, r1, c1)
								best = (1+r1-r0)*(1+c1-c0)
							}
						}
					}
				}
			}
		}

		outputExpected := fmt.Sprintf("Case #1: %d\n", best)
		j, closeResFunc := JamNewMock(b.String())
		solveAll(j)
		res := closeResFunc()
		if res != outputExpected {
			t.Errorf("INPUT:\n%s\nOUTPUT:\n%s", b.String(), outputExpected)
			t.Fatalf("test %d: sum is not correct, expected %d, returned %s", z, best, res)
		}
	}
}

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
