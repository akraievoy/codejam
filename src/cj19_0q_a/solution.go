package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	aDigs, bDigs := make([]rune, 0, 100), make([]rune, 0, 100)

	caseCount := readInt64(scanner)
	for index := int64(0); index < caseCount; index++ {
		aDigs, bDigs = aDigs[:0], bDigs[:0]
		for _, nDig := range []rune(readString(scanner)) {
			if nDig == '4' {
				aDigs = append(aDigs, '2')
				bDigs = append(bDigs, '2')
			} else {
				if len(bDigs) > 0 {
					bDigs = append(bDigs, '0')
				}
				aDigs = append(aDigs, nDig)
			}
		}
		writef(writer, "Case #%d: %s %s\n", 1+index, string(aDigs), string(bDigs))
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