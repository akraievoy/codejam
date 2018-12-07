package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type caseInput struct {
	originalSize uint32
	valueToLastIndex map[uint64]uint32
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	valueToLastIndex := make(map[uint64]uint32)
	size := uint32(readInt(scanner))

	for i := uint32(0); i < size; i++ {
		v := uint64(readInt64(scanner))

		for {
			_, prs := valueToLastIndex[v]

			if !prs {
				valueToLastIndex[v] = i
				break
			} else {
				delete(valueToLastIndex, v)
				v *= 2
			}
		}
	}

	return caseInput{size, valueToLastIndex}
}

type caseOutput struct {
	nums []uint64
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", len(out.nums))
	for idx, num := range out.nums {
		if idx+1 < len(out.nums) {
			writef(writer, "%d ", num)
		} else {
			writef(writer, "%d\n", num)
		}
	}
}

func solveCase(in caseInput) caseOutput {
	temp := make([]uint64, in.originalSize)
	for value, lastIndex := range in.valueToLastIndex {
		temp[lastIndex] = value
	}

	res := make([]uint64, 0, len(in.valueToLastIndex))
	for _, v := range temp {
		if v > 0 {
			res = append(res, v)
		}
	}

	return caseOutput{res}
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
