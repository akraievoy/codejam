package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
)

type caseInput struct {
	index int
	v     []int
}

func readCaseInput(scanner *bufio.Scanner, index int) caseInput {
	size := readInt(scanner)
	v := make([]int, size)
	for i := range v {
		v[i] = readInt(scanner)
	}
	in := caseInput{index, v}
	return in
}

type caseOutput struct {
	index             int
	sortingErrorIndex int
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	//	FIXME write the out
	if out.sortingErrorIndex < 0 {
		writef(writer, "Case #%d: OK\n", 1+out.index)
	} else {
		writef(writer, "Case #%d: %d\n", 1+out.index, out.sortingErrorIndex)
	}
}

func troubleSortQuadratic(in *caseInput) {
	for {
		done := true
		for i := 0; i < len(in.v)-2; i++ {
			if in.v[i] > in.v[i+2] {
				done = false
				in.v[i], in.v[i+2] = in.v[i+2], in.v[i]
			}
		}
		if done {
			break
		}
	}
}

type EvenTrouble struct {
	ci *caseInput
}

func (et EvenTrouble) Len() int           { return (len(et.ci.v) + 1) / 2 }
func (et EvenTrouble) Swap(i, j int)      { et.ci.v[i*2], et.ci.v[j*2] = et.ci.v[j*2], et.ci.v[i*2] }
func (et EvenTrouble) Less(i, j int) bool { return et.ci.v[i*2] < et.ci.v[j*2] }

type OddTrouble struct {
	ci *caseInput
}

func (ot OddTrouble) Len() int           { return len(ot.ci.v) / 2 }
func (ot OddTrouble) Swap(i, j int)      { ot.ci.v[i*2+1], ot.ci.v[j*2+1] = ot.ci.v[j*2+1], ot.ci.v[i*2+1] }
func (ot OddTrouble) Less(i, j int) bool { return ot.ci.v[i*2+1] < ot.ci.v[j*2+1] }

func troubleSort(in *caseInput) {
	sort.Sort(EvenTrouble{in})
	sort.Sort(OddTrouble{in})
}

func solveCase(in caseInput) caseOutput {
	troubleSort(&in)

	sortingErrorIndex := -1
	for i := range in.v {
		if i < len(in.v)-1 && in.v[i] > in.v[i+1] {
			sortingErrorIndex = i
			break
		}
	}
	return caseOutput{in.index, sortingErrorIndex}
}

//	everything below is reusable boilerplate
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

