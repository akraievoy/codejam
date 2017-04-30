package main

import (
	"bufio"
	"runtime"
	"sort"
	"math"
	"strconv"
	"fmt"
	"os"
)


type PC struct {
	r, h int
}

type PCSort []PC

func (pc_sort PCSort) Len() int {
	return len(pc_sort)
}
func (pc_sort PCSort) Swap(i, j int) {
	pc_sort[i], pc_sort[j] = pc_sort[j], pc_sort[i]
}
func (pc_sort PCSort) Less(i, j int) bool {
	pci, pcj := pc_sort[i], pc_sort[j]
	return pci.r > pcj.r || pci.r == pcj.r && pci.h > pcj.h
}

type HSort []PC

func (h_sort HSort) Len() int {
	return len(h_sort)
}
func (h_sort HSort) Swap(i, j int) {
	h_sort[i], h_sort[j] = h_sort[j], h_sort[i]
}
func (h_sort HSort) Less(i, j int) bool {
	pci, pcj := h_sort[i], h_sort[j]
	return pci.h * pci.r > pcj.h * pcj.r
}

type In struct {
	index int
	n, k  int
	pcs   []PC
}

type Out struct {
	index int
	maxS  float64
}

func solveBrute(in In) Out {
	sort.Sort(PCSort(in.pcs))
	maxS := -1
	sel := make([]bool, in.n)
	for {
		curS := 0
		nSel := 0
		for pci, pc := range in.pcs {
			if !sel[pci] {
				continue
			}
			nSel += 1
			if curS == 0 {
				curS += pc.r * pc.r + 2 * pc.r * pc.h
			} else {
				curS += 2 * pc.r * pc.h
			}
		}
		if nSel == in.k && curS > maxS {
			maxS = curS
		}
		rollPos := len(sel) - 1
		for rollPos >= 0 && sel[rollPos] {
			rollPos--
		}
		if (rollPos < 0) {
			break
		}
		sel[rollPos] = true
		rollPos += 1
		for ; rollPos < len(sel); rollPos++ {
			sel[rollPos] = false
		}
	}
	return Out{in.index, float64(maxS)}
}

func solve(in In) Out {
	sort.Sort(PCSort(in.pcs))
	maxS := -1
	for baseIdx := 0; baseIdx < in.n - in.k + 1; baseIdx++ {
		pcb := in.pcs[baseIdx]
		curS := pcb.r * pcb.r + 2 * pcb.r * pcb.h

		upperChoices := make([]PC, in.n - baseIdx - 1)
		copy(upperChoices, in.pcs[baseIdx + 1:])
		sort.Sort(HSort(upperChoices))

		for upperIdx := 1; upperIdx < in.k; upperIdx++ {
			pcup := upperChoices[upperIdx - 1]
			curS += 2 * pcup.r * pcup.h
		}

		if curS > maxS {
			maxS = curS
		}
	}
	return Out{in.index, float64(maxS)}
}

func solveInput(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseCount := ReadInt(scanner)

	cores := runtime.NumCPU()
	var ins = make(chan In, cores)
	var outs = make(chan Out, caseCount)
	for t := 0; t < cores; t++ {
		go solveChannel(ins, outs)
	}

	outsSlice := make([]Out, caseCount)
	for index := 0; index < caseCount; index++ {
		n := ReadInt(scanner)
		k := ReadInt(scanner)
		pcs := make([]PC, n)
		for i := range pcs {
			pcs[i] = PC{ReadInt(scanner), ReadInt(scanner)}
		}
		in := In{index, n, k, pcs}
		ins <- in
	}
	close(ins)

	for index := 0; index < caseCount; index++ {
		out := <-outs
		outsSlice[out.index] = out
	}
	close(outs)

	for _, out := range outsSlice {
		Writef(writer, "Case #%d:", 1 + out.index)

		Writef(writer, " %.9f", math.Pi * out.maxS)

		Writef(writer, "\n")
	}
}

func solveChannel(ins <-chan In, outs chan <- Out) {
	for in := range ins {
		outs <- solve(in)
	}
}

func min(a, b int) int {
	if (a < b) {
		return a
	}
	return b
}

func max(a, b int) int {
	if (a > b) {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min64(a, b int64) int64 {
	if (a < b) {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if (a > b) {
		return a
	}
	return b
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
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

	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solveInput(scanner, writer)
}

func ReadString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func ReadInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func ReadInt(sc *bufio.Scanner) int {
	return int(ReadInt64(sc))
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
