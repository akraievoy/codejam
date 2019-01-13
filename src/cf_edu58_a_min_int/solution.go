package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type LRD struct {
	L, R, D uint64
}

type CaseInput struct {
	lrds []LRD
}

func readCaseInput(scanner *bufio.Scanner) CaseInput {
	size := readInt(scanner)
	nums := make([]LRD, size)
	for i := range nums {
		nums[i] = LRD{
			uint64(readInt64(scanner)),
			uint64(readInt64(scanner)),
			uint64(readInt64(scanner)),
		}
	}
	in := CaseInput{nums}
	return in
}

type caseOutput struct {
	xs []uint64
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	for _, x := range out.xs {
		writef(writer, "%d\n", x)
	}
}

func solveCase(in CaseInput) caseOutput {
	xs := make([]uint64, len(in.lrds))
	for i, lrd := range in.lrds {
		if lrd.L <= lrd.D {
			xs[i] = lrd.R + lrd.D - (lrd.R % lrd.D)
		} else {
			xs[i] = lrd.D
		}
	}
	return caseOutput{xs}
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