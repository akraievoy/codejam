package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var DEBUG = true

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF
	//	Live Templates: for0l for1l for0ui for1ui forr vl0 vln vui0 vuin ; Casts: l i

	T := l()
	for t := int64(1); t <= T; t++ {
		n := l()
		d("test %d: target = %d\n", t, n)
		maxRows := uint64(30)

		height := n
		included := NewBitSet(maxRows)

		minHeight := int64(0)
		for pow := int64(maxRows); pow > 0; pow-- {
			if minHeight == 0 {
				if height - (1 << uint64(pow)) >= pow {
					included.Set(uint64(pow))
					minHeight = pow
					height -= (1 << uint64(pow)) - 1
					d("included pow %d, height = %d\n", pow, height)
				} else {
					d("passed on pow %d\n", pow)
				}
			} else {
				if height - (1 << uint64(pow)) >= minHeight {
					included.Set(uint64(pow))
					height -= (1 << uint64(pow)) - 1
					d("included pow %d, height = %d\n", pow, height)
				} else {
					d("passed on pow %d\n", pow)
				}
			}
		}

		p("Case #%d:\n", t)
		sum := int64(0)
		left := true
		for h := int64(0); h < height; h++ {
			incl := false
			if h <= int64(maxRows) && included.IsSet(uint64(h)) {
				incl = true
				if left {
					for c := int64(0); c <= h; c++ {
						d("%d %d\n", h+1, c+1)
						p("%d %d\n", h+1, c+1)
					}
				} else {
					for c := h; c >= int64(0); c-- {
						d("%d %d\n", h+1, c+1)
						p("%d %d\n", h+1, c+1)
					}
				}
				sum += 1 << uint64(h)
				left = !left
			} else {
				if left {
					d("%d %d\n", h+1, 1)
					p("%d %d\n", h+1, 1)
				} else {
					d("%d %d\n", h+1, h+1)
					p("%d %d\n", h+1, h+1)
				}
				sum += 1
			}
			d("%d %v %d\n", h, incl, sum)
		}

		if sum != n {
			panic(fmt.Sprintf("HOUSTON: %v != %v", sum, n))
		}

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

