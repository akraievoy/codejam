package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF

	n, x := l(), l()
	a := make([]int64, n, n)
	var b [300001][3][3]int64

	for i := int64(0); i < n; i++ {
		a[i] = l()
	}
	b[0][0][0] = 0

	for i := int64(0); i <= n; i++ {
		for j := int64(0); j < 3; j++ {
			for k := int64(0); k < 3; k++ {
				if k < 2 {
					b[i][j][k+1] =
						max(
							b[i][j][k+1],
							b[i][j][k],
						)
				}
				if j < 2 {
					b[i][j+1][k] =
						max(
							b[i][j+1][k],
							b[i][j][k],
						)
				}
				if i < n {
					b[i+1][j][k] =
						max(
							b[i+1][j][k],
							b[i][j][k]+t(j == 1, a[i], 0)*t(k == 1, x, 1),
						)
				}
			}
		}
	}

	pf("%d\n", b[n][2][2])

	if false { d("%v", []interface{}{s, l, ui, f, d, p, pf}) }
}

func t(b bool, t, f int64) int64 {
	if b {
		return t
	}
	return f
}
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	jam, closeFunc := JamNew()
	defer closeFunc()
	solveAll(jam)
}

type Jam interface {
	Scanner() *bufio.Scanner
	Writer() *bufio.Writer
	Close()

	Str() string
	Long() int64
	Int() uint32
	Float() float64

	D(format string, values ...interface{})
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

	var writer = bufio.NewWriterSize(os.Stdout, 1024*1024)
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

func (j *jam) Long() int64 {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	res, err := strconv.ParseInt(j.sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	return res
}

func (j *jam) Int() uint32 {
	return uint32(j.Long())
}

func (j *jam) Float() float64 {
	j.sc.Scan()
	res, err := strconv.ParseFloat(j.sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (j *jam) D(format string, values ...interface{}) {
	_ /*bytesWritten*/, err := fmt.Fprintf(os.Stderr, format, values...)
	if err != nil {
		panic(err)
	}
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
