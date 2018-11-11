package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type caseInput struct {
	index   int
	wordLen int
	words   []string
}

func readCaseInput(scanner *bufio.Scanner, index int) caseInput {
	size := readInt(scanner)
	wordLen := readInt(scanner)
	words := make([]string, size)
	for i := range words {
		words[i] = readString(scanner)
	}
	in := caseInput{index, wordLen, words}
	return in
}

type caseOutput struct {
	index int
	word  string
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(
		writer,
		"Case #%d: %s\n", 1+out.index,
		out.word,
	)
}

func solveCase(in caseInput) caseOutput {
	runesAtPos := make([][]rune, in.wordLen)
	for pos := range runesAtPos {
		runesAtPos[pos] = make([]rune, 0, 26)
	}

	allWords := make(map[string]bool)
	for _, word := range in.words {
		for wordPos, wordRune := range []rune(word) {
			runeFound := false
			for _, runeAt := range runesAtPos[wordPos] {
				if runeAt == wordRune {
					runeFound = true
					break
				}
			}
			if !runeFound {
				runesAtPos[wordPos] = append(runesAtPos[wordPos], wordRune)
			}
		}
		allWords[word] = true
	}

	runeIdx := make([]int, in.wordLen)
	runes := make([]rune, in.wordLen)

	for {
		for pos := range runeIdx {
			runes[pos] = runesAtPos[pos][runeIdx[pos]];
		}

		word := string(runes)
		_, present := allWords[word]
		if !present {
			return caseOutput{in.index, word}
		}

		rollPos := in.wordLen - 1
		for rollPos>=0 && runeIdx[rollPos] + 1 == len(runesAtPos[rollPos]) {
			rollPos -= 1
		}
		if rollPos < 0 {
			break // enumerated all possible words
		}
		runeIdx[rollPos] += 1
		rollPos += 1
		for rollPos < in.wordLen {
			runeIdx[rollPos] = 0
			rollPos += 1
		}

	}

	return caseOutput{in.index, "-"}
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
