package main

import (
	"bufio"
	"runtime"
	"strconv"
	"fmt"
	"os"
)

/*
RYB
000 -
001 B
010 Y
011 G
100 R
101 V
110 O
111 -
 */
var numToColor = []rune("-BYGRVO-")

type In struct {
	index            int
	n                int //	3..1000
	r, o, y, g, b, v int //	1000
}

type Out struct {
	index int
	manes []int
}

func solve(in In) Out {
	counts := []int{0, 2 * in.b, 2 * in.y, 2 * in.g, 2 * in.r, 2 * in.v, 2 * in.o}
	routes := []int{0, 3, 3, 1, 3, 1, 1}
	manes := make([]int, in.n)
	pos := 0

	start := 1
	if counts[2] > counts[start] {
		start = 2
	}
	if (counts[4] > counts[start]) {
		start = 4
	}

	if counts[start] == 0 {
		//fmt.Println(fmt.Sprintf("manes=%v", manes))
		return Out{in.index, nil}
	}

	counts[start] -= 1
	manes[0] = start;
	pos++;

	for pos < in.n {
		next := 0
		for dest := 1; dest < 7; dest++ {
			if (manes[pos - 1] & dest) == 0 &&
				counts[dest] > 1 &&
				(
					next == 0 ||
						counts[next] < counts[dest] ||
						routes[next] > routes[dest]) {
				next = dest
			}
		}
		if (next == 0) {
			//fmt.Println(fmt.Sprintf("manes=%v", manes))
			return Out{in.index, nil}
		} else {
			manes[pos] = next
			counts[next] -= 2
			pos++
		}
	}
	if pos < in.n || (manes[0] & manes[in.n-1]) > 0 {
		//fmt.Println(fmt.Sprintf("manes=%v", manes))
		return Out{in.index, nil}
	}
	return Out{in.index, manes}
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
		n, r, o, y, g, b, v := ReadInt(scanner), ReadInt(scanner), ReadInt(scanner), ReadInt(scanner), ReadInt(scanner), ReadInt(scanner), ReadInt(scanner)
		in := In{index, n, r, o, y, g, b, v}
		ins <- in
	}
	close(ins)

	for index := 0; index < caseCount; index++ {
		out := <-outs
		outsSlice[out.index] = out
	}
	close(outs)

	for _, out := range outsSlice {
		Writef(writer, "Case #%d: ", 1 + out.index)

		if out.manes != nil {
			for _, num := range out.manes {
				Writef(writer, "%c", numToColor[num])
			}
		} else {
			Writef(writer, "IMPOSSIBLE")
		}

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
