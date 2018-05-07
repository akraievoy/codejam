package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type caseInput struct {
	index   int
	a       int
	scanner *bufio.Scanner
	writer  *bufio.Writer
}

func readCaseInput(scanner *bufio.Scanner, writer *bufio.Writer, index int) caseInput {
	return caseInput{index, readInt(scanner), scanner, writer}
}

func solveCase(in caseInput) bool {
	w := in.a
	h := 3
	p := 2 * (w + h - 2)
	a := w * h

	for wI := 3; wI <= in.a; wI++ {
		for hI := 3; hI <= wI; hI++ {
			pI := 2 * (wI + hI - 2)
			aI := wI * hI
			if aI >= in.a && (aI < a || aI == a && pI < p) {
				w, h, p, a = wI, hI, pI, aI
			}
		}
	}

	d := []int{-1, 0, 1}

	dugCounts := make([][]uint8, h)
	dugFlags := make([][]bool, h)
	for hI := range dugCounts {
		dugCounts[hI] = make([]uint8, w)
		dugFlags[hI] = make([]bool, w)
	}

	for {
		hD, wD := 1, 1
		for hI := 1; hI < len(dugCounts)-1; hI++ {
			for wI := 1; wI < len(dugCounts[hI])-1; wI++ {
				if dugCounts[hI][wI] < dugCounts[hD][wD] {
					hD, wD = hI, wI
				}
			}
		}

		writef(in.writer, "%d %d\n", 1+hD, 1+wD)
		in.writer.Flush()

		hDA := readInt(in.scanner) - 1
		wDA := readInt(in.scanner) - 1

		if hDA < -1 {
			return false
		} else if hDA == -1 {
			return true
		}

		if !dugFlags[hDA][wDA] {
			dugFlags[hDA][wDA] = true
			for _, dH := range d {
				for _, dW := range d {
					wDD, hDD  := wDA + dW, hDA + dH
					if wDD >= 0 && wDD < w && hDD >= 0 && hDD < h {
						dugCounts[hDD][wDD] += 1
					}
				}
			}
		}
	}

	return false
}

//	everything below is reusable boilerplate
func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseCount := readInt(scanner)
	for index := 0; index < caseCount; index++ {
		if !solveCase(readCaseInput(scanner, writer, index)) {
			break
		}
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