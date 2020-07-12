package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DEBUG = true

func solveAll(jm Jam) {
	i8, pf := jm.ScanI8, jm.Printf

	T := i8()
	for t := int64(1); t <= T; t++ {
		n, a, b := i8(), i8(), i8()

		if n > 100 {
			panic("not today")
		}

		parent := make([]int64, n, n)
		parent[0] = -1
		for i := int64(1); i < n; i++ {
			parent[i] = i8() - 1
		}

		paintedTotal := int64(0)
		painted := make([]int64, n, n)
		for aStart := int64(0); aStart < n; aStart++ {
			for bStart := int64(0); bStart < n; bStart++ {
				mark := 1 + aStart * n + bStart
				paintedNow := int64(0)
				aNode, aPos := aStart, int64(0)
				for aNode >= 0 {
					if aPos % a == 0 && painted[aNode] != mark {
						painted[aNode] = mark
						paintedNow += 1
					}
					aNode = parent[aNode]
					aPos += 1
				}
				bNode, bPos := bStart, int64(0)
				for bNode >= 0 {
					if bPos % b == 0 && painted[bNode] != mark {
						painted[bNode] = mark
						paintedNow += 1
					}
					bNode = parent[bNode]
					bPos += 1
				}
				paintedTotal += paintedNow
			}
		}

		pf("Case #%d: %.6f\n", t, float64(paintedTotal) / float64(n*n))
	}
}

//	works for one painter prob estimation only ((:trolled:))
//		childrenMap := make(map[int64][]int64)
//		for i := int64(1); i < n; i++ {
//			parent[i] = i8() - 1
//			children, prs := childrenMap[parent[i]]
//			if !prs {
//				children = make([]int64, 0, 2)
//			}
//			children = append(children, i)
//			childrenMap[parent[i]] = children
//		}
//
//		levelStats := make([]int64, n)
//		var countLevel func(int64, int64)
//		countLevel = func(at, level int64) {
//			levelStats[level] += 1
//			for _, child := range childrenMap[at] {
//				countLevel(child, level+1)
//			}
//		}
//		countLevel(0, 0)
//
//		d(levelStats)
//
//		paintedAtLevel := int64(0)
//		paintedTotal := int64(0)
//		for level, count := range levelStats {
//			if count == 0 {
//				break
//			}
//			if int64(level) % a == 0 || int64(level) % b == 0 {
//				paintedAtLevel += 1
//			}
//			d(level, count, paintedAtLevel)
//			paintedTotal += paintedAtLevel * count
//		}

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

