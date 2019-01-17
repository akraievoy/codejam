package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	n := uint64(readInt64(scanner))
	m := readInt64(scanner)

	minCount := uint64(0)
	aToCount := make([]uint64, n)

	countToDistinctA := make([]uint64, m+1)
	countToDistinctA[0] = n

	for i := int64(0); i < m; i++ {
		a := readInt64(scanner) - 1

		aCount := aToCount[a]

		countToDistinctA[aCount] -= 1
		if countToDistinctA[aCount] == 0 && minCount == aCount {
			minCount += 1
			writef(writer, "1")
		} else {
			writef(writer, "0")
		}
		aToCount[a] += 1

		countToDistinctA[aCount+1] += 1
	}

	writef(writer, "\n")
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