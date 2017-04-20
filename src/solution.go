package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	// FIXME test case input structure
	nums  []int16
}

type Out struct {
	// FIXME test case output structure
	sum   int32
}

func solve(in In) (out Out) {
	// FIXME actual solution
	sum := int32(0)
	for _, v := range in.nums {
		sum += int32(v)
	}
	return Out{sum}
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

	//	FIXME read the in
	size := int16(ReadInt(scanner))
	nums := make([]int16, size)
	for i := range nums {
		nums[i] = int16(ReadInt(scanner))
	}
	in := In{nums}

	out := solve(in)

	//	FIXME write the out
	Writef(writer, "%d\n", out.sum)
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
