package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
)

type caseInput struct {
	a, p  uint32
	seats string
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	_ = uint32(readInt64(scanner))

	a := uint32(readInt64(scanner))
	p := uint32(readInt64(scanner))
	seats := readString(scanner)

	return caseInput{a, p, seats}
}

type caseOutput struct {
	seated uint32
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", out.seated)
}

type Uint32Slice []uint32

func (p Uint32Slice) Len() int           { return len(p) }
func (p Uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func solveCase(in caseInput) caseOutput {
	freeZoneSizes := make([]uint32, 0, 0)

	firstFreeSeat := -1
	for i, c := range []rune(in.seats) {
		if c == '*' {
			if firstFreeSeat >= 0 {
				freeZoneSizes = append(freeZoneSizes, uint32(i-firstFreeSeat))
				firstFreeSeat = -1
			}
		} else if c == '.' {
			if firstFreeSeat == -1 {
				firstFreeSeat = i
			}
		}
	}
	if firstFreeSeat != -1 {
		freeZoneSizes = append(freeZoneSizes, uint32(len(in.seats)-firstFreeSeat))
	}

	sort.Sort(Uint32Slice(freeZoneSizes))

	seated := uint32(0)
	a, p := in.a, in.p
	for len(freeZoneSizes) > 0 && a+p > 0 {
		freeZone := freeZoneSizes[0]
		freeZoneSizes = freeZoneSizes[1:]

		bigger := (freeZone + 1) / 2
		smaller := freeZone / 2

		seatedA := uint32(0)
		seatedP := uint32(0)
		if a > p {
			seatedA = min_u32(a, bigger)
			seatedP = min_u32(p, smaller)
		} else {
			seatedA = min_u32(a, smaller)
			seatedP = min_u32(p, bigger)
		}

		seated += seatedA
		seated += seatedP
		a -= seatedA
		p -= seatedP
	}

	return caseOutput{seated}
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

func readString(sc *bufio.Scanner) string {
	if !sc.Scan() {
		panic("failed to scan next token")
	}

	return sc.Text()
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

func min_u32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}