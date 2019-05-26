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

	dc, dr := make(map[rune]int32), make(map[rune]int32)
	dirs := "NSEW"
	dc['N'], dc['S'], dc['E'], dc['W'] = 0, 0, 1, -1
	dr['N'], dr['S'], dr['E'], dr['W'] = -1, 1, 0, 0

	type pos struct{ r, c int32 }
	type posdir struct {
		p pos;
		d rune
	}

	T := l()
	d("%d tests", T)
	for t := int64(1); t <= T; t++ {
		skip := make(map[posdir]pos)

		n, _ /*r*/, _ /*c*/, sr, sc := l(), l(), l(), int32(l()), int32(l())
		poses := make([]pos, 0, n)
		route := s()

		for _, dir := range []rune(dirs) {
			next := pos{sr + dr[dir], sc + dc[dir]}
			nextSkip, prs := skip[posdir{next, dir}]
			if prs {
				skip[posdir{pos{sr, sc}, dir}] = nextSkip
			} else {
				skip[posdir{pos{sr, sc}, dir}] = next
			}
		}
		for _, dir := range []rune(route) {
			sr += dr[dir]
			sc += dc[dir]

			poses = poses[:0]
			for {
				nextSkip, prs := skip[posdir{pos{sr, sc}, dir}]
				if !prs {
					break
				}
				poses = append(poses, pos{sr, sc})
				sr, sc = nextSkip.r, nextSkip.c
			}
			for _, prevPos := range poses {
				skip[posdir{prevPos, dir}] = pos{sr, sc}
			}

			for _, dir := range []rune(dirs) {
				next := pos{sr + dr[dir], sc + dc[dir]}
				nextSkip, prs := skip[posdir{next, dir}]
				if prs {
					skip[posdir{pos{sr, sc}, dir}] = nextSkip
				} else {
					skip[posdir{pos{sr, sc}, dir}] = next
				}
			}
		}
		p("Case #%d: %d %d\n", t, sr, sc)
	}

	if false {
		d("%v", []interface{}{s, l, ui, f, d, p, pf})
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
