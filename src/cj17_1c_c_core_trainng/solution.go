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
	n, k  int
	u     float64
	p     []float64
}

type Out struct {
	index   int
	maxProb float64
}

func solve(in In) Out {
	uleft := in.u
	p := make([]float64, len(in.p))
	copy(p, in.p)

	for uleft > 0 {
		minP := p[0]
		minPCount := 1
		for i := 1; i < len(p); i++ {
			if (minP == p[i]) {
				minPCount += 1
			} else if minP > p[i] {
				minP = p[i]
				minPCount = 1
			}
		}
		nextMinI := -1
		for i, pi := range p {
			if nextMinI >= 0 && pi > minP && pi < p[nextMinI] {
				nextMinI = i
			}
			if nextMinI < 0 && pi > minP {
				nextMinI = i
			}
		}
		//fmt.Println(fmt.Sprintf("p=%v", p))
		//fmt.Println(fmt.Sprintf("minP=%v minPCount=%v, nextMinI=%v", minP, minPCount, nextMinI))
		if nextMinI >= 0 && (p[nextMinI] - minP) * float64(minPCount) <= uleft {
			uleft -= (p[nextMinI] - minP) * float64(minPCount)
			for i, pi := range p {
				if pi == minP {
					p[i] = p[nextMinI]
				}
			}
		} else {
			increment := uleft / float64(minPCount)
			uleft = 0
			for i, pi := range p {
				if pi == minP {
					p[i] = pi + increment
				}
			}
		}
	}

	maxProb := 1.0
	for _, p := range p {
		maxProb *= p
	}
	return Out{in.index, maxProb}
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
		u := ReadFloat64(scanner)
		p := make([]float64, n)
		for i := range p {
			p[i] = ReadFloat64(scanner)
		}
		in := In{index, n, k, u, p}
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

		Writef(writer, " %.9f", out.maxProb)

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

func ReadFloat64(sc *bufio.Scanner) float64 {
	sc.Scan()
	res, err := strconv.ParseFloat(sc.Text(), 64)
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
