package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	swipes []int
}

type Out struct {
	reward   int
}

func solve(in In) (out Out) {
	reward := 0
	for _,s := range in.swipes {
		reward += min(s * 10, 100)
	}
	return Out{reward}
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

	in := In{[]int{ReadInt(scanner), ReadInt(scanner), ReadInt(scanner)}}

	out := solve(in)

	Writef(writer, "%d\n", out.reward)
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
