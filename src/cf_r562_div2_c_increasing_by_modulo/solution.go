package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF
	//	Live Templates: for0l for1l for0ui for1ui forr vl0 vln vui0 vuin ; Casts: l i

	n,m := ui(), ui()
	a := make([]uint32, n, n)
	for i := uint32(0); i < n; i++ {
		a[i] = ui()
	}

	p(
		"%d\n",
		sort.Search(
			int(m),
			func(probeInt int) bool {
				probe := uint32(probeInt)
				//d("PROBE=%d\n", probe)
				minBound := ti((a[0] + probe) >= m, 0, a[0])
				//d("%d:%d ", a[0], minBound)
				for i := uint32(1); i < n; i++ {
					if minBound > a[i] + probe  {
						//d("%d:FALSE\n", a[i])
						return false
					}
					if (a[i] + probe) >= m && (a[i] + probe) % m >= minBound {
						// do nothing
					} else {
						minBound = maxi(a[i], minBound)
					}
					//d("%d:%d ", a[i], minBound)
				}
				//d(" TRUE\n")
					return true
			},
		),
	)

	if false { d("%v", []interface{}{s, l, ui, f, d, p, pf}) }
}

//	TODO wipe unused shorthand methods for math
type Uint32Sort []uint32
func (s Uint32Sort) Len() int           { return len(s) }
func (s Uint32Sort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Uint32Sort) Less(i, j int) bool { return s[i] < s[j] }

func t(b bool, t, f int64) int64 { if b {return t};return f }
func min(a, b int64) int64 { if a < b {return a}; return b }
func max(a, b int64) int64 { if a > b {return a}; return b }
func abs(a int64) int64 { if a < 0 { return -a }; return a }
func ti(b bool, t, f uint32) uint32 { if b {return t};return f }
func mini(a, b uint32) uint32 { if a < b { return a }; return b }
func maxi(a, b uint32) uint32 { if a > b { return a }; return b }

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

