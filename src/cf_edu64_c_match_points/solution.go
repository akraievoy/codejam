package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type SortableInts64 []int64

func (si SortableInts64) Len() int           { return len(si) }
func (si SortableInts64) Less(i, j int) bool { return si[i] < si[j] }
func (si SortableInts64) Swap(i, j int)      { si[i], si[j] = si[j], si[i] }

func solve(j Jam) {
	n, z := j.Int(), j.Int()
	x := make([]int64, n, n)
	for i := range x {
		x[i] = j.Int()
	}
	sort.Sort(SortableInts64(x))
	left, right, matchedPairs := int64(0), (n+1)/2, int64(0)
	for left < right && left < (n+1)/2 && right < n {
		if x[right]-x[left] >= z {
			left += 1
			right += 1
			matchedPairs += 1
		} else {
			right += 1
		}
	}
	j.P("%d\n", matchedPairs)
}

func main() {
	jam, closeFunc := JamNew()
	defer closeFunc()
	solve(jam)
}

type Jam interface {
	Scanner() *bufio.Scanner
	Writer() *bufio.Writer
	Close()

	Str() string
	Int() int64
	Float() float64

	P(format string, values ...interface{})
	PF(format string, values ...interface{})
}

func JamNew() (Jam, func()) {
	if len(os.Args) > 1 {
		panic("running with input file path is not supported")
	}

	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	jam := &jam{scanner, writer}
	return jam, jam.Close
}

type jam struct {
	sc *bufio.Scanner
	wr *bufio.Writer
}

func (j *jam) Close() {
	if err := j.wr.Flush(); err != nil {
		panic(err)
	}
}

func (j *jam) Scanner() *bufio.Scanner {
	return j.sc
}

func (j *jam) Writer() *bufio.Writer {
	return j.wr
}

func (j *jam) Str() string {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	return j.sc.Text()
}

func (j *jam) Int() int64 {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	res, err := strconv.ParseInt(j.sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	return res
}

func (j *jam) Float() float64 {
	j.sc.Scan()
	res, err := strconv.ParseFloat(j.sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (j *jam) P(format string, values ...interface{}) {
	_, err := fmt.Fprintf(j.wr, format, values...)
	if err != nil {
		panic(err)
	}
}

func (j *jam) PF(format string, values ...interface{}) {
	_, err := fmt.Fprintf(j.wr, format, values...)
	if err != nil {
		panic(err)
	}
	if err = j.wr.Flush(); err != nil {
		panic(err)
	}
}
