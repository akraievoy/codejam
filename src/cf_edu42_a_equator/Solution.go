package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type caseInput struct {
	nums []int16
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	size := readInt(scanner)
	nums := make([]int16, size)
	for i := range nums {
		nums[i] = int16(readInt(scanner))
	}
	in := caseInput{nums}
	return in
}

type caseOutput struct {
	equat int
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", out.equat)
}

func solveCase(in caseInput) caseOutput {
	sum := 0
	for _, v := range in.nums {
		sum += int(v)
	}
	sum2 := int(0)
	for i, v := range in.nums {
		sum2 += int(v)
		if sum2 * 2 >= sum {
			return caseOutput{i + 1}
		}
	}
	return caseOutput{len(in.nums)}
}

//	everything below is reusable boilerplate
func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	writeCaseOutput(writer, solveCase(readCaseInput(scanner)))
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

	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solveSequential(scanner, writer)
}

func readInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func readInt(sc *bufio.Scanner) int {
	return int(readInt64(sc))
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
