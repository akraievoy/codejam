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
	s []rune
}

type Out struct {
	index int
	lw []rune
}

func solve(in In) Out {
	lw := make([]rune, 0, len(in.s))
	for _, r := range in.s {
		if len(lw) == 0 || r < lw[0] {
			lw = append(lw, r)
		} else {
			lw = append([]rune{r}, lw...)
		}
	}
	return Out{in.index, lw}
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
		in := In{index, []rune(ReadString(scanner))}
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

		Writef(writer, " %s", string(out.lw))

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
