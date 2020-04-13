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
	d("%d tests", T)
	for t := int64(1); t <= T; t++ {
		size := uint64(l())
		trace, dupedRows, dupedCols := uint64(0), int64(0), int64(0)
		dupedRow, dupedCol := NewBitSet(size), NewBitSet(size)
		seenInRow, seenInCol := make([]BitSet, size), make([]BitSet, size)
		for i := uint64(0); i < size; i++ {
			seenInRow[i], seenInCol[i] = NewBitSet(size), NewBitSet(size)
		}
		for row := uint64(0); row < size; row++ {
			for col := uint64(0); col < size; col++ {
				elem := uint64(l())
				if row == col {
					trace += elem
				}
				if seenInRow[row].IsSet(elem-1) {
					if !dupedRow.IsSet(row)	 {
						dupedRows +=1
					}
					dupedRow.Set(row)
				}
				seenInRow[row].Set(elem-1)
				if seenInCol[col].IsSet(elem-1) {
					if !dupedCol.IsSet(col) {
						dupedCols += 1
					}
					dupedCol.Set(col)
				}
				seenInCol[col].Set(elem-1)
			}
		}
		p("Case #%d: %d %d %d\n", t, trace, dupedRows, dupedCols)
	}

	if false { d("%v", []interface{}{s, l, ui, f, d, p, pf}) }
}

type BitSetElem uint64
const BitSetElemBits = 64

// BitSet is a set of bits that can be set, cleared and queried.
type BitSet []BitSetElem

func NewBitSet(bitlen uint64) BitSet {
	return make([]BitSetElem, (bitlen + BitSetElemBits - 1)/BitSetElemBits)
}

// Set ensures that the given bit is set in the BitSet.
func (s *BitSet) Set(i uint64) {
	(*s)[i/BitSetElemBits] |= 1 << (i % BitSetElemBits)
}

// Clear ensures that the given bit is cleared (not set) in the BitSet.
func (s *BitSet) Clear(i uint64) {
	if len(*s) >= int(i/BitSetElemBits+1) {
		(*s)[i/BitSetElemBits] &^= 1 << (i % BitSetElemBits)
	}
}

// IsSet returns true if the given bit is set, false if it is cleared.
func (s *BitSet) IsSet(i uint64) bool {
	return (*s)[i/BitSetElemBits]&(1<<(i%BitSetElemBits)) != 0
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

