package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var DEBUG = false

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF
	//	Live Templates: for0l for1l for0ui for1ui forr vl0 vln vui0 vuin ; Casts: l i

	T := l()
	for t := int64(1); t <= T; t++ {
		d("test %d of %d\n", t, T)
		size := l()
		arr := make([]int64, 0, size)
		for i, e := range arr {
			d("%d %d\n", i, e)
		}
		sum := int64(0)
		for i := int64(0); i < size; i++ {
			sum += l()
		}
		p("Case #%d: %d\n", t, sum)
	}

	if false { d("%v", []interface{}{s, l, ui, f, d, p, pf}) }
}

//	TODO wipe unused shorthand methods for math
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

func sliceUI2D(R, C int64) [][]uint32 {
	res := make([][]uint32, 0, R)
	for r := int64(0); r < R; r++ {
		res = append(res, make([]uint32, C, C))
	}
	return res
}

func sliceUI3D(L, R, C int64) [][][]uint32 {
	res := make([][][]uint32, L, L)
	for l := int64(0); l < L; l++ {
		res[l] = make([][]uint32, 0, R)
		for r := int64(0); r < R; r++ {
			res[l] = append(res[l], make([]uint32, C, C))
		}
	}
	return res
}

func sliceF2D(R, C int64) [][]float64 {
	res := make([][]float64, 0, R)
	for r := int64(0); r < R; r++ {
		res = append(res, make([]float64, C, C))
	}
	return res
}

func sliceF3D(L, R, C int64) [][][]float64 {
	res := make([][][]float64, L, L)
	for l := int64(0); l < L; l++ {
		res[l] = make([][]float64, 0, R)
		for r := int64(0); r < R; r++ {
			res[l] = append(res[l], make([]float64, C, C))
		}
	}
	return res
}

type Uint32Sort []uint32
func (s Uint32Sort) Len() int           { return len(s) }
func (s Uint32Sort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Uint32Sort) Less(i, j int) bool { return s[i] < s[j] }

func tern(b bool, t, f int64) int64 { if b {return t};return f }
func min(a, b int64) int64 { if a < b {return a}; return b }
func max(a, b int64) int64 { if a > b {return a}; return b }
func abs(a int64) int64 { if a < 0 { return -a }; return a }
func ternUI(b bool, t, f uint32) uint32 { if b {return t};return f }
func minUI(a, b uint32) uint32 { if a < b { return a }; return b }
func maxUI(a, b uint32) uint32 { if a > b { return a }; return b }
func ternS(b bool, t, f string) string { if b {return t};return f }

func round(x float64) float64 { //	https://www.cockroachlabs.com/blog/rounding-implementations-in-go/
	const (
		mask     = 0x7FF
		shift    = 64 - 11 - 1
		bias     = 1023

		signMask = 1 << 63
		fracMask = (1 << shift) - 1
		halfMask = 1 << (shift - 1)
		one      = bias << shift
	)

	bits := math.Float64bits(x)
	e := uint(bits>>shift) & mask
	switch {
	case e < bias:
		// Round abs(x)<1 including denormals.
		bits &= signMask // +-0
		if e == bias-1 {
			bits |= one // +-1
		}
	case e < bias+shift:
		// Round any abs(x)>=1 containing a fractional component [0,1).
		e -= bias
		bits += halfMask >> e
		bits &^= fracMask >> e
	}
	return math.Float64frombits(bits)
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
