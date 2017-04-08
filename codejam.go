package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type In struct {
	index int32
	// FIXME test case input structure
	nums []int16
}

type Out struct {
	index int32
	// FIXME test case output structure
	sum int32
}

func solve(in In) (out Out) {
	// FIXME actual solution
	sum := int32(0)
	for _, v := range in.nums {
		sum += int32(v)
	}
	return Out{in.index, sum}
}

func solveChannel(ins <-chan In, outs chan<- Out) {
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
		//	FIXME read the in
		size := ReadInt16(scanner)
		nums := make([]int16, size)
		for i := range nums {
			nums[i] = ReadInt16(scanner)
		}
		in := In{index, nums}
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
		//	FIXME write the out
		Writef(writer, "Case #%d: %d\n", 1 + out.index, out.sum)
	}
}

//	boring IO
func ReadInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return int64(res)
}

func ReadInt32(sc *bufio.Scanner) int32 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(res)
}

func ReadInt16(sc *bufio.Scanner) int16 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 16)
	if err != nil {
		panic(err)
	}
	return int16(res)
}

func ReadInt8(sc *bufio.Scanner) int8 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 8)
	if err != nil {
		panic(err)
	}
	return int8(res)
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
