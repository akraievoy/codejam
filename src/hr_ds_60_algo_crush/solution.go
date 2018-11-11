package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func readUint64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}

type Op struct {
	index int32
	delta int64
}

type OpSort []Op

func (ops OpSort) Len() int {
	return len(ops)
}
func (ops OpSort) Swap(i, j int) {
	ops[i], ops[j] = ops[j], ops[i]
}
func (ops OpSort) Less(i, j int) bool {
	return ops[i].index < ops[j].index
}

func main() {
	var scanner *bufio.Scanner;
	if len(os.Args) > 1 {
		reader, err := os.Open(os.Args[1])
		if (err != nil) {
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

	var _, M = readUint64(scanner), int32(readUint64(scanner))
	var ops []Op = make([]Op, 2*M)

	for opI := int32(0); opI < M; opI++ {
		a, b := int32(readUint64(scanner)), int32(readUint64(scanner))
		k := int64(readUint64(scanner))
		ops[opI * 2] = Op{a, k}
		ops[opI * 2 + 1] = Op{b + 1, -k}
	}

	sort.Sort(OpSort(ops))

	pos := int32(-1)
	sum, maxSum := int64(0), int64(0)
	for opI := int32(0); opI < 2*M; opI++ {
		op := ops[opI]
		if (op.index > pos) {
			if pos < 0 || sum > maxSum {
				maxSum = sum
			}
			pos = op.index
		}
		sum += op.delta
	}
	if (sum > maxSum) {
		maxSum = sum
	}
	writef(writer, "%d", maxSum)
}
