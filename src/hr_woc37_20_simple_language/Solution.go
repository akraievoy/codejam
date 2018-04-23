package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type instruction struct {
	op string
	amount int
}

type caseInput struct {
	instructions []instruction
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	size := readInt(scanner)
	instructions := make([]instruction, size)
	for i := range instructions {
		instructions[i] = instruction{readString(scanner), readInt(scanner)}
	}
	return caseInput{instructions}
}

type caseOutput struct {
	maxResult int
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", out.maxResult)
}

func solveCase(in caseInput) caseOutput {
	maxResult := 0
	for _, instr := range in.instructions {
		if instr.op == "add" {
			if instr.amount > 0 {
				maxResult += instr.amount
			}
		} else if instr.op == "set" {
			if instr.amount > maxResult {
				maxResult = instr.amount
			}
		}
	}
	return caseOutput{maxResult}
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

	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solveSequential(scanner, writer)
}

func readString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func readInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func readInt(sc *bufio.Scanner) int {
	return int(readInt64(sc))
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
