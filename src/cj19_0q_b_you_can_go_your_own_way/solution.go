package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	res := make([]rune, 0, 100000)

	caseCount := readInt64(scanner)
	for index := int64(0); index < caseCount; index++ {
		res = res[:0]
		_ = readInt64(scanner)
		for _, nMove := range []rune(readString(scanner)) {
			if nMove == 'E' {
				res = append(res, 'S')
			} else {
				res = append(res, 'E')
			}
		}
		writef(writer, "Case #%d: %s\n", 1+index, string(res))
	}
}

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) > 1 {
		reader, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := reader.Close(); err != nil {
				panic(err)
			}
		}()
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	defer func() {
		if err := writer.Flush(); err != nil {
			panic(err)
		}
	}()

	solveSequential(scanner, writer)
}

func readString(sc *bufio.Scanner) string {
	if !sc.Scan() {
		panic("failed to scan next token")
	}

	return sc.Text()
}

func readInt64(sc *bufio.Scanner) int64 {
	if !sc.Scan() {
		panic("failed to scan next token")
	}

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