package main

import (
	"bufio"
	"runtime"
	"strconv"
	"fmt"
	"os"
)


type A struct {
	s, e int
}

type In struct {
	index  int
	ac, aj []A
}

type Out struct {
	index        int
	minExchanges int
}

func within(outer, inner A) bool {
	return inner.s >= outer.s && inner.e <= outer.e
}

func intersects(a, b A) bool {
	return max(a.s, b.s) < min(a.e, b.e)
}

func solve(in In) Out {
	for e1 := 0; e1 < 720; e1++ {
		e1a := A{e1, e1+720}

		e1IsCamerons := true
		for _, ac := range in.ac {
			if !within(e1a, ac) {
				e1IsCamerons = false
			}
		}
		for _, aj := range in.aj {
			if intersects(e1a, aj) {
				e1IsCamerons = false
			}
		}

		e1IsJamies := true
		for _, aj := range in.aj {
			if !within(e1a, aj) {
				e1IsJamies = false
			}
		}
		for _, ac := range in.ac {
			if intersects(e1a, ac) {
				e1IsJamies = false
			}
		}

		if (e1IsJamies || e1IsCamerons) {
			return Out{in.index, 2}
		}
	}

	return Out{in.index, 4}
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
		acn := ReadInt(scanner)
		ac := make([]A, acn)
		ajn := ReadInt(scanner)
		aj := make([]A, ajn)
		for i := range ac {
			ac[i] = A{ReadInt(scanner), ReadInt(scanner)}
		}
		for i := range aj {
			aj[i] = A{ReadInt(scanner), ReadInt(scanner)}
		}
		in := In{index, ac, aj}
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

		Writef(writer, " %d", out.minExchanges)

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
