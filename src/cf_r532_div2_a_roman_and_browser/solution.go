package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type CaseInput struct {
	K int64
	Tabs []int64
}

func readCaseInput(scanner *bufio.Scanner) CaseInput {
	size := readInt64(scanner)
	k := readInt64(scanner)
	tabs := make([]int64, size)
	for i := range tabs {
		tabs[i] = readInt64(scanner)
	}
	return CaseInput{k, tabs}
}

type CaseOutput struct {
	MaxDiff int64
}

func writeCaseOutput(writer *bufio.Writer, out CaseOutput) {
	writef(writer, "%d\n", abs64(out.MaxDiff))
}

func solveCase(in CaseInput) CaseOutput {
	diffs := make([]int64, in.K)
	diffsTotal := int64(0)
	for i, t := range in.Tabs {
		diffs[int64(i) % in.K] += t
		diffsTotal += t
	}

	maxDiff := abs64(diffsTotal - diffs[0])
	for b := int64(1); b < in.K; b++ {
		maxDiff = max64(abs64(diffsTotal - diffs[b]), maxDiff)
	}
	return CaseOutput{maxDiff}
}

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
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solveSequential(scanner, writer)
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

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
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
