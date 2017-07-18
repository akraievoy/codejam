package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	num int
}

type Out struct {
	nextLucky int
}

func lower(num int) int {
  return num % 10 + (num / 10 % 10) + (num / 100 % 10)
}

func upper(num int) int {
  return (num /1000) % 10 + (num / 10000 % 10) + (num / 100000 % 10)
}

func solve(in In) (out Out) {
	nextLucky := in.num + 1

	for lower(nextLucky) != upper(nextLucky) {
		nextLucky++
	}

	return Out{nextLucky}
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

	out := solve(In{ReadInt(scanner)})

	Writef(writer, "%d\n", out.nextLucky)
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
