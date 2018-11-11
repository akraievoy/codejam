package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
)

type caseInput struct {
	ratings []int
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	size := readInt(scanner)
	ratings := make([]int, size)
	for i := range ratings {
		ratings[i] = readInt(scanner)
	}
	return caseInput{ratings}
}

type caseOutput struct {
	averageOfTop string
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%s\n", out.averageOfTop)
}

//	https://www.cockroachlabs.com/blog/rounding-implementations-in-go/
func Round(x float64) float64 {
	const (
		mask     = 0x7FF
		shift    = 64 - 11 - 1
		bias     = 1023

		signMask = 1 << 63
		fracMask = (1 << shift) - 1
		halfMask = 1 << (shift - 1)
		one      = bias << shift
	)

	bits := math.Float64bits(x)
	e := uint(bits>>shift) & mask
	switch {
	case e < bias:
		// Round abs(x)<1 including denormals.
		bits &= signMask // +-0
		if e == bias-1 {
			bits |= one // +-1
		}
	case e < bias+shift:
		// Round any abs(x)>=1 containing a fractional component [0,1).
		e -= bias
		bits += halfMask >> e
		bits &^= fracMask >> e
	}
	return math.Float64frombits(bits)
}

func solveCase(in caseInput) caseOutput {
	topSum := 0
	topCount := 0
	for _, v := range in.ratings {
		if v >= 90 {
			topSum += v
			topCount += 1
		}
	}

	averageOfTop := Round(100 * float64(topSum) / float64(topCount)) / 100.0
	averageOfTopPrinted :=
		strconv.FormatFloat(averageOfTop, 'f', 2, 64)

	return caseOutput{averageOfTopPrinted}
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