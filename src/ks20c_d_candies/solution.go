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
		n, q := i8(), i8()
		cell := MaxI8(1, int64(math.Sqrt(float64(n))))
		d("cell", cell)
		cells := n / cell
		arr := make([]int64, n, n)
		for i := int64(0); i < n; i++ {
			arr[i] = i8()
		}
		signedSums := make([]int64, cells, cells)
		weightedSums := make([]int64, cells, cells)
		for c := int64(0); c < cells; c++ {
			for i := int64(0); i < cell; i++ {
				sign := TernI8(i%2 == 0, 1, -1)
				signedSums[c] += arr[c*cell+i] * sign
				weightedSums[c] += arr[c*cell+i] * sign * (i + 1)
			}
		}
		d(signedSums)
		d(weightedSums)
		sweetenessSum := int64(0)
		for qi := int64(0); qi < q; qi++ {
			query := s()
			if query == "U" {
				pos := i8() - 1
				newVal := i8()
				oldVal := arr[pos]
				arr[pos] = newVal
				c := pos / cell
				if c < cells {
					i := pos % cell
					sign := TernI8(i%2 == 0, 1, -1)
					signedSums[c] += (newVal - oldVal) * sign
					weightedSums[c] += (newVal - oldVal) * sign * (i + 1)
				}
				d(signedSums)
				d(weightedSums)
			} else if query == "Q" {
				querySweeteness, start, end := int64(0), i8()-1, i8()-1
				for i := start; i <= end; {
					relPos := i - start + 1
					mult := TernI8(relPos%2 == 1, relPos, -relPos)
					if i%cell == 0 && i+cell - 1 <= end {
						c := i / cell
						querySweeteness +=
							TernI8(relPos%2 == 1, weightedSums[c], -weightedSums[c]) +
								(mult+TernI8(relPos%2==1, -1, 1))*signedSums[c]
						//d("q(", start, ",", end, ") @", i, " :cell ", TernI8(relPos%2 == 1, weightedSums[c], -weightedSums[c]),
						//	(mult+TernI8(relPos%2==1, -1, 1))*signedSums[c], " --> ", querySweeteness)
						i += cell
					} else {
						querySweeteness += mult * arr[i]
						//d("q(", start, ",", end, ") @", i, " :elem ", arr[i], "*", mult, " --> ", querySweeteness)
						i += 1
					}
				}
				sweetenessSum += querySweeteness
			} else {
				panic("houston")
			}
		}
		pf("Case #%d: %d\n", t, sweetenessSum)
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

func MaxI8(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func TernI8(b bool, t, f int64) int64 {
	if b {
		return t
	}
	return f
}

