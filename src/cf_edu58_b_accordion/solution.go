package main

import (
	"bufio"
	"os"
	"fmt"
)

type caseInput struct {
	BayanStr string
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	return caseInput{readString(scanner)}
}

type caseOutput struct {
	BayanLen int32
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", out.BayanLen)
}

func solveCase(in caseInput) caseOutput {
	opened := false
	firstColonSeen := false
	closingColonSeen := false
	closed := false

	pipesTotal := uint32(0)
	pipesCloseable := uint32(0)
	pipesPendingClose := uint32(0)

	for _, c := range []rune(in.BayanStr) {
		if !opened {
			if c == '[' {
				opened = true
			}
			continue
		}

		if opened && !firstColonSeen {
			if c == ':' {
				firstColonSeen = true
			}
			continue
		}

		if opened && firstColonSeen {
			if c == '|' {
				pipesPendingClose += 1
			} else if c == ':' {
				pipesCloseable += pipesPendingClose
				pipesPendingClose = 0
				closingColonSeen = true
			} else if c == ']' {
				if closingColonSeen {
					closed = true
					pipesTotal = pipesCloseable
				}
			}
		}

	}

	if closed {
		return caseOutput{int32(4 + pipesTotal)}
	}
	return caseOutput{int32(-1)}
}

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	writeCaseOutput(writer, solveCase(readCaseInput(scanner)))
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
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solveSequential(scanner, writer)
}

func readString(sc *bufio.Scanner) string {
	if !sc.Scan() {
		panic("failed to scan next token")
	}

	return sc.Text()
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}