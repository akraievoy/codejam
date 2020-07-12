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
	i8, pf := jm.ScanI8, jm.Printf

	T := i8()
	for t := int64(1); t <= T; t++ {
		n, q := i8(), i8()
		if n > 1000 || q > 1000 {
			panic("not today")
		}
		drs := make([]int64, n-1, n-1)
		for i := int64(0); i < n-1; i++ {
			drs[i] = i8()
		}
		pf("Case #%d: ", t)
		for i := int64(0); i < q; i++ {
			s, k := i8()-1, i8()-1
			l, r := s, s
			ans := s
			for k > 0 {
				if l == 0 {
					r += 1
					ans = r
				} else if r == n-1 {
					l -= 1
					ans = l
				} else if drs[l-1] < drs[r] {
					l -= 1
					ans = l
				} else {
					r += 1
					ans = r
				}
				k -= 1
			}
			pf("%d", ans+1)
			if i+1 < q {
				pf(" ")
			}
		}
		pf("\n")
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
