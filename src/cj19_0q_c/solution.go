package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
)

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	alphabet := make([]*big.Int, 0, 26)
	products, letters, cleartext := make([]*big.Int, 0, 100), make([]*big.Int, 0, 100), make([]rune, 0, 100)

	caseCount := readInt64(scanner)
	for index := int64(0); index < caseCount; index++ {
		alphabet, letters, products, cleartext := alphabet[:0], letters[:0], products[:0], cleartext[:0]
		readString(scanner) // prime limit, meh, not needed
		productCount := readInt64(scanner)
		for prodIndex := int64(0); prodIndex < productCount; prodIndex++ {
			parsed, success := big.NewInt(0).SetString(readString(scanner), 10)
			if !success {
				panic(fmt.Sprintf("failed to parse product %d in test %d", prodIndex+1, index+1))
			}
			products = append(products, parsed)
			if prodIndex == 0 {
				letters = append(letters, nil)
			} else {
				if products[prodIndex-1].Cmp(products[prodIndex]) != 0 {
					letters = append(
						letters,
						big.NewInt(0).GCD(nil, nil, products[prodIndex-1], products[prodIndex]),
					)
				} else { //	this is possible for cases like ZZZ OR ZYZ
					letters = append(letters, nil)
				}
			}
		}

		//	if we know at least one letter anywhere, propagate it both left and right
		for prodIndex := int64(0); prodIndex+1 < productCount; prodIndex++ {
			if letters[prodIndex] != nil && letters[prodIndex+1] == nil {
				letters[prodIndex+1] = big.NewInt(0).Div(products[prodIndex], letters[prodIndex])
			}
		}
		for prodIndex := productCount - 2; prodIndex >= 0; prodIndex-- {
			if letters[prodIndex+1] != nil && letters[prodIndex] == nil {
				letters[prodIndex] = big.NewInt(0).Div(products[prodIndex], letters[prodIndex+1])
			}
		}
		letters = append(letters, big.NewInt(0).Div(products[productCount-1], letters[productCount-1]))

		for _, letter := range letters {
			searchIdx := sort.Search(
				len(alphabet),
				func(i int) bool { return alphabet[i].Cmp(letter) >= 0 },
			)
			if searchIdx < len(alphabet) {
				if alphabet[searchIdx].Cmp(letter) > 0 {
					alphabet = append(alphabet[:searchIdx], append([]*big.Int{letter}, alphabet[searchIdx:]...)...)
				}
			} else {
				alphabet = append(alphabet, letter)
			}
		}

		for _, letter := range letters {
			searchIdx := sort.Search(
				len(alphabet),
				func(i int) bool { return alphabet[i].Cmp(letter) >= 0 },
			)
			cleartext = append(cleartext, rune('A'+searchIdx))
		}

		writef(writer, "Case #%d: %s\n", 1+index, string(cleartext))
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
