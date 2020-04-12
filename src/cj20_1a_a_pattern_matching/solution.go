package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DEBUG = false

type patt struct {
	pref string
	inf []string
	suff string

}

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF
	//	Live Templates: for0l for1l for0ui for1ui forr vl0 vln vui0 vuin ; Casts: l i

	T := l()
	d("%d tests", T)
	for t := int64(1); t <= T; t++ {
		n := l()
		patts := make([]patt, 0, n)
		for i := int64(0); i < n; i++ {
			p := s()
			comps := strings.Split(p, "*")
			patts = append(patts, patt{comps[0], comps[1:len(comps)-1], comps[len(comps)-1]})
		}
		ans := &patt{"", []string{}, ""}
		for _, p := range patts {
			if len(p.pref) >= len(ans.pref) {
				if strings.HasPrefix(p.pref, ans.pref) {
					ans.pref = p.pref
				} else {
					ans = nil
					break // no joy
				}
			} else {
				if strings.HasPrefix(ans.pref, p.pref) {
					// we're cool
				} else {
					ans = nil
					break // no joy
				}
			}
			if len(p.suff) >= len(ans.suff) {
				if strings.HasSuffix(p.suff, ans.suff) {
					ans.suff = p.suff
				} else {
					ans = nil
					break // no joy
				}
			} else {
				if strings.HasSuffix(ans.suff, p.suff) {
					// we're cool
				} else {
					ans = nil
					break // no joy
				}
			}
			ans.inf = append(ans.inf, p.inf...)
			d("%v\n", ans)
		}

		if ans == nil {
			p("Case #%d: %s\n", t, "*")
		} else {
			ansStr := ans.pref + strings.Join(ans.inf, "") + ans.suff
			p("Case #%d: %s\n", t, ansStr)
		}
	}

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

