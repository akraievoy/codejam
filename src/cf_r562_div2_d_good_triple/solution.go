package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF
	//	Live Templates: for0l for1l for0ui for1ui forr vl0 vln vui0 vuin ; Casts: l i

	noTriples := make(map[string]bool)
	noTriples["00"], noTriples["01"], noTriples["10"], noTriples["11"] = true, true, true, true

	added := true
	ln := 2
	for added {
		added = false

		for k := range noTriples {
			if len(k) != ln {
				continue
			}
			for _, suff := range [2]rune{'0', '1'} {
				test := k + string([]rune{suff})
				noTriplesInTest := true
				for k := 1; 2*k < len(test) && noTriplesInTest; k++ {
					for x := 0; x < len(test) - 2*k && noTriplesInTest; x++ {
						if test[x] == test[x+k] && test[x+k] == test[x+2*k] {
							noTriplesInTest = false
						}
					}
				}
				if noTriplesInTest {
					//d("%s ", test)
					noTriples[test] = true
					added = true
				}
			}
		}
		ln += 1
		//d("\n")
	}

	noTriples[""], noTriples["0"], noTriples["1"] = true, true, true

	str := s()
	pairs := uint64(len(str))*uint64(len(str)-1)/2
	//d("INIT: %d\n", pairs)
	buf := ""
	for _, r := range []rune(str) {
		buf = buf + string([]rune{r})
		//d("%s %s", string([]rune{r}), buf)
		for !noTriples[buf] {
			buf = buf[1:]
			//d(" :%s", buf)
		}
		if noTriples[buf] {
			//d(" -%d", len(buf)-1)
			pairs -= uint64(len(buf)-1)
		}
		//d("\n")
	}

	p("%d", pairs)

	if false { d("%v", []interface{}{s, l, ui, f, d, p, pf}) }
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

