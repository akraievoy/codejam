package main

import (
	"bufio"
	"runtime"
	"strings"
	"fmt"
	"strconv"
	"os"
)

type In struct {
	index int
	c     []int
	j     []int
}

type Out struct {
	index int
	c     []int
	j     []int
}

type FreePos struct {
	jc    *[]int
	pos   int
	value int
}

func valDiff(c, j []int) (cVal, jVal, absDiff int) {
	cVal, jVal = 0, 0
	weight := 1
	for i := len(c) - 1; i >= 0; i-- {
		cVal += weight * c[i]
		jVal += weight * j[i]
		weight *= 10
	}
	absDiff = abs(cVal - jVal)
	return
}

func unrollAt(pos int, x, xIn []int) [][]int {
	changed := false
	xOut := make([]int, len(x))
	copy(xOut, x)
	if xIn[pos] < 0 && x[pos] > 0 {
		xOut[pos]--
		changed = true
	}
	for i := pos + 1; i < len(x); i++ {
		if (xIn[i] < 0) {
			xOut[i] = 9
			changed = true
		}
	}

	if changed {
		return [][]int{xOut}
	} else {
		return [][]int{}
	}
}

func rollAt(pos int, x, xIn []int) [][]int {
	changed := false
	xOut := make([]int, len(x))
	copy(xOut, x)
	if xIn[pos] < 0 && x[pos] < 9 {
		xOut[pos]++
		changed = true
	}
	for i := pos + 1; i < len(x); i++ {
		if (xIn[i] < 0) {
			xOut[i] = 0
			changed = true
		}
	}

	if changed {
		return [][]int{xOut}
	} else {
		return [][]int{}
	}
}

func solve(in In) Out {
	cn, jn := make([]int, len(in.c)), make([]int, len(in.j))
	conflictIdx := -1

	for i := range cn {
		if in.c[i] == in.j[i] {
			if in.c[i] < 0 {
				cn[i], jn[i] = 0, 0
			} else {
				cn[i], jn[i] = in.c[i], in.j[i]
			}
		} else if min(in.c[i], in.j[i]) == -1 {
			if in.c[i] == -1 {
				jn[i] = in.j[i]
				cn[i] = jn[i]
			} else {
				cn[i] = in.c[i]
				jn[i] = cn[i]
			}
		} else {
			if conflictIdx < 0 {
				conflictIdx = i
			}
			cn[i] = in.c[i]
			jn[i] = in.j[i]
		}
	}

	if conflictIdx < 0 {
		return Out{in.index, cn, jn}
	}

	var cs = [][]int{cn}
	var js = [][]int{jn}
	for rollIdx := 0; rollIdx <= conflictIdx; rollIdx++ {
		cs = append(cs, rollAt(rollIdx, cn, in.c)...)
		cs = append(cs, unrollAt(rollIdx, cn, in.c)...)
		js = append(js, rollAt(rollIdx, jn, in.j)...)
		js = append(js, unrollAt(rollIdx, jn, in.j)...)
	}

	if false {
		fmt.Println(fmt.Sprintf("conflictIdx=%d\nin.c=%v\ncs=%v\nin.j=%v\njs=%v\n", conflictIdx, in.c, cs, in.j, js))
	}

	cVal, jVal, absDiff := valDiff(cn, jn)
	cRes, jRes := cn, jn
	for _, c := range cs {
		for _, j := range js {
			cValI, jValI, absDiffI := valDiff(c, j)
			if absDiffI < absDiff ||
				absDiffI == absDiff && cValI < cVal ||
				absDiffI == absDiff && cValI == cVal && jValI < jVal {
				cVal, jVal, absDiff = cValI, jValI, absDiffI
				cRes, jRes = c, j
			}
		}
	}

	return Out{in.index, cRes, jRes}
}

func solveBrute(in In) Out {
	hasOpenPositions := false
	for i := range in.c {
		if in.c[i] < 0 || in.j[i] < 0 {
			hasOpenPositions = true
		}
	}

	if !hasOpenPositions {
		return Out{in.index, in.c, in.j}
	}

	equal := true
	for i := range in.c {
		if in.c[i] != in.j[i] && max(in.c[i], in.j[i]) >= 0 {
			equal = false
			break
		}
	}

	if equal {
		cOut, jOut := make([]int, len(in.c)), make([]int, len(in.c))
		for i := range in.c {
			if max(in.c[i], in.j[i]) < 0 {
				cOut[i] = 0
				jOut[i] = 0
			} else if in.c[i] < 0 {
				cOut[i] = in.j[i]
				jOut[i] = in.j[i]
			} else if in.j[i] < 0 {
				cOut[i] = in.c[i]
				jOut[i] = in.c[i]
			} else {
				if in.c[i] != in.j[i] {
					panic("this branch should handle cases with equality only, but oh well")
				}
				cOut[i] = in.c[i]
				jOut[i] = in.j[i]
			}
		}
		return Out{in.index, cOut, jOut}
	}

	c, j := make([]int, len(in.c)), make([]int, len(in.c))
	fp := make([]FreePos, 0)

	for i := 0; i < len(in.c); i++ {
		c[i] = in.c[i]
		if in.c[i] == -1 {
			fp = append(fp, FreePos{&c, i, 0})
			c[i] = 0
		} else {
			c[i] = in.c[i]
		}
	}
	for i := 0; i < len(in.c); i++ {
		j[i] = in.j[i]
		if in.j[i] == -1 {
			fp = append(fp, FreePos{&j, i, 0})
			j[i] = 0
		} else {
			j[i] = in.j[i]
		}
	}

	minCVal, minJVal := -1, -1
	var minC, minJ []int = make([]int, len(in.c)), make([]int, len(in.c))

	for {
		cVal, jVal, weight := 0, 0, 1
		for i := len(in.c) - 1; i >= 0; i-- {
			cVal += weight * c[i]
			jVal += weight * j[i]
			weight *= 10
		}

		minDiff := abs(minCVal - minJVal)
		diff := abs(cVal - jVal)
		if minCVal < 0 ||
			minDiff > diff ||
			minDiff == diff && minCVal > cVal ||
			minDiff == diff && minCVal == cVal && minJVal > jVal {
			for i := 0; i < len(in.c); i++ {
				minC[i] = c[i]
				minJ[i] = j[i]
			}
			minCVal = cVal
			minJVal = jVal
		}

		increment := len(fp) - 1
		for increment >= 0 && fp[increment].value == 9 {
			increment--
		}
		if (increment < 0) {
			break
		}

		fp[increment].value += 1
		(*fp[increment].jc)[fp[increment].pos] = fp[increment].value
		increment += 1
		for increment < len(fp) {
			fp[increment].value = 0
			(*fp[increment].jc)[fp[increment].pos] = fp[increment].value
			increment += 1
		}
	}

	return Out{in.index, minC, minJ}
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
		decoding := "0123456789"
		cStr, jStr := []rune(ReadString(scanner)), []rune(ReadString(scanner))
		c, j := make([]int, len(cStr)), make([]int, len(cStr))
		for i := range c {
			c[i] = strings.IndexRune(decoding, cStr[i])
			j[i] = strings.IndexRune(decoding, jStr[i])
		}
		in := In{index, c, j}
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

		Writef(writer, " ")
		for _, d := range out.c {
			Writef(writer, "%d", d)
		}
		Writef(writer, " ")
		for _, d := range out.j {
			Writef(writer, "%d", d)
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
