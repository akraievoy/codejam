package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type caseInput struct {
	index int
	d     int
	p     string
}

func readCaseInput(scanner *bufio.Scanner, index int) caseInput {
	d := readInt(scanner)
	p := readString(scanner)
	in := caseInput{index, d, p}
	return in
}

type caseOutput struct {
	index int
	hacks int
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	if out.hacks >= 0 {
		writef(writer, "Case #%d: %d\n", 1+out.index, out.hacks)
	} else {
		writef(writer, "Case #%d: IMPOSSIBLE\n", 1+out.index)
	}
}

func solveCase(in caseInput) caseOutput {
	damageForCharges := make([]int, 48)
	for idx := range damageForCharges {
		if idx == 0 {
			damageForCharges[idx] = 1
		} else {
			damageForCharges[idx] = damageForCharges[idx-1] * 2
		}
	}

	firesForCharges := make([]int, 48)
	charges := 0
	damageTotal := 0
	for _, a := range in.p {
		if a == 'C' {
			charges += 1
		} else {
			damageTotal += damageForCharges[charges]
			firesForCharges[charges] += 1
		}
	}

	if damageTotal <= in.d {
		return caseOutput{in.index, 0}
	}

	hacks := 0
	for {
		chargesToHack := -1
		for charges, fires := range firesForCharges {
			if charges > 0 && fires > 0 {
				chargesToHack = charges
			}
		}
		if chargesToHack == -1 {
			return caseOutput{in.index, -1}
		} else {
			damageTotal -= damageForCharges[chargesToHack]
			firesForCharges[chargesToHack] -= 1
			firesForCharges[chargesToHack - 1] += 1
			damageTotal += damageForCharges[chargesToHack - 1]
			hacks += 1
			if damageTotal <= in.d {
				return caseOutput{in.index, hacks}
			}
		}
	}
	return caseOutput{in.index, -1}
}

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseCount := readInt(scanner)
	for index := 0; index < caseCount; index++ {
		writeCaseOutput(writer, solveCase(readCaseInput(scanner, index)))
	}
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