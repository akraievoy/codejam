package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var DEBUG = false

func solveAll(jm Jam) {
	//	Live Templates: for0i8 for1i8 forr v0i8 vi8 ; Casts: i8
	s, i8, u4, f8,
	df, pf, ff,
	d, p, f :=
		jm.ScanS, jm.ScanI8, jm.ScanU4, jm.ScanF8,
		jm.Debugf, jm.Printf, jm.Flushf,
		jm.Debug, jm.Print, jm.Flush

	T := i8()
	for t := int64(1); t <= T; t++ {
		df("test %d of %d\n", t, T)
		size := i8()
		arr := make([]int64, 0, size)
		for i, e := range arr {
			df("%d %d\n", i, e)
		}
		sum := int64(0)
		for i := int64(0); i < size; i++ {
			sum += i8()
		}
		pf("Case #%d: %d\n", t, sum)
	}

	if false {
		df("%v", []interface{}{s, i8, u4, f8, df, pf, ff, d, p, f})
	}
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

	ScanS() string
	ScanF8() float64
	ScanI8() int64
	ScanU8() uint64
	ScanI4() int32
	ScanU4() uint32

	Debugf(format string, values ...interface{})
	Debug(values ...interface{})
	Printf(format string, values ...interface{})
	Print(values ...interface{})
	Flushf(format string, values ...interface{})
	Flush(values ...interface{})
}

