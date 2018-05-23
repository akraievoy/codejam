package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type caseInput struct {
	index      int
	rows, cols int
	bw         [][]bool
}

func readCaseInput(scanner *bufio.Scanner, index int) caseInput {
	rows := readInt(scanner)
	cols := readInt(scanner)
	bw := make([][]bool, rows)
	for r := range bw {
		bw[r] = make([]bool, cols)
		row := readString(scanner)
		for c, v := range []rune(row) {
			bw[r][c] = v == 'W'
		}
	}
	return caseInput{index, rows, cols, bw}
}

type caseOutput struct {
	index          int
	largestPattern uint32
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "Case #%d: %d\n", 1+out.index, out.largestPattern)
}

func quadrant(qr, qc, r, c int) int {
	if r < qr {
		if c < qc {
			return 0
		}
		return 1
	} else {
		if c < qc {
			return 2
		}
		return 3
	}
}

func qmask(ul, ur, ll, lr bool) int {
	res := 0
	if ul {
		res += 1
	}
	if ur {
		res += 2
	}
	if ll {
		res += 4
	}
	if lr {
		res += 8
	}
	return res
}

func solveCase(in caseInput) caseOutput {
	qMasks := []int{ /*ul*/ 1, /*ur*/ 2, /*ll*/ 4, /*lr*/ 8}

	largestPattern := uint32(1)

	quadrants := make([]bool, 16)
	if in.rows == 1 {
		for c := 0; c+1 < in.cols; c++ {
			ul, ur := in.bw[0][c], in.bw[0][c+1]
			quadrants[qmask(ul, ur, ul, ur)] = true
			quadrants[qmask(ul, ul, ul, ul)] = true
			quadrants[qmask(ur, ur, ur, ur)] = true
		}
	}
	if in.cols == 1 {
		for r := 0; r+1 < in.rows; r++ {
			ul, ll := in.bw[r][0], in.bw[r+1][0]
			quadrants[qmask(ul, ul, ll, ll)] = true
			quadrants[qmask(ul, ul, ul, ul)] = true
			quadrants[qmask(ll, ll, ll, ll)] = true
		}
	}

	for r := 0; r + 1 < in.rows; r++ {
		for c := 0; c + 1 < in.cols; c++ {
			ul, ur, ll, lr := in.bw[r][c], in.bw[r][c+1], in.bw[r+1][c], in.bw[r+1][c+1]
			quadrants[qmask(ul, ur, ll, lr)] = true
			quadrants[qmask(ul, ur, ul, ur)] = true
			quadrants[qmask(ll, lr, ll, lr)] = true
			quadrants[qmask(ul, ul, ll, ll)] = true
			quadrants[qmask(ur, ur, lr, lr)] = true
			quadrants[qmask(ul, ul, ul, ul)] = true
			quadrants[qmask(ur, ur, ur, ur)] = true
			quadrants[qmask(ll, ll, ll, ll)] = true
			quadrants[qmask(lr, lr, lr, lr)] = true
		}
	}

	qVals := make([]bool, 4)
	uf := UnionFindNew(in.rows * in.cols)
	for q := range quadrants {
		if !quadrants[q] {
			continue
		}

		for i, m := range qMasks {
			qVals[i] = (q & m) > 0
		}

		for qr := 0; qr <= in.rows; qr++ {
			for qc := 0; qc <= in.cols; qc++ {

				uf.Reset()
				for r := 0; r < in.rows; r++ {
					for c := 0; c < in.cols; c++ {

						matchCur := in.bw[r][c] == qVals[quadrant(qr, qc, r, c)]
						if !matchCur {
							continue
						}

						matchLeft := false
						if c > 0 {
							matchLeft = in.bw[r][c-1] == qVals[quadrant(qr, qc, r, c-1)]
						}

						matchUpper := false
						if r > 0 {
							matchUpper = in.bw[r-1][c] == qVals[quadrant(qr, qc, r-1, c)]
						}

						ufIdx := uint32(r*in.cols + c)
						if matchLeft {
							uf.Union(ufIdx-1, ufIdx)
						}
						if matchUpper {
							uf.Union(ufIdx-uint32(in.cols), ufIdx)
						}
						if matchLeft || matchUpper {
							newSize := uf.Size(ufIdx)
							if largestPattern < newSize {
								/*
								fmt.Printf(
									"%d <- left=%v upper=%v (%d, %d) with [ul ur ll lr] %v at (%d, %d)\n",
									newSize, matchLeft, matchUpper, r, c, qVals, qr, qc,
								)
								*/
								largestPattern = newSize;
							}
						}
					}
				}
			}
		}
	}
	return caseOutput{in.index, largestPattern}
}

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseCount := readInt(scanner)
	for index := 0; index < caseCount; index++ {
		writeCaseOutput(writer, solveCase(readCaseInput(scanner, index)))
	}
}

func main() {
	var scanner *bufio.Scanner
	if len(os.Getenv("CODEJAM_INPUT")) > 0 {
		reader, err := os.Open(os.Getenv("CODEJAM_INPUT"))
		if err != nil {
			panic(err)
		}
		defer reader.Close()
		scanner = bufio.NewScanner(reader)
	} else if len(os.Args) > 1 {
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

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

//	http://stackoverflow.com/questions/30957644
//	http://www.math.tau.ac.il/~michas/ufind.pdf
//	http://algs4.cs.princeton.edu/15uf/WeightedQuickUnionPathCompressionUF.java.html
type UnionFind struct {
	parent []uint32
	size   []uint32
	count  uint32
}

func UnionFindNew(count int) *UnionFind {
	parent := make([]uint32, count, count)
	size := make([]uint32, count, count)
	lim := uint32(count)
	for i := uint32(0); i < lim; i++ {
		parent[i], size[i] = i, 1
	}
	return &UnionFind{parent, size, lim}
}

func (ufp *UnionFind) Reset() {
	lim := uint32(len(ufp.parent))
	for i := uint32(0); i < lim; i++ {
		ufp.parent[i], ufp.size[i] = i, 1
	}
	ufp.count = lim
}

func (ufp *UnionFind) validate(p uint32) {
	n := uint32(len(ufp.parent))
	if p >= n {
		panic(fmt.Sprintf("index %d is not between 0 and %d", p, n-1))
	}
}

func (ufp *UnionFind) Count() uint32 {
	return ufp.count
}

func (ufp *UnionFind) Find(p uint32) uint32 {
	ufp.validate(p)

	root := p
	for root != ufp.parent[root] {
		root = ufp.parent[root]
	}
	for p != root {
		newp := ufp.parent[p]
		ufp.parent[p] = root
		p = newp
	}

	return root
}

func (ufp *UnionFind) Size(p uint32) uint32 {
	return ufp.size[ufp.Find(p)]
}

func (ufp *UnionFind) Connected(p, q uint32) bool {
	return ufp.Find(p) == ufp.Find(q)
}

func (ufp *UnionFind) Union(p, q uint32) {
	rootP := ufp.Find(p)
	rootQ := ufp.Find(q)

	if rootP == rootQ {
		return
	}

	if ufp.size[rootP] < ufp.size[rootQ] {
		ufp.parent[rootP] = rootQ
		ufp.size[rootQ] += ufp.size[rootP]
	} else {
		ufp.parent[rootQ] = rootP
		ufp.size[rootP] += ufp.size[rootQ]
	}

	ufp.count -= 1
}
