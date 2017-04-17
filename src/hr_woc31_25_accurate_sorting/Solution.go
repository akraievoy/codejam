package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	nums []int32
}

type Out struct {
	sortable bool
}

func solve(in In) (out Out) {
	furtherSortable := true
	fullySorted := false

	for furtherSortable {
		furtherSortable = false
		fullySorted = true
		for i := 1; i < len(in.nums); i++ {
			cur := in.nums[i]
			prev := in.nums[i - 1]

			if (cur < prev) {
				fullySorted = false;
			}
			if (cur == prev - 1) {
				furtherSortable = true
				in.nums[i - 1], in.nums[i] = cur, prev
			}
		}
	}

	return Out{fullySorted}
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

	tests := int(ReadInt32(scanner))

	for i := 0; i < tests; i++ {
		size := ReadInt32(scanner)
		nums := make([]int32, size)
		for i := range nums {
			nums[i] = ReadInt32(scanner)
		}

		out := solve(In{nums})

		res := "No"
		if (out.sortable) {
			res = "Yes"
		}
		Writef(writer, "%s\n", res)
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