func JamNew() (Jam, func()) {
	var scanner *bufio.Scanner
	if len(os.Args) > 1 {
		reader, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		//	FIXME defer reader.Close()
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var outWriter = bufio.NewWriterSize(os.Stdout, 1024*1024)
	var errWriter = bufio.NewWriterSize(os.Stderr, 1024*1024)
	jam := &jam{scanner, outWriter, errWriter}
	return jam, jam.Close
}

type jam struct {
	inScanner *bufio.Scanner
	outWriter *bufio.Writer
	errWriter *bufio.Writer
}

func (jm *jam) Close() {
	if err := jm.outWriter.Flush(); err != nil {
		panic(err)
	}
	if err := jm.errWriter.Flush(); err != nil {
		panic(err)
	}
}

func (jm *jam) Scanner() *bufio.Scanner {
	return jm.inScanner
}

func (jm *jam) Writer() *bufio.Writer {
	return jm.outWriter
}

func (jm *jam) ScanS() string {
	if !jm.inScanner.Scan() {
		panic("failed to scan next token")
	}

	return jm.inScanner.Text()
}

func (jm *jam) ScanF8() float64 {
	jm.inScanner.Scan()
	res, err := strconv.ParseFloat(jm.inScanner.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (jm *jam) ScanI8() int64 {
	if !jm.inScanner.Scan() {
		panic("failed to scan next token")
	}

	res, err := strconv.ParseInt(jm.inScanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	return res
}

func (jm *jam) ScanU8() uint64 {
	return uint64(jm.ScanI8())
}

func (jm *jam) ScanI4() int32 {
	return int32(jm.ScanI8())
}

func (jm *jam) ScanU4() uint32 {
	return uint32(jm.ScanI8())
}

func (jm *jam) Debugf(format string, values ...interface{}) {
	//noinspection GoBoolExpressions
	if !DEBUG {
		return
	}
	_ /*bytesWritten*/, err := fmt.Fprintf(jm.errWriter, format, values...)
	if err != nil {
		panic(err)
	}
	err = jm.errWriter.Flush()
	if err != nil {
		panic(err)
	}
}

func (jm *jam) Debug(values ...interface{}) {
	jm.Debugf(strings.Repeat("%v ", len(values)-1)+"%v\n", values...)
}

func (jm *jam) Printf(format string, values ...interface{}) {
	_, err := fmt.Fprintf(jm.outWriter, format, values...)
	if err != nil {
		panic(err)
	}
}

func (jm *jam) Print(values ...interface{}) {
	jm.Printf(strings.Repeat("%v ", len(values)-1)+"%v\n", values...)
}

func (jm *jam) Flushf(format string, values ...interface{}) {
	jm.Printf(format, values...)
	if err := jm.outWriter.Flush(); err != nil {
		panic(err)
	}
}

func (jm *jam) Flush(values ...interface{}) {
	jm.Flushf(strings.Repeat("%v ", len(values)-1)+"%v\n", values...)
}

//	TODO the below is typical golang hand-roll-yer-own-max-function boilerplate

func Slice2DI8(R, C int64) [][]int64 {
	res := make([][]int64, 0, R)
	for r := int64(0); r < R; r++ {
		res = append(res, make([]int64, C, C))
	}
	return res
}

func Slice2DU8(R, C int64) [][]uint64 {
	res := make([][]uint64, 0, R)
	for r := int64(0); r < R; r++ {
		res = append(res, make([]uint64, C, C))
	}
	return res
}

func Slice2DI4(R, C int64) [][]int32 {
	res := make([][]int32, 0, R)
	for r := int64(0); r < R; r++ {
		res = append(res, make([]int32, C, C))
	}
	return res
}

func Slice2DU4(R, C int64) [][]uint32 {
	res := make([][]uint32, 0, R)
	for r := int64(0); r < R; r++ {
		res = append(res, make([]uint32, C, C))
	}
	return res
}

func Slice2DF8(R, C int64) [][]float64 {
	res := make([][]float64, 0, R)
	for r := int64(0); r < R; r++ {
		res = append(res, make([]float64, C, C))
	}
	return res
}

func Slice3DI8(L, R, C int64) [][][]int64 {
	res := make([][][]int64, 0, L)
	for l := int64(0); l < L; l++ {
		res = append(res, Slice2DI8(R, C))
	}
	return res
}

func Slice3DU8(L, R, C int64) [][][]uint64 {
	res := make([][][]uint64, 0, L)
	for l := int64(0); l < L; l++ {
		res = append(res, Slice2DU8(R, C))
	}
	return res
}

func Slice3DI4(L, R, C int64) [][][]int32 {
	res := make([][][]int32, 0, L)
	for l := int64(0); l < L; l++ {
		res = append(res, Slice2DI4(R, C))
	}
	return res
}

func Slice3DU4(L, R, C int64) [][][]uint32 {
	res := make([][][]uint32, 0, L)
	for l := int64(0); l < L; l++ {
		res = append(res, Slice2DU4(R, C))
	}
	return res
}

func Slice3DF8(L, R, C int64) [][][]float64 {
	res := make([][][]float64, 0, L)
	for l := int64(0); l < L; l++ {
		res = append(res, Slice2DF8(R, C))
	}
	return res
}

type SortI8 []int64

func (s SortI8) Len() int           { return len(s) }
func (s SortI8) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortI8) Less(i, j int) bool { return s[i] < s[j] }

type SortU8 []uint64

func (s SortU8) Len() int           { return len(s) }
func (s SortU8) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortU8) Less(i, j int) bool { return s[i] < s[j] }

type SortI4 []int32

func (s SortI4) Len() int           { return len(s) }
func (s SortI4) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortI4) Less(i, j int) bool { return s[i] < s[j] }

type SortU4 []uint32

func (s SortU4) Len() int           { return len(s) }
func (s SortU4) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortU4) Less(i, j int) bool { return s[i] < s[j] }

func MinF8(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func MinI8(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func MinU8(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
func MinI4(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
func MinU4(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func MaxF8(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
func MaxI8(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func MaxU8(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
func MaxI4(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
func MaxU4(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func AbsF8(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
func AbsI8(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
func AbsU8(a uint64) uint64 {
	if a < 0 {
		return -a
	}
	return a
}
func AbsI4(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}
func AbsU4(a uint32) uint32 {
	if a < 0 {
		return -a
	}
	return a
}

func TernS(b bool, t, f string) string {
	if b {
		return t
	}
	return f
}
func TernF8(b bool, t, f float64) float64 {
	if b {
		return t
	}
	return f
}
func TernI8(b bool, t, f int64) int64 {
	if b {
		return t
	}
	return f
}
func TernU8(b bool, t, f uint64) uint64 {
	if b {
		return t
	}
	return f
}
func TernI4(b bool, t, f int32) int32 {
	if b {
		return t
	}
	return f
}
func TernU4(b bool, t, f uint32) uint32 {
	if b {
		return t
	}
	return f
}

func Round(x float64) float64 {
	//	https://www.cockroachlabs.com/blog/rounding-implementations-in-go/
	const (
		mask  = 0x7FF
		shift = 64 - 11 - 1
		bias  = 1023

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