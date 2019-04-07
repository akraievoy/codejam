package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseCount := readInt64(scanner)
	testStoreQuery := make([]rune, 0, 1024)
	testStoreResult := make([]int64, 0, 1024)
	for index := int64(0); index < caseCount; index++ {
		n, b, f := readInt64(scanner), readInt64(scanner), readInt64(scanner)

		testStoreResult = testStoreResult[:0]
		for i := int64(0); i < n - b; i++ {
			testStoreResult = append(testStoreResult, 0)
		}

		for pow := int64(0); pow < f; pow++ {
			powVal := int64(1) << uint(pow)
			testStoreQuery := testStoreQuery[:0]
			for i := int64(0); i < n; i++ {
				appended := '0'
				if i & powVal > 0 {
					appended = '1'
				}
				testStoreQuery = append(testStoreQuery, appended)
			}
			writef(writer, "%s\n", string(testStoreQuery))
			if err := writer.Flush(); err != nil {
				panic(err)
			}
			reply := readString(scanner)
			if reply == "-1" {
				panic(reply)
			}
			for i, r := range []rune(reply) {
				if r == '1' {
					testStoreResult[i] |= powVal
				}
			}
		}

		mask := (int64(1) << uint(f)) - 1
		resultPos := 0
		subsequent := false

		for brokerId := int64(0); brokerId < n; brokerId++ {
			if resultPos < len(testStoreResult) && testStoreResult[resultPos] == brokerId & mask {
				resultPos += 1
			} else {
				if subsequent {
					writef(writer, " ")
				}
				writef(writer, "%d", brokerId)
				subsequent = true
			}
		}
		writef(writer, "\n")

		if err := writer.Flush(); err != nil {
			panic(err)
		}

		testCaseReply := readString(scanner)
		if testCaseReply == "1" {
			// cool we're still on track
		} else {
			panic(testCaseReply)
		}

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
