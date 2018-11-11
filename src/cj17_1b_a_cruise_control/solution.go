package main

import (
	"bufio"
	"runtime"
	"strconv"
	"fmt"
	"os"
)

type In struct {
	index int
	d int //	10^9
	n int //	10^3
	k []int //	10^9
	s []int // 1^4
}

type Out struct {
	index int
	maxSpeed float64
}

func solve(in In) Out {
	maxT := float64(0)
	maxTi := 0
	for i := 0; i < in.n; i++ {
		t := float64(in.d - in.k[i]) / float64(in.s[i])
		if i == 0 {
			maxT = t
			maxTi = 0
		} else {
			if (t > maxT) {
				maxT = t
				maxTi = i
			}
		}
	}
	maxSpeed := float64(in.d * in.s[maxTi]) / float64(in.d - in.k[maxTi])

	return Out{in.index, maxSpeed}
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
		d := ReadInt(scanner)
		n := ReadInt(scanner)
		k := make([]int, n)
		s := make([]int, n)
		for i := range k {
			k[i] = ReadInt(scanner)
			s[i] = ReadInt(scanner)
		}
		in := In{index, d, n, k, s}
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

		Writef(writer, " %.6f", out.maxSpeed)

		Writef(writer, "\n")
	}
}

func solveChannel(ins <-chan In, outs chan<- Out) {
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
