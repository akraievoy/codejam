package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
)

type LR struct{ L, R, I uint32 }

type CaseInput struct {
	lrs []LR
}

func readCaseInput(scanner *bufio.Scanner) CaseInput {
	size := readInt(scanner)
	nums := make([]LR, size)
	for i := range nums {
		nums[i] = LR{
			uint32(readInt(scanner)),
			uint32(readInt(scanner)),
			uint32(i),
		}
	}
	return CaseInput{nums}
}

type caseOutput struct {
	Grouping []bool
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	if len(out.Grouping) == 0 {
		writef(writer, "-1\n")
	} else {
		for _, g := range out.Grouping {
			if g {
				writef(writer, "2 ")
			} else {
				writef(writer, "1 ")
			}
		}
		writef(writer, "\n")
	}
}

type LRSlice []LR

func (lrs LRSlice) Len() int           { return len(lrs) }
func (lrs LRSlice) Less(i, j int) bool { return lrs[i].L < lrs[j].L || lrs[i].L == lrs[j].L && lrs[i].R < lrs[j].R }
func (lrs LRSlice) Swap(i, j int)      { lrs[i], lrs[j] = lrs[j], lrs[i] }

func solveCase(in CaseInput) caseOutput {
	sort.Sort(LRSlice(in.lrs))

	maxR := in.lrs[0].R
	for i, lr := range in.lrs {
		if i == 0 {
			continue
		}

		if lr.L > maxR {
			grouping := make([]bool, len(in.lrs))
			for _, resLR := range in.lrs[:i] {
				grouping[resLR.I] = true
			}
			return caseOutput{ grouping }
		}

		if maxR < lr.R {
			maxR = lr.R
		}
	}

	return caseOutput{make([]bool, 0)}
}

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseNum := readInt(scanner)
	for i := 0; i < caseNum; i++ {
		writeCaseOutput(writer, solveCase(readCaseInput(scanner)))
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