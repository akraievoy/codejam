package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(j Jam) {
	size := int32(j.Int())
	infinite := false
	sum := int32(0)
	prevPrevFigure := int64(0)
	prevFigure := int64(0)
	for i := int32(0); i < size; i++ {
		figure := j.Int()
		switch figure {
		case 1: //	circle
			switch prevFigure {
			case 1:
				infinite = true
			case 2:
				sum += 3
			case 3:
				sum += 4
			}
		case 2: //	triangle
			switch prevFigure {
			case 1:
				if prevPrevFigure == 3 {
					sum += 2
				} else {
					sum += 3
				}
			case 2:
				infinite = true
			case 3:
				infinite = true
			}
		case 3: //	square
			switch prevFigure {
			case 1:
				sum += 4
			case 2:
				infinite = true
			case 3:
				infinite = true
			}
		}
		if infinite {
			break
		}

		prevPrevFigure = prevFigure
		prevFigure = figure
	}
	if infinite {
		j.P("Infinite\n")
	} else {
		j.P("Finite\n")
		j.P("%d\n", sum)
	}
}

func main() {
	jam, closeFunc := JamNew()
	defer closeFunc()
	solve(jam)
}

type Jam interface {
	Scanner() *bufio.Scanner
	Writer() *bufio.Writer
	Close()

	Str() string
	Int() int64
	Float() float64

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

	var writer = bufio.NewWriter(os.Stdout)
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

func (j *jam) Int() int64 {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	res, err := strconv.ParseInt(j.sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprintf(os.Stderr, "scanned %d", res)
	_ = os.Stderr.Sync()

	return res
}

func (j *jam) Float() float64 {
	j.sc.Scan()
	res, err := strconv.ParseFloat(j.sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
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
