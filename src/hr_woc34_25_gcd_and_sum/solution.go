package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

type In struct {
	A []int
	B []int
}

type Out struct {
	sum int
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

func solve(in In) Out {
	sort.Ints(in.A)
	sort.Ints(in.B)

	maxXForDivisor := make([]int, 1000001, 1000001)
	factors, powers := make([]int, 20), make([]int, 20)
	enumDeltas, enumPowers := make([]int, 20), make([]int, 20)

	prevX := -1
	for _, x := range in.A {
		if (prevX == x) {
			continue
		}
		prevX = x

		factors, powers = factorize(x, factors, powers)
		enumDeltas, enumPowers = enumDeltas[:0], enumPowers[:0]
		for range factors {
			enumDeltas = append(enumDeltas, 1)
			enumPowers = append(enumPowers, 0)
		}
		divisor := 1
		for true {
			maxXForDivisor[divisor] = max(maxXForDivisor[divisor], x)

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
			if (enumDeltas[rollPos] > 0) {
				divisor *= factors[rollPos]
			} else {
				divisor /= factors[rollPos]
			}
		}
	}

	maxGCD := 1
	maxSum := 2
	prevY := -1
	for _, y := range in.B {
		if (prevY == y) {
			continue
		}
		prevY = y

		factors, powers = factorize(y, factors, powers)
		enumDeltas, enumPowers = enumDeltas[:0], enumPowers[:0]
		for range factors {
			enumDeltas = append(enumDeltas, 1)
			enumPowers = append(enumPowers, 0)
		}
		divisor := 1
		for true {
			if (divisor >= maxGCD && maxXForDivisor[divisor] > 0) {
				if (divisor == maxGCD) {
					maxSum = max(maxSum, maxXForDivisor[divisor] + y)
				} else {
					maxSum = maxXForDivisor[divisor] + y
				}
				maxGCD = divisor
			} else {
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
			if (enumDeltas[rollPos] > 0) {
				divisor *= factors[rollPos]
			} else {
				divisor /= factors[rollPos]
			}
		}
	}

	return Out{maxSum}
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

	size := ReadInt(scanner)
	A := make([]int, size)
	for i := range A {
		A[i] = ReadInt(scanner)
	}
	B := make([]int, size)
	for i := range A {
		B[i] = ReadInt(scanner)
	}
	in := In{A,B}

	out := solve(in)

	Writef(writer, "%d\n", out.sum)
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

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}

func max(a, b int) int {
	if (a > b) {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
