package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type caseInput struct {
	index   int
	n, k, m int
	a       []int
}

func readCaseInput(scanner *bufio.Scanner, index int) caseInput {
	n, k, m := readInt(scanner), readInt(scanner), readInt(scanner)
	a := make([]int, m)
	for i := range a {
		a[i] = readInt(scanner)
	}
	in := caseInput{index, n, k, m, a}
	return in
}

type caseOutput struct {
	index       int
	sectionForK int
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", out.sectionForK)
}

func solveCase(in caseInput) caseOutput {
	sectionForK := -1
	sectionStart := 0
	for sectionIdx, sectionSize := range in.a {
		sectionStart += sectionSize
		if in.k <= sectionStart {
			sectionForK = sectionIdx + 1
			break
		}
	}
	return caseOutput{in.index, sectionForK}
}

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