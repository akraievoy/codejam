package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	nums  []int
}

type Out struct {
	removed   int
}

func compare(a, b int) int {
	if (a < b) {
		return -1
	}
	return 1
}

func solve(in In) (out Out) {
	removed := 0

	for i := 0; i < len(in.nums) - 2; i++ {
		if compare(in.nums[i], in.nums[i+1]) == compare(in.nums[i+1], in.nums[i+2]) {
			//fmt.Println(fmt.Sprintf("%v", in.nums[i + 1]))
			removed++
		}
	}

	return Out{removed}
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

	size := ReadInt(scanner)
	nums := make([]int, size)
	for i := range nums {
		nums[i] = ReadInt(scanner)
	}
	in := In{nums}

	out := solve(in)

	Writef(writer, "%d\n", out.removed)
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
