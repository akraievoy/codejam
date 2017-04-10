package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type In struct {
	index  int32
	digits []int8
}

type Out struct {
	index  int32
	digits []int8
}

func solve(in In) (out Out) {
	tidy := make([]int8, len(in.digits))
	for i := range in.digits {
		tidy[i] = in.digits[i]
	}

	mostSignificantUntidyPos := -1
	for i := 1; i < len(tidy); i++ {
		if (tidy[i] < tidy[i - 1]) {
			mostSignificantUntidyPos = i
			break
		}
	}

	if (mostSignificantUntidyPos < 0) {
		return Out{in.index, tidy}
	}

	rollDownPos := mostSignificantUntidyPos - 1
	for rollDownPos > 0 && tidy[rollDownPos - 1] == tidy[rollDownPos] {
		rollDownPos--
	}

	tidy[rollDownPos] -= 1;

	for i := rollDownPos + 1; i < len(tidy); i++ {
		tidy[i] = 9
	}

	if (tidy[0] == 0) {
		tidy = tidy[1:]
	}

	return Out{in.index, tidy}
}

func solveChannel(ins <-chan In, outs chan <- Out) {
	for in := range ins {
		outs <- solve(in)
	}
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

	caseCount := ReadInt32(scanner)

	cores := runtime.NumCPU()
	var ins = make(chan In, cores)
	var outs = make(chan Out, caseCount)
	for t := 0; t < cores; t++ {
		go solveChannel(ins, outs)
	}

	for index := int32(0); index < caseCount; index++ {
		scanner.Scan()
		numStr := []rune(scanner.Text())
		digits := make([]int8, len(numStr))
		for i, r := range numStr {
			digits[i] = int8(r - '0')
		}
		in := In{index, digits}
		ins <- in
	}
	close(ins)

	outsSlice := make([]Out, caseCount)
	for index := int32(0); index < caseCount; index++ {
		out := <-outs
		outsSlice[out.index] = out
	}
	close(outs)

	for _, out := range outsSlice {
		numStr := make([]rune, len(out.digits))
		for i, d := range out.digits {
			numStr[i] = '0' + int32(d)
		}
		Writef(writer, "Case #%d: %s\n", 1 + out.index, string(numStr))
	}
}

//	boring IO
func ReadInt32(sc *bufio.Scanner) int32 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(res)
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
