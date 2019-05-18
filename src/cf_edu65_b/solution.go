package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(j Jam) {
	j.PF("? 1 2\n")
	p01 := j.Int()
	j.PF("? 2 3\n")
	p12 := j.Int()
	j.PF("? 3 4\n")
	p23 := j.Int()
	j.PF("? 4 5\n")
	p34 := j.Int()

	nums := [6]int64{4, 8, 15, 16, 23, 42}
	perm := PermNew(6)

	for {
		t01 := nums[perm[0]] * nums[perm[1]]
		t12 := nums[perm[1]] * nums[perm[2]]
		t23 := nums[perm[2]] * nums[perm[3]]
		t34 := nums[perm[3]] * nums[perm[4]]
		if p01 == t01 && p12 == t12 && p23 == t23 && p34 == t34 {
			fmt.Printf("! %d %d %d %d %d %d\n",
				nums[perm[0]], nums[perm[1]],
				nums[perm[2]], nums[perm[3]],
				nums[perm[4]], nums[perm[5]],
			)
			return
		}
		if !perm.Next() {
			panic("not possible")
		}
	}
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

//	GoLang shorthand methods for math go below
//	TODO wipe unused methods before submitting

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

type Perm []uint32

func PermNew(size int) Perm {
	perm := make([]uint32, size)
	for i := range perm {
		perm[i] = uint32(i)
	}
	return perm
}

//	https://www.nayuki.io/page/next-lexicographical-permutation-algorithm

func (perm Perm) Next() bool {
	rollPos := len(perm) - 2
	for rollPos >= 0 && perm[rollPos] > perm[rollPos+1] {
		rollPos -= 1
	}
	if rollPos < 0 {
		return false
	}

	swapPos := rollPos + 1
	for pos := rollPos + 2; pos < len(perm); pos++ {
		if perm[rollPos] < perm[pos] && perm[pos] < perm[swapPos] {
			swapPos = pos
		}
	}
	perm[rollPos], perm[swapPos] = perm[swapPos], perm[rollPos]

	for i, j := rollPos+1, len(perm)-1; i < j; i, j = i+1, j-1 {
		perm[i], perm[j] = perm[j], perm[i]
	}

	return true
}
