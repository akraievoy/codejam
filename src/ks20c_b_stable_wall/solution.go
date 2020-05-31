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

	T := i8()
	for t := int64(1); t <= T; t++ {
		df("test %d of %d\n", t, T)
		r, c := i8(), i8()
		upperTiles := make([]rune, c, c)
		polyToBelow := make(map[rune]map[rune]bool)
		allPolys := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		for _, poly := range allPolys {
			polyToBelow[poly] = make(map[rune]bool)
		}
		usedPolysMap := make(map[rune]bool)
		for ri := int64(0); ri < r; ri++ {
			row := []rune(s())
			for ci := int64(0); ci < c; ci++ {
				if ri != 0 {
					if upperTiles[ci] != row[ci] {
						polyToBelow[upperTiles[ci]][row[ci]] = true
					}
				}
				upperTiles[ci] = row[ci]
				usedPolysMap[row[ci]] = true
			}
		}
		usedPolys := make([]rune, 0, 26)
		for p := range usedPolysMap {
			usedPolys = append(usedPolys, p)
		}
		possible := true
		placedPolys := make(map[rune]bool)
		polyPlacementOrder := make([]rune, 0, 26)
		for len(polyPlacementOrder) < len(usedPolys) {
			atLeastOnePlaced := false
			for _, poly := range usedPolys {
				if placedPolys[poly] {
					continue
				}
				reqmsMet := true
				for belowReqm := range polyToBelow[poly] {
					if !placedPolys[belowReqm] {
						reqmsMet = false
						break
					}
				}
				if reqmsMet {
					atLeastOnePlaced = true
					placedPolys[poly] = true
					polyPlacementOrder = append(polyPlacementOrder, poly)
				}
			}
			if !atLeastOnePlaced {
				possible = false
				break
			}
		}
		pf("Case #%d: %s\n", t, TernS(possible, string(polyPlacementOrder), "-1"))
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

func TernS(b bool, t, f string) string {
	if b {
		return t
	}
	return f
}

