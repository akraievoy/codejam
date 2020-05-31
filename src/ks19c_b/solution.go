package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var DEBUG = false

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF
	//	Live Templates: for0l for1l for0ui for1ui forr vl0 vln vui0 vuin ; Casts: l i

	T := l()
	for t := int64(1); t <= T; t++ {
		R, C, K := l(), l(), uint16(l())

		validCellsToCInR := make([][]uint16, R, R)
		minFromC := make([]uint16, C, C)
		maxFromC := make([]uint16, C, C)
		for r := int64(0); r < R; r++ {
			validCellsToCInR[r] = make([]uint16, C, C)

			for c := int64(0); c < C; c++ {
				validCellsToCInR[r][c] = 1
				v := uint16(l())
				minFromC[c], maxFromC[c] = v, v
				for f := c - 1; f >= 0; f-- {
					minFromC[f], maxFromC[f] = min(minFromC[f], v), max(maxFromC[f], v)
					if maxFromC[f]-minFromC[f] <= K {
						validCellsToCInR[r][c] = uint16(c - f + 1)
					}
				}
			}
		}
		best := R
		type validCellsSince struct {
			validCells uint16;
			since      int64
		}
		history := make([]validCellsSince, 0, 0)
		for c := int64(0); c < C; c++ {
			history = append(
				history[:0],
				validCellsSince{validCellsToCInR[0][c], 0},
			)
			for r := int64(1); r < R; r++ {
				v := validCellsToCInR[r][c]
				hNext := validCellsSince{v, r}
				for len(history) > 0 && history[len(history)-1].validCells > v {
					h := history[len(history)-1]
					best = max64(int64(h.validCells)*(r-h.since), best)
					hNext.since = h.since
					history = history[:len(history)-1]
					if len(history) > 0 {
						h = history[len(history)-1]
					} else {
						break
					}
				}
				history = append(history, hNext)
			}
			for len(history) > 0 {
				h := history[len(history)-1]
				best = max64(int64(h.validCells)*(R-h.since), best)
				history = history[:len(history)-1]
			}
		}
		p("Case #%d: %d\n", t, best)
	}

	if false {
		d("%v", []interface{}{s, l, ui, f, d, p, pf})
	}
}

func min(a, b uint16) uint16 {
	if a < b {
		return a
	};
	return b
}
func max(a, b uint16) uint16 {
	if a > b {
		return a
	};
	return b
}
func max64(a, b int64) int64 {
	if a > b {
		return a
	};
	return b
}
func min64(a, b int64) int64 {
	if a < b {
		return a
	};
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
	//noinspection GoBoolExpressions
	if !DEBUG {
		return
	}
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
