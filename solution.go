package main

import (
	"bufio"
	"runtime"
)

type In struct {
	index int
	// FIXME test case input structure
	nums []int16
}

type Out struct {
	index int
	// FIXME test case output structure
	sum int32
}

func solve(in In) Out {
	// FIXME actual solution
	sum := int32(0)
	for _, v := range in.nums {
		sum += int32(v)
	}
	return Out{in.index, sum}
}

func solveInput(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseCount := int(ReadInt32(scanner))

	cores := runtime.NumCPU()
	var ins = make(chan In, cores)
	var outs = make(chan Out, caseCount)
	for t := 0; t < cores; t++ {
		go solveChannel(ins, outs)
	}

	outsSlice := make([]Out, caseCount)
	for index := 0; index < caseCount; index++ {
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

	for index := 0; index < caseCount; index++ {
		out := <-outs
		outsSlice[out.index] = out
	}
	close(outs)

	for _, out := range outsSlice {
		Writef(writer, "Case #%d:", 1 + out.index)

		//	FIXME write the out
		Writef(writer, " %d", out.sum)

		Writef(writer, "\n")
	}
}

func solveChannel(ins <-chan In, outs chan<- Out) {
	for in := range ins {
		outs <- solve(in)
	}
}
