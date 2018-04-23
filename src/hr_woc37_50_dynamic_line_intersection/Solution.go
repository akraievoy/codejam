package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type query interface {
	isAdd() bool
	isRemove() bool
	isAsk() bool
	getK() uint32
	getB() uint32
	getQ() uint32
}

type addQuery struct {
	k, b uint32
}

func (aq addQuery) isAdd() bool    { return true }
func (aq addQuery) isRemove() bool { return false }
func (aq addQuery) isAsk() bool    { return false }
func (aq addQuery) getK() uint32   { return aq.k }
func (aq addQuery) getB() uint32   { return aq.b }
func (aq addQuery) getQ() uint32   { return 0 }

type removeQuery struct {
	k, b uint32
}

func (rq removeQuery) isAdd() bool    { return false }
func (rq removeQuery) isRemove() bool { return true }
func (rq removeQuery) isAsk() bool    { return false }
func (rq removeQuery) getK() uint32   { return rq.k }
func (rq removeQuery) getB() uint32   { return rq.b }
func (rq removeQuery) getQ() uint32   { return 0 }

type askQuery struct {
	q uint32
}

func (aq askQuery) isAdd() bool    { return false }
func (aq askQuery) isRemove() bool { return false }
func (aq askQuery) isAsk() bool    { return true }
func (aq askQuery) getK() uint32   { return 0 }
func (aq askQuery) getB() uint32   { return 0 }
func (aq askQuery) getQ() uint32   { return aq.q }

type caseInput struct {
	queries []query
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	queryN := readInt(scanner)
	queries := make([]query, queryN)
	for i := range queries {
		q := readString(scanner)
		switch q {
		case "+":
			queries[i] =
				addQuery{
					uint32(readInt(scanner)), uint32(readInt(scanner)),
				}
		case "-":
			queries[i] =
				removeQuery{
					uint32(readInt(scanner)), uint32(readInt(scanner)),
				}
		case "?":
			queries[i] =
				askQuery{
					uint32(readInt(scanner)),
				}
		}
	}
	return caseInput{queries}
}

type caseOutput struct {
	intersections []uint32
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	for _, v := range out.intersections {
		writef(writer, "%d\n", v)
	}
}

func solveCase(in caseInput) caseOutput {
	limit := uint32(100000)
	linesAt := make([]uint32, limit+1)
	modLimit := uint32(500)

	cachedModKsValid := true
	cachedModKs := make([]uint32, 0, modLimit)
	bTotalsForModK := make([]uint32, modLimit+1, modLimit+1)
	bCountsForModK := make([][]uint32, modLimit+1, modLimit+1)

	intersections := make([]uint32, 0, len(in.queries))
	for _, que := range in.queries {
		if que.isAdd() {
			k, b := que.getK(), que.getB()
			if k <= modLimit {
				cachedModKsValid = cachedModKsValid && bTotalsForModK[k] > 0
				bTotalsForModK[k] += 1
				if bCountsForModK[k] == nil {
					bCountsForModK[k] = make([]uint32, k, k)
				}
				bCountsForModK[k][b%k] += 1
			} else {
				for at := b % k; at <= limit; at += k {
					linesAt[at] += 1
				}
			}
		} else if que.isRemove() {
			k, b := que.getK(), que.getB()
			if k <= modLimit {
				cachedModKsValid = cachedModKsValid && bTotalsForModK[k] > 1
				bTotalsForModK[k] -= 1
				bCountsForModK[k][b%k] -= 1
			} else {
				for at := b % k; at <= limit; at += k {
					linesAt[at] -= 1
				}
			}
		} else if que.isAsk() {
			q := que.getQ()
			intersectionsRes := linesAt[q]
			if !cachedModKsValid {
				cachedModKsValid = true
				cachedModKs = cachedModKs[:0]
				for k, bT := range bTotalsForModK {
					if bT > 0 {
						cachedModKs = append(cachedModKs, uint32(k))
					}
				}
			}
			for _, k := range cachedModKs {
				intersectionsRes += bCountsForModK[k][q%k]
			}
			intersections = append(intersections, intersectionsRes)
		}
	}

	return caseOutput{intersections}
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

func readString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
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
