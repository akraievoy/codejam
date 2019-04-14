package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solveSequential(j Jam) {
	T := j.Ri()
	for t := int64(0); t < T; t++ {
		N := j.Ri();
		W := make([]string, 0, N)
		rhymingSuffixLen := make([]int, N, N)
		maxLen := 1
		for i := int64(0); i < N; i++ {
			w := j.Rs()
			W = append(W, w)
			maxLen = max(maxLen, len(w))
			rhymingSuffixLen[i] = -1
		}

		for suffixLen := maxLen - 1; suffixLen > 0; suffixLen-- {
			rhymesToW := make(map[string][]int)
			for i, w := range W {
				if rhymingSuffixLen[i] >= 0 || len(w) < suffixLen {
					continue
				}
				suffix := w[len(w)-suffixLen:]
				ws, prs := rhymesToW[suffix]
				if !prs {
					ws = make([]int, 0, N)
				}
				ws = append(ws, i)
				rhymesToW[suffix] = ws
			}

			for _, ws := range rhymesToW {
				if len(ws) > 1 {
					w0, w1 := ws[0], ws[1]
					rhymingSuffixLen[w0] = suffixLen
					rhymingSuffixLen[w1] = suffixLen
				}
			}
		}

		rhymingCount := 0
		for _, rsl := range rhymingSuffixLen {
			if rsl > 0 {
				rhymingCount += 1
			}
		}

		j.W("Case #%d: %d\n", 1+t, rhymingCount)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}