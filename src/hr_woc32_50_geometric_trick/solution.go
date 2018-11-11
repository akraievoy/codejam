package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	abc []rune
}

type Out struct {
	triples int
}

var PRIMES = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
	31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
	73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
	233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
	283, 293, 307, 311, 313, 317, 331, 337, 347, 349,
	353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
	419, 421, 431, 433, 439, 443, 449, 457, 461, 463,
	467, 479, 487, 491, 499, 503, 509, 521, 523, 541,
	547, 557, 563, 569, 571, 577, 587, 593, 599, 601,
	607, 613, 617, 619, 631, 641, 643, 647, 653, 659,
	661, 673, 677, 683, 691, 701, 709, 719, 727, 733,
	739, 743, 751, 757, 761, 769, 773, 787, 797, 809,
	811, 821, 823, 827, 829, 839, 853, 857, 859, 863,
	877, 881, 883, 887, 907, 911, 919, 929, 937, 941,
	947, 953, 967, 971, 977, 983, 991, 997 }

func factorize(num int, factors []int, powers []int) ([]int, []int) {
	factors = factors[:0]
	powers = powers[:0]
	divIdx := 0

	if num == 1 {
		factors = append(factors, 1)
		powers = append(powers, 1)
		return factors, powers
	}

	for num > 1 {
		var div int
		if divIdx == len(PRIMES) {
			div = num
		} else {
			div = PRIMES[divIdx]
			divIdx++
		}
		if num % div > 0 {
			continue
		}
		factors = append(factors, div)
		powers = append(powers, 1)
		num /= div
		for num % div == 0 {
			powers[len(powers) - 1]++
			num /= div
		}
	}

	return factors, powers
}

func solve(in In) (out Out) {
	iFlags := make([]bool, len(in.abc) + 1)
	jArr := make([]int, 0, len(in.abc))
	kFlags := make([]bool, len(in.abc) + 1)

	for index, letter := range in.abc {
		if letter == 'a' {
			iFlags[index + 1] = true
		} else if letter == 'b' {
			jArr = append(jArr, index + 1)
		} else if letter == 'c' {
			kFlags[index + 1] = true
		} else {
			panic(fmt.Sprintf("unexpected rune %c at index %d", letter, index))
		}
	}

	factors, powers, enumDeltas, enumPowers :=
		make([]int, 20), make([]int, 20), make([]int, 20), make([]int, 20)
	triples := 0
	for _, j := range jArr {
		factors, powers = factorize(j, factors, powers)
		i, k := 1, j * j

		enumDeltas, enumPowers = enumDeltas[:0], enumPowers[:0]
		for fi := range factors {
			enumDeltas = append(enumDeltas, 1)
			enumPowers = append(enumPowers, 0)
			powers[fi] *= 2
		}

		for true {
			if i < len(iFlags) && k < len(kFlags) && iFlags[i] && kFlags[k] {
				triples += 1
			}

			rollPos := 0
			rollPow := enumPowers[rollPos] + enumDeltas[rollPos]
			for rollPos < len(enumPowers) && (rollPow < 0 || rollPow > powers[rollPos]) {
				enumDeltas[rollPos] *= -1
				rollPos++
				if rollPos < len(enumPowers) {
					rollPow = enumPowers[rollPos] + enumDeltas[rollPos]
				}
			}

			if rollPos >= len(enumPowers) {
				break
			}
			enumPowers[rollPos] = rollPow
			if enumDeltas[rollPos] > 0 {
				i *= factors[rollPos]
				k /= factors[rollPos]
			} else {
				i /= factors[rollPos]
				k *= factors[rollPos]
			}
		}

	}

	return Out{triples}
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
	scanner.Buffer(make([]byte, 1000000), 1000000)
	scanner.Split(bufio.ScanWords)

	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	_, text := ReadInt(scanner), ReadString(scanner)
	in := In{[]rune(text)}

	out := solve(in)

	Writef(writer, "%d\n", out.triples)
}

func ReadInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func ReadInt(sc *bufio.Scanner) int {
	return int(ReadInt64(sc))
}

func ReadString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
