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
type In struct {
	index int
	n     int     // 1 <= n <= 50
	p     int     // 1 <= p <= 50
	r     []int   //	1 <= r[i] <= 10^6, 0 <= i < n
	q     [][]int // 1 <= q[i][j] <= 10^6, 0 <= i < n, 0 <= j < p
}

type Out struct {
	index   int
	maxKits int
}

func kitRange(amt, req int) (int, int) {
	kMin := int(math.Ceil(float64(10*amt)/float64(req*11)))
	kMax := int(math.Floor(float64(10*amt)/float64(9*req)))

	if (kMax < kMin || kMax == 0) {
		return -1, -1
	} else {
		return kMin, kMax
	}
}

func solve(in In) Out {
	for i := 0; i < in.n; i++ {
		sort.Ints(in.q[i])
	}

	js := make([]int, in.n)
	maxJ, maxKits := 0, 0
	for maxJ < in.p {
		first := true
		kMin, kMax := 1, 1
		kMaxMin, kMaxMinI := 0, 0
		for i := 0; kMax > 0 && i < in.n; i++ {
			kiMin, kiMax := kitRange(in.q[i][js[i]], in.r[i])
			if (first) {
				kMin, kMax = kiMin, kiMax
				kMaxMin, kMaxMinI = kiMax, 0
				first = false
		} else {
				kMin = max(kMin, kiMin)
				kMax = min(kMax, kiMax)
				if (kiMax < kMaxMin) {
					kMaxMin = kiMax
					kMaxMinI = i
				}
			}
		}
		if (kMax >= kMin && kMax > 0) {
			maxKits += 1
			for i := 0; i < in.n; i++ {
				js[i]++
				maxJ = max(maxJ, js[i])
			}
		} else {
			js[kMaxMinI]++
			maxJ = max(maxJ, js[kMaxMinI])
		}
	}

	return Out{in.index, maxKits}
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
		n, p := ReadInt(scanner), ReadInt(scanner)
		r := make([]int, n)
		for i := range r {
			r[i] = ReadInt(scanner)
		}
		q := make([][]int, n)
		for i := range q {
			q[i] = make([]int, p)
			for j := range q[i] {
				q[i][j] = ReadInt(scanner)
			}
		}
		in := In{index, n, p, r, q}
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
		Writef(writer, " %d", out.maxKits)
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
