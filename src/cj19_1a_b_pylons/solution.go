package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fill(R, C int64, p func(ri, ci int64)) {
	for c := int64(0); c < C; c++ {
		for r := int64(0); r < R; r++ {
			rs := r
			cs := c
			if R == 4 && (c+2 < C && c%2 > 0) {
				rs = r ^ 1
			}
			if R == C && R % 2 > 0 || C == R + 2 {
				if c == C - 2 {
					cs = C - 1
				} else if c == C - 1 {
					cs = C - 2
				}
			}
			if rs%2 == 0 {
				p(rs, (cs+2)%C)
			} else {
				p(rs, cs)
			}
		}
	}
}

func solveSequential(j Jam) {
	T := j.Ri()
	for t := int64(0); t < T; t++ {
		R, C := j.Ri(), j.Ri()
		transpose := false
		if R > C {
			transpose = true
			R, C = C, R
		}

		var p = func(ri, ci int64) {
			if transpose {
				j.W("%d %d\n", ci+1, ri+1)
			} else {
				j.W("%d %d\n", ri+1, ci+1)
			}
		}

		if R == 2 {
			if C < 5 {
				j.W("Case #%d: IMPOSSIBLE\n", 1+t)
			} else {
				j.W("Case #%d: POSSIBLE\n", 1+t)
				fill(R, C, p)
			}
		} else if R == 3 {
			if C == 3 {
				j.W("Case #%d: IMPOSSIBLE\n", 1+t)
			} else {
				j.W("Case #%d: POSSIBLE\n", 1+t)
				fill(R, C, p)
			}
		} else {
			j.W("Case #%d: POSSIBLE\n", 1+t)
			fill(R, C, p)
		}
	}
}

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) > 1 {
		reader, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := reader.Close(); err != nil {
				panic(err)
			}
		}()
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	defer func() {
		if err := writer.Flush(); err != nil {
			panic(err)
		}
	}()

	solveSequential(&jam{scanner, writer})
}

type Jam interface {
	Sc() *bufio.Scanner
	Wr() *bufio.Writer

	Rs() string
	Ri() int64
	Rf() float64

	W(format string, values ...interface{})
}

type jam struct {
	sc *bufio.Scanner
	wr *bufio.Writer
}

func (j *jam) Sc() *bufio.Scanner {
	return j.sc
}

func (j *jam) Wr() *bufio.Writer {
	return j.wr
}

func (j *jam) Rs() string {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	return j.sc.Text()
}

func (j *jam) Ri() int64 {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	res, err := strconv.ParseInt(j.sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	return res
}

func (j *jam) Rf() float64 {
	j.sc.Scan()
	res, err := strconv.ParseFloat(j.sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (j *jam) W(format string, values ...interface{}) {
	out := fmt.Sprintf(format, values...)
	_, err := j.wr.WriteString(out)
	if err != nil {
		panic(err)
	}
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
