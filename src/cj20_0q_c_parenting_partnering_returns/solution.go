package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var DEBUG = false

type Event struct {
	AtTime        int64
	ActivityIndex int64
}

type EventSlice []*Event

func (events EventSlice) Len() int {
	return len(events)
}

func (events EventSlice) Swap(i, j int) {
	events[i], events[j] = events[j], events[i]
}

func (events EventSlice) Less(i, j int) bool {
	if events[i].AtTime == events[j].AtTime {
		return events[i].ActivityIndex < events[j].ActivityIndex
	}
	return events[i].AtTime < events[j].AtTime
}

func solveAll(jam Jam) {
	s, l, ui, f, d, p, pf := jam.Str, jam.Long, jam.Int, jam.Float, jam.D, jam.P, jam.PF
	//	Live Templates: for0l for1l for0ui for1ui forr vl0 vln vui0 vuin ; Casts: l i

	T := l()
	d("%d tests", T)
	for t := int64(1); t <= T; t++ {
		size := l()
		cj := make([]rune, size)
		possible := true
		events := make([]*Event, 2*size)
		for i := int64(0); i < size; i++ {
			events[2*i], events[2*i+1] = &Event{l(), i + 1}, &Event{l(), -i - 1}
		}
		sort.Sort(EventSlice(events))
		activitiesScheduled := uint64(0)
		cameronActivity := int64(0)
		jamieActivity := int64(0)
		for _, e := range events {
			if e.ActivityIndex < 0 {
				activitiesScheduled -= 1
				if activitiesScheduled < 0 {
					panic("HOUSTON 0")
				}
				if cameronActivity == abs(e.ActivityIndex) {
					cameronActivity = 0
				} else if jamieActivity == abs(e.ActivityIndex) {
					jamieActivity = 0
				} else {
					panic("HOUSTON 1")
				}
			} else {
				activitiesScheduled += 1
				if activitiesScheduled > 2 {
					possible = false
					break
				}
				if cameronActivity == 0 {
					cameronActivity = e.ActivityIndex
					cj[e.ActivityIndex-1] = 'C'
				} else if jamieActivity == 0 {
					jamieActivity = e.ActivityIndex
					cj[e.ActivityIndex-1] = 'J'
				} else {
					panic("HOUSTON 2")
				}
			}
		}
		p("Case #%d: %s\n", t, ts(possible, string(cj), "IMPOSSIBLE"))
	}

	if false {
		d("%v", []interface{}{s, l, ui, f, d, p, pf})
	}
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
func ts(b bool, t, f string) string {
	if b {
		return t
	}
	return f
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
