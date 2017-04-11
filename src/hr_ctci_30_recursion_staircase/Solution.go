package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	nums []int16
}

type Out struct {
	counts []uint64
}

func steps(num int16) (uint64, uint64, uint64) {
	if (num == 1) {
		return 0, 0, 1
	} else if (num == 2) {
		return 0, 1, 2
	} else if (num == 3) {
		return 1, 2, 4
	} else {
		counts2, counts1, counts := steps(num - 1)
		return counts1, counts, counts2 + counts1 + counts
	}
}

func solve(in In) Out {
	counts := make([]uint64, len(in.nums))
	for i, num := range in.nums {
		_, _, stepsRes := steps(num)
		counts[i] = stepsRes
	}

	return Out{counts}
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

	size := ReadInt16(scanner)
	nums := make([]int16, size)
	for i := range nums {
		nums[i] = ReadInt16(scanner)
	}
	out := solve(In{nums})

	for _, count := range out.counts {
		Writef(writer, "%d\n", count)
	}
}

//	boring IO
func ReadInt16(sc *bufio.Scanner) int16 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 16)
	if err != nil {
		panic(err)
	}
	return int16(res)
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
