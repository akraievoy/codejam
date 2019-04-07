package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type caseInput struct {
	index int
	// FIXME test case input structure
	nums []int16
}

//	FIXME read the input
func readCaseInput(scanner *bufio.Scanner, index int) caseInput {
	size := int16(readInt(scanner))
	nums := make([]int16, size)
	for i := range nums {
		nums[i] = int16(readInt(scanner))
	}
	in := caseInput{index, nums}
	return in
}

type caseOutput struct {
	index int
	// FIXME test case output structure
	sum int32
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	//	FIXME write the out
	writef(writer, "Case #%d: %d\n", 1+out.index, out.sum)
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
func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseCount := readInt(scanner)
	for index := 0; index < caseCount; index++ {
		writeCaseOutput(writer, solveCase(readCaseInput(scanner, index)))
	}
}

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) > 1 {
		reader, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer reader.Close()
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solveSequential(scanner, writer)
}

//	GoLang shorthand methods for I/O go below
//	TODO wipe unused methods before submitting

func readString(sc *bufio.Scanner) string {
	if !sc.Scan() {
		panic("failed to scan next token")
	}

	return sc.Text()
}

func readInt64(sc *bufio.Scanner) int64 {
	if !sc.Scan() {
		panic("failed to scan next token")
	}

	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func readInt(sc *bufio.Scanner) int {
	return int(readInt64(sc))
}

func readFloat64(sc *bufio.Scanner) float64 {
	sc.Scan()
	res, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
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