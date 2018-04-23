package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type caseInput struct {
	index int
	freqs []int
}

func readCaseInput(scanner *bufio.Scanner, index int) caseInput {
	freqs := make([]int, 26)
	for i := range freqs {
		freqs[i] = readInt(scanner)
	}
	return caseInput{index, freqs}
}

type caseOutput struct {
	index     int
	superiors int
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", out.superiors)
}

func solveCase(in caseInput) caseOutput {
	total := 0
	//					 0 1 2 3 4 5 6 7 8 91011
	//	A: 2 1 1 1 7 -> [0   0   1   2   3   4 4] (4 superiors)
	//	                   4   4   4   4   4
	//	B: 2 0 8 1 1 -> [0   0   2   2   2   2 2] (3 superiors)
	//	                   2   2   2   3   4
	//	C: 0 10 1 1 0-> [1   1   1   1   1   1 1] (2 superiors)
	//	                   1   1   1   2   3
	//	D: 1 4 6     -> [0   1   1   1   1   2]   (4 superiors)
	//                     2   2   2   2   2
	for _, v := range in.freqs {
		total += v
	}

	medianStart := 0 //	A: 5 B: 2 C: 0
	medianC := 0     //	A: 4 B: 2 C: 1
	{
		median := (total + 1) / 2 //	A B C D: 6

		totalTmp := 0
		for c, v := range in.freqs {
			if totalTmp+v >= median {
				medianStart = totalTmp
				medianC = c
				break
			}
			totalTmp += v
		}
	}

	medianEvenFirst := medianStart * 2                          //	A: 10 B: 4 C: 0
	medianOdds := in.freqs[medianC] - (total + 1) / 2 + medianStart //	A: 6 B: 4 C: 4
	if total%2 == 0 {
		medianOdds -= 1 //	A: 5 B: 3 C: 3
	}
	medianOddLast := (medianOdds-1)*2 + 1 //	A: 9 B: 5 C: 5

	if medianOddLast < medianEvenFirst-1 {
		return caseOutput{in.index, (total - 1) / 2}
	}

	overlap := medianOddLast - max(medianEvenFirst, 2) + 3 //	A: 2 B: 4 C: 6

	return caseOutput{in.index, (total - 1 - overlap) / 2}
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}