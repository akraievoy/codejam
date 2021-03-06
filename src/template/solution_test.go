package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
	"math/rand"
)

func TestSimple(t *testing.T) {
	seed := time.Now().UnixNano()
	fmt.Printf("my seed is %d\n", seed)
	var src = rand.NewSource(seed)
	var rnd = rand.New(src)

	a := rnd.Int31n(100)
	b := rnd.Int31n(100)
	c := rnd.Int31n(100)

	input := fmt.Sprintf("1\n3\n%d %d %d", a, b, c)
	outputExpected := fmt.Sprintf("Case #1: %d\n", a+b+c)

	j, closeResFunc := JamNewMock(input)
	solveAll(j)
	res := closeResFunc()

	if res != outputExpected {
		t.Errorf("sum is not correct, expected %d, returned %s", a+b+c, res)
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
