package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(j Jam) {
	T := int16(j.Int())
	for t := int16(0); t < T; t++ {
		in := j.Str()
		runeCount := make([]int, 26)
		for _, r := range []rune(in) {
			rIdx := int(r - 'a')
			runeCount[rIdx] += 1
		}
		presentRunes := make([]int, 0, 26)
		for r := range runeCount {
			if runeCount[r] > 0 {
				presentRunes = append(presentRunes, r)
			}
		}
		result := make([]rune, 0)
		seed := presentRunes[0]
		for i := 0; i < runeCount[seed]; i++ {
			result = append(result, rune('a'+seed))
		}
		presentRunes = presentRunes[1:]
		left, right := seed, seed
		retryRunes := make([]int, 0, 26)
		for _, r := range presentRunes {
			if right+1 != r {
				for i := 0; i < runeCount[r]; i++ {
					result = append(result, rune('a'+r))
				}
				right = r
			} else if left+1 != r {
				for i := 0; i < runeCount[r]; i++ {
					result = append([]rune{rune('a' + r)}, result...)
				}
				left = r
			} else {
				retryRunes = append(retryRunes, r)
			}
		}
		valid := true
		for _, r := range retryRunes {
			if abs(right-r) != 1 {
				for i := 0; i < runeCount[r]; i++ {
					result = append(result, rune('a'+r))
				}
				right = r
			} else if abs(left-r) != 1 {
				for i := 0; i < runeCount[r]; i++ {
					result = append([]rune{rune('a' + r)}, result...)
				}
				left = r
			} else {
				valid = false
				break
			}
		}
		if !valid {
			j.P("No answer\n")
		} else {
			j.P("%s\n", string(result))
		}
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

//	GoLang shorthand methods for math go below
//	TODO wipe unused methods before submitting

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
