package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type caseInput struct {
	index int64
	// FIXME test case input structure
	nums []int16
}

//	FIXME read the input
func readCaseInput(j Jam, index int64) caseInput {
	size := int16(j.Ri())
	nums := make([]int16, size)
	for i := range nums {
		nums[i] = int16(j.Ri())
	}
	in := caseInput{index, nums}
	return in
}

type caseOutput struct {
	index int64
	// FIXME test case output structure
	sum int32
}

func writeCaseOutput(j Jam, out caseOutput) {
	//	FIXME write the out
	j.W("Case #%d: %d\n", 1+out.index, out.sum)
}

func solveCase(in caseInput) caseOutput {
	// FIXME actual solution
	sum := int32(0)
	for _, v := range in.nums {
		sum += int32(v)
	}
	return caseOutput{in.index, sum}
}

//	everything below is reusable boilerplate
func solveSequential(j Jam) {
	caseCount := j.Ri()
	for index := int64(0); index < caseCount; index++ {
		writeCaseOutput(j, solveCase(readCaseInput(j, index)))
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