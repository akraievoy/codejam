package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"strings"
	"testing"
	"time"
	"math/rand"
)

func TestCompare(t *testing.T) {
	{
		actual := compare(4, -2, math.MaxInt64, 1)
		expected := -1

		if actual != expected {
			t.Errorf("compare(4, -2, math.MaxInt64, 1) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}
	{
		actual := compare(-1, 3, 0, 1)
		expected := -1

		if actual != expected {
			t.Errorf("compare(-1, 3, 0, 1) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}
	{
		actual := compare(1, -1, 1000, 1)
		expected := -1

		if actual != expected {
			t.Errorf("compare(1, -1, 9223372036854775807, 1) is not correct\n\texpected:\t%v\n\tactual:\t%v", expected, actual)
		}
	}

}

func TestSimple(t *testing.T) {
	t.SkipNow()

	for z := 0; z < 1000; z++ {
		seed := time.Now().UnixNano()
		fmt.Printf("my seed is %d\n", seed)
		var src = rand.NewSource(seed)
		var rnd = rand.New(src)
		c, j := make([]int64, 3, 3), make([]int64, 3, 3)
		for i := int64(0); i < 3; i++ {
			c[i], j[i] = rnd.Int63n(100)+1, rnd.Int63n(100)+1
		}
		notFound := true
		var outputExpected string
		for C := int64(1); C <= 10000 && notFound; C++ {
			for J := int64(1); J < 10000 && notFound; J++ {
				if c[0]*C+j[0]*J < c[1]*C+j[1]*J && c[1]*C+j[1]*J < c[2]*C+j[2]*J {
					outputExpected = fmt.Sprintf("Case #1: %d %d\n", C, J)
					notFound = false
				}
			}
		}
		if notFound {
			outputExpected = fmt.Sprintf("Case #1: IMPOSSIBLE\n")
		}
		input := fmt.Sprintf("1\n3\n%d %d\n%d %d\n%d %d\n", c[0], j[0], c[1], j[1], c[2], j[2], )
		jam, closeResFunc := JamNewMock(input)
		solveAll(jam)
		res := closeResFunc()
		if res != outputExpected {
			t.Fatalf("INPUT:\n%s\nEXPECTED:\n%s\nACTUAL:\n%s\n", input, outputExpected, res)
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
