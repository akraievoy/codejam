package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	x int
}

type Out struct {
	charAtX int
}

func solve(in In) (out Out) {
	x := in.x
	charAtX := 0
	for x > 0 {
		charAtX = 1 - charAtX
		x = x & (x - 1)
	}
	return Out{charAtX}
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

	q := ReadInt(scanner)
	for i := 0; i < q; i++ {
		out := solve(In{ReadInt(scanner)})
		Writef(writer, "%d\n", out.charAtX)
	}
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
