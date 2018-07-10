package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type caseInput struct {
	k    int32
	nums []int32
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	size := readInt(scanner)
	k := int32(readInt(scanner))
	nums := make([]int32, size)
	for i := range nums {
		nums[i] = int32(readInt(scanner))
	}
	in := caseInput{k, nums}
	return in
}

type caseOutput struct {
	minutes uint32
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", out.minutes)
}

func solveCase(in caseInput) caseOutput {
	offsetStats := make(map[int64]uint32)
	mostPopularOffsetStat := uint32(0)

	for i, num := range in.nums {
		offset := int64(num) - int64(i) * int64(in.k)
		prevStat, pres := offsetStats[offset]
		var stat uint32
		if pres {
			stat = prevStat + 1
		} else {
			stat = 1
		}

		offsetStats[offset] = stat
		mostPopularOffsetStat = max(mostPopularOffsetStat, stat)
	}

	return caseOutput{uint32(len(in.nums)) - mostPopularOffsetStat}
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

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}