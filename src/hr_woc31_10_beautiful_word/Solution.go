package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type In struct {
	w  string
}

type Out struct {
	beautiful   bool
}

func Vowel(r rune) bool {
	return strings.IndexRune("aeiouy", r) >= 0
}

func solve(in In) Out {
	runes := []rune(in.w)
	beautiful := true
	prevVowel := Vowel(runes[0])
	for i := 1; beautiful && i < len(runes); i++ {
		vowel := Vowel(runes[i])
		if (vowel && prevVowel) {
			beautiful = false;
		}
		if (runes[i] == runes[i - 1]) {
			beautiful = false
		}
		prevVowel = vowel
	}

	return Out{beautiful}
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

	scanner.Scan()
	in := In{scanner.Text()}

	out := solve(in)

	res := "No"
	if (out.beautiful) {
		res = "Yes"
	}
	Writef(writer, "%s\n", res)
}

//	boring IO
func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
