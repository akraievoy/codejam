package main

import (
	"bufio"
	"fmt"
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

	T, B := i8(), i8()
	for t := int64(1); t <= T; t++ {
		df("test %d of %d\n", t, T)

		bits := make([]int64, B, B)
		for i := range bits {
			bits[i] = -1 // unknown
		}
		samePos := int64(-1)
		diffPos := int64(-1)

		completedQueries := 0
		knownBits := int64(0)

		for knownBits < B || knownBits == B && completedQueries%10 < 2 {
			firstUnknownBit := int64(0)
			for i := int64(0); i < B; i++ {
				pos := TernI8(
					i%2 == 0,
					i/2,
					B-i/2-1,
				)
				if bits[pos] < 0 {
					firstUnknownBit = pos
					break
				}
			}
			queryPos :=
				TernI8(
					completedQueries%10 == 0 && samePos >= 0,
					samePos,
					TernI8(
						completedQueries%10 == 1 && diffPos >= 0,
						diffPos,
						firstUnknownBit,
					),
				)

			d("query to", queryPos+1)
			f(queryPos + 1)
			v := i8()
			d("received", v)

			if completedQueries%10 == 0 {
				for i := int64(0); i < B/2; i++ {
					if bits[i] >= 0 && bits[B-i-1] < 0 {
						knownBits -= 1
						bits[i] = -1
					} else if bits[B-i-1] >= 0 && bits[i] < 0 {
						knownBits -= 1
						bits[B-i-1] = -1
					}
				}
				if samePos >= 0 && bits[samePos] != v {
					d("complement busted, recovering equal pairs")
					for i := int64(0); i < B/2; i++ {
						if bits[i] >= 0 && bits[B-i-1] == bits[i] {
							bits[i] = 1 - bits[i]
							bits[B-i-1] = 1 - bits[B-i-1]
						}
					}
					if B%2 > 0 && bits[B/2] >= 0 {
						bits[B/2] = 1 - bits[B/2]
					}
				}
			} else if completedQueries%10 == 1 {
				if diffPos >= 0 && bits[diffPos] != v {
					d("complement or reversal (but not both) busted, recovering unequal pairs")
					for i := int64(0); i < B/2; i++ {
						if bits[i] >= 0 && bits[B-i-1] >= 0 && bits[B-i-1] != bits[i] {
							bits[i] = 1 - bits[i]
							bits[B-i-1] = 1 - bits[B-i-1]
						}
					}
				}
			}

			if bits[queryPos] < 0 {
				knownBits += 1
				bits[queryPos] = v
			}

			if samePos == -1 && bits[queryPos] == bits[B-queryPos-1] {
				samePos = queryPos
			}
			if diffPos == -1 && bits[B-queryPos-1] >= 0 && bits[queryPos] == 1-bits[B-queryPos-1] {
				diffPos = queryPos
			}

			completedQueries += 1
			d("completedQueries =", completedQueries, "knownBits =", knownBits, "bits =", bits)
		}

		res := ""
		for _, b := range bits {
			res += TernS(b > 0, "1", "0")
		}
		d("res =", res)
		f(res)
		if s() != "Y" {
			panic("HOUSTON 1")
		}
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
	if len(os.Args) > 1 {
		panic("running with input file path is not supported")
	}

	var scanner = bufio.NewScanner(os.Stdin)
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

func TernS(b bool, t, f string) string {
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
