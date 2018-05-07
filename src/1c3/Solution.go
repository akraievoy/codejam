package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
)

type caseInput struct {
	index   int
	weights []int
}

func readCaseInput(scanner *bufio.Scanner, index int) caseInput {
	size := readInt(scanner)
	nums := make([]int, size)
	for i := range nums {
		nums[i] = readInt(scanner)
	}
	return caseInput{index, nums}
}

type caseOutput struct {
	index              int
	longestSubsequence int
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "Case #%d: %d\n", 1+out.index, out.longestSubsequence)
}

func solveCase(in caseInput) caseOutput {
	//	sub-sequences from zero to I, inclusive
	//	sub-sequence length is the inner index + 1
	lightest := make([][]int, len(in.weights))

	minWeight := math.MaxInt64
	longestSubsequence := 0
	for i, weight := range in.weights {
		minWeight = min(minWeight, weight)
		lightest[i] = append(make([]int, 0, 140), minWeight)
		longestSubsequence = max(longestSubsequence, 1)

		for prev := max(i-1, 0); prev < i; prev++ {
			for prevLenMinusOne, lightestPrev := range lightest[prev] {
				for len(lightest[i]) < prevLenMinusOne + 1 {
					lightest[i] = append(lightest[i], math.MaxInt64)
				}
				lightest[i][prevLenMinusOne] = min(lightest[i][prevLenMinusOne], lightest[prev][prevLenMinusOne])

				lengthWithPrev := prevLenMinusOne + 1 + 1
				lightestWithPrev := lightestPrev + weight
				if lightestPrev > 6*weight {
					continue
				}
				if len(lightest[i]) < lengthWithPrev || lightest[i][lengthWithPrev - 1] > lightestPrev {
					for len(lightest[i]) < lengthWithPrev {
						lightest[i] = append(lightest[i], math.MaxInt64)
					}
					lightest[i][lengthWithPrev - 1] = lightestWithPrev
				}
				longestSubsequence = max(longestSubsequence, lengthWithPrev)
			}
		}
	}

	return caseOutput{in.index, longestSubsequence}
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
