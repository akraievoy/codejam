package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type NK struct {
	N, K int
}

type In struct {
	precomputed [][]int
	nk          []NK
}

type Out struct {
	counts []int
}

const MOD = 1000000007

func precompute() [][]int {
	lim := 3001
	result := make([][]int, lim + 1)
	for n := 1; n <= lim; n++ {
		result[n] = make([]int, lim + 1)
		if n == 1 {
			result[n][0] = 1
		} else {
			for k := 1; k <= n; k++ {
				if result[n - 1][k] > 0 && 2 * k - n + 2 > 0 {
					result[n][k] =
						result[n - 1][k] * (2 * k - n + 2) % MOD
				}
				if result[n - 1][k - 1] > 0 && n - k > 0 {
					result[n][k] =
						(result[n][k] + result[n - 1][k - 1] * (2 * (n - k))) % MOD
				}
			}
		}
	}
	return result
}

func solve(in In) (out Out) {
	counts := make([]int, len(in.nk))
	for q, nk := range in.nk {
		sum := 0
		for _, v := range in.precomputed[nk.N][nk.K:] {
			sum = (sum + v) % MOD
		}
		counts[q] = sum
	}
	return Out{counts}
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

	q := ReadInt(scanner)
	nk := make([]NK, q)
	for i := range nk {
		nk[i] = NK{ReadInt(scanner), ReadInt(scanner)}
	}
	in := In{precompute(), nk}

	out := solve(in)

	for _, c := range out.counts {
		Writef(writer, "%d\n", c)
	}
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
