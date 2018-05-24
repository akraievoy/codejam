package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
)

type caseInput struct {
	persons uint32
	groups  []group
}

type group struct {
	members    []uint32
	efficiency uint64
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	persons := uint32(readInt64(scanner))
	groups := make([]group, uint32(readInt64(scanner)))
	for i := range groups {
		members := make([]uint32, readInt(scanner))
		efficiency := uint64(readInt64(scanner))
		for m := range members {
			members[m] = uint32(readInt64(scanner))
		}
		groups[i] = group{members, efficiency}
	}
	return caseInput{persons, groups}
}

type caseOutput struct {
	efficiencyTotal uint64
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", out.efficiencyTotal)
}

func solveCase(in caseInput) caseOutput {
	totalEfficiency := uint64(0)
	for _, g := range in.groups {
		totalEfficiency += g.efficiency
	}

	cm := CapacityMatrixNew(in.persons)

	for _, g := range in.groups {
		if len(g.members) == 2 {
			cm.AddLink(g.members[0]-1, g.members[1]-1, uint64(g.efficiency) * 2)
		} else {
			cm.AddLink(g.members[0]-1, g.members[1]-1, uint64(g.efficiency))
			cm.AddLink(g.members[1]-1, g.members[2]-1, uint64(g.efficiency))
			cm.AddLink(g.members[2]-1, g.members[0]-1, uint64(g.efficiency))
		}
	}

	minCut := cm.MinCutStoerWagner()

	return caseOutput{totalEfficiency - uint64(minCut) / 2}
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

type CapacityMatrix [][]uint64

func CapacityMatrixNew(n uint32) CapacityMatrix {
	outerSlice := make([][]uint64, n)
	for i := range outerSlice {
		outerSlice[i] = make([]uint64, n)
	}
	return outerSlice
}

func (cm *CapacityMatrix) Size() int {
	return len(*cm)
}

func (cm *CapacityMatrix) AddLink(i, j uint32, capacity uint64) {
	(*cm)[i][j] += capacity
	(*cm)[j][i] += capacity
}

func (cm *CapacityMatrix) Clone() CapacityMatrix {
	n := cm.Size()

	outerSlice := make([][]uint64, n)

	for i := range outerSlice {
		outerSlice[i] = make([]uint64, n)
		copy(outerSlice[i], (*cm)[i])
	}

	return outerSlice
}

/*
	http://www.e-maxx-ru.1gb.ru/algo/stoer_wagner_mincut#4

	http://www.e-maxx-ru.1gb.ru/bookz/files/stoer_wagner_mincut.pdf
	http://www.e-maxx-ru.1gb.ru/bookz/files/mehlhorn_mincut_stoer_wagner.pdf
 */
func (cm *CapacityMatrix) MinCutStoerWagner() uint64 {
	minCutSet := false
	minCut := uint64(math.MaxUint64)

	g := *cm
	n := len(g)
	w := make([]uint64, n)
	exist := make([]bool, n)
	for i := range exist {
		exist[i] = true
	}
	inA := make([]bool, n)

	for phase := 0; phase < n-1; phase++ {
		for i := range inA {
			inA[i] = false
			w[i] = 0
		}

		prev := -1
		for it := 0; it < n-phase; it++ {

			sel := -1
			for i := 0; i < n; i++ {
				if exist[i] && !inA[i] && (sel == -1 || w[i] > w[sel]) {
					sel = i;
				}
			}

			if it == n-phase-1 {
				if !minCutSet || w[sel] < minCut {
					minCutSet = true
					minCut = w[sel]
				}

				for i := 0; i < n; i++ {
					g[i][prev] += g[sel][i]
					g[prev][i] = g[i][prev]
				}
				exist[sel] = false
			} else {
				inA[sel] = true
				for i := 0; i < n; i++ {
					w[i] += g[sel][i]
				}
				prev = sel
			}

		}
	}

	return minCut;
}
