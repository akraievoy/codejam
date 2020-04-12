package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var DEBUG = false

type pos struct{ r, c int64 }

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF
	//	Live Templates: for0l for1l for0ui for1ui forr vl0 vln vui0 vuin ; Casts: l i

	T := l()
	dU, dL, dD, dR := 0, 1, 2, 3
	dirs := []int{dU, dL, dD, dR}
	for t := int64(1); t <= T; t++ {
		d("test %d of %d\n", t, T)
		R, C := l(), l()
		interest := int64(0)
		scores := sliceL2D(R, C)
		comp := sliceL3D(R, C, int64(len(dirs)))
		qIn, qOut, qElim := make(map[pos]bool), make(map[pos]bool), make(map[pos]bool)
		for r := int64(0); r < R; r++ {
			for c := int64(0); c < C; c++ {
				scores[r][c] = l()

				interest += scores[r][c]
				comp[r][c][dU] = tern(r > 0, r-1, -1)
				comp[r][c][dD] = tern(r < R-1, r+1, - 1)
				comp[r][c][dL] = tern(C > 0, c-1, - 1)
				comp[r][c][dR] = tern(c < C-1, c+1, - 1)
			}
		}

		interestTotal := int64(0)

		for r := int64(0); r < R; r++ {
			for c := int64(0); c < C; c++ {
				qIn[pos{r, c}] = true
			}
		}

		for len(qIn) > 0 {
			interestTotal += interest
			d("next round, interest total %d, interest remaining %d, len(qIn) == %d\n", interestTotal, interest, len(qIn))
			for p := range qIn {
				neighbourScores := int64(0)
				neighbourCount := int64(0)
				for _, d := range dirs {
					rOrC := comp[p.r][p.c][d]
					if rOrC < 0 {
						continue
					}
					if d == dU || d == dD {
						neighbourScores += scores[rOrC][p.c]
					} else {
						neighbourScores += scores[p.r][rOrC]
					}
					neighbourCount += 1
				}

				if scores[p.r][p.c]*neighbourCount >= neighbourScores {
					continue
				}
				qElim[p] = true
			}
			for p := range qElim {
				d("eliminating %d %d %d\n", p.r, p.c, scores[p.r][p.c])
				if scores[p.r][p.c] == 0 {
					continue
				}
				interest -= scores[p.r][p.c]
				scores[p.r][p.c] = 0
				for _, d := range dirs {
					dOppo := (d + 2) % 4
					rOrC := comp[p.r][p.c][d]
					if rOrC < 0 {
						continue
					}
					if d == dU || d == dD {
						comp[rOrC][p.c][dOppo] = comp[p.r][p.c][dOppo]
						qOut[pos{rOrC, p.c}] = true
					} else {
						comp[p.r][rOrC][dOppo] = comp[p.r][p.c][dOppo]
						qOut[pos{p.r, rOrC}] = true
					}
				}
			}

			qIn = qOut
			qOut, qElim = make(map[pos]bool), make(map[pos]bool)
		}

		p("Case #%d: %d\n", t, interestTotal)
	}

	if false {
		d("%v", []interface{}{s, l, ui, f, d, p, pf})
	}
}

func sliceL2D(R, C int64) [][]int64 {
	res := make([][]int64, 0, R)
	for r := int64(0); r < R; r++ {
		res = append(res, make([]int64, C, C))
	}
	return res
}

func sliceL3D(L, R, C int64) [][][]int64 {
	res := make([][][]int64, L, L)
	for l := int64(0); l < L; l++ {
		res[l] = make([][]int64, 0, R)
		for r := int64(0); r < R; r++ {
			res[l] = append(res[l], make([]int64, C, C))
		}
	}
	return res
}

func tern(b bool, t, f int64) int64 {
	if b {
		return t
	}
	return f
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
