package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math/big"
)

func readUint64(sc *bufio.Scanner) int64 {
	sc.Scan()
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

func GenerateWeights() [10][16]int64 {
	var res [10][16]int64

	for digit := 0; digit < 10; digit++ {
		v := int64(digit)
		for pow := 0; pow < 16; pow++ {
			res[digit][pow] = v
			v *= 10
		}
	}

	return res
}

var debug = false
var Weights [10][16]int64 = GenerateWeights()

var Megas = [10]bool{false, false, true, true, false, true, false, true, false, false}
var NextMegas = [10]rune{'2', '2', '3', '5', '5', '7', '7', '2', '2', '2'}
var FirstMega = '2'

var TrailingMegas = [10]bool{false, false, false, true, false, false, false, true, false, false}
var TrailingNextMegas = [10]rune{'3', '3', '3', '7', '7', '7', '7', '3', '3', '3'}
var TrailingFirstMega = '3'

func Trailing(pos, len int) bool {
	return pos > 0 && pos == len - 1
}

func IsMegaDigit(digit rune, trailing bool) bool {
	if (trailing) {
		return TrailingMegas[digit - '0']
	}
	return Megas[digit - '0']
}

func NextMegaDigit(digit rune, trailing bool) rune {
	if (trailing) {
		return TrailingNextMegas[digit - '0']
	}
	return NextMegas[digit - '0'];
}

func IsMega(curElem []rune) bool {
	for i, e := range curElem {
		if (!IsMegaDigit(e, i > 0 && i == len(curElem) - 1)) {
			return false
		}
	}
	return true
}

func NextMega(curElem []rune) []rune {
	curLen := len(curElem)
	rollPos := 0
	for rollPos + 1 < curLen && IsMegaDigit(curElem[rollPos], false) {
		rollPos += 1
	}

	carryOver := rollPos
	for carryOver >= 0 &&
		NextMegaDigit(curElem[carryOver], Trailing(carryOver, curLen)) < curElem[carryOver] {
		carryOver--
	}

	if (carryOver < 0) {
		res := make([]rune, curLen + 1)
		for i := range res {
			if (i < curLen) {
				res[i] = FirstMega
			} else {
				res[i] = TrailingFirstMega
			}
		}
		return res
	}

	curElem[carryOver] =
		NextMegaDigit(curElem[carryOver], Trailing(carryOver, curLen))
	for i := carryOver + 1; i < curLen; i++ {
		if (Trailing(i, curLen)) {
			curElem[i] = TrailingFirstMega
		} else {
			curElem[i] = FirstMega
		}
	}

	return curElem
}

func WeightValue(curElem []rune) int64 {
	var res int64
	for i, digit := range curElem {
		res += Weights[digit - '0'][len(curElem) - i - 1];
	}
	return res
}

func main() {
	var scanner *bufio.Scanner;
	if len(os.Args) > 1 {
		//debug = true
		reader, err := os.Open(os.Args[1])
		if (err != nil) {
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

	var first = readUint64(scanner)
	var last = readUint64(scanner)
	if (last == 1) {
		last = 2; //  1 is some weird special case, let it burn somewhere else
	}

	var curElem = []rune(fmt.Sprintf("%d", first))

	if (!IsMega(curElem)) {
		curElem = NextMega(curElem)
	}

	megaPrimes := 0
	curVal := WeightValue(curElem)

	col := 1
	for curVal <= last {
		if (col == 0 && debug) {
			os.Stderr.WriteString("\n")
		}
		if big.NewInt(curVal).ProbablyPrime(8) {
			if debug {
				os.Stderr.WriteString(
					fmt.Sprintf("  [%6d]", curVal))
			}
			megaPrimes += 1
		} else {
			if debug {
				os.Stderr.WriteString(
					fmt.Sprintf("   %6d ", curVal))
			}
		}
		curElem = NextMega(curElem)
		curVal = WeightValue(curElem)
		if debug {
			col = (col + 1) % 16
		}
	}

	writef(writer, "%d\n", megaPrimes)
}
