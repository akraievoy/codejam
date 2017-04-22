package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math/big"
)

type In struct {
	n int   //	[1..250]
	m int   //	[1..50]
	c []int //	[1..50]
}

type Out struct {
	waysToChange *big.Int
}

func solve(in In) Out {
	bigZero, bigOne := big.NewInt(0), big.NewInt(1)
	ways := make([][]*big.Int, in.m)
	for m := 0; m < in.m; m++ {
		ways[m] = make([]*big.Int, in.n+1)
		for n := 0; n <= in.n; n++ {
			if m == 0 {
				if n % in.c[m] == 0 {
					ways[m][n] = bigOne
				} else {
					ways[m][n] = bigZero
				}
			} else {
				if n == 0 {
					ways[m][n] = bigOne
				} else if (n < in.c[m]) {
					ways[m][n] = big.NewInt(0).Add(ways[m - 1][n], bigZero)
				} else {
					ways[m][n] = big.NewInt(0).Add(ways[m - 1][n], ways[m][n - in.c[m]])
				}
			}
		}
	}

	//fmt.Println(fmt.Sprintf("ways=%v", ways))
	
	return Out{ways[in.m - 1][in.n]}
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

	n,m := ReadInt(scanner),ReadInt(scanner)
	c := make([]int, m)
	for i := range c {
		c[i] = ReadInt(scanner)
	}
	in := In{n,m,c}

	out := solve(in)

	Writef(writer, "%s\n", out.waysToChange)
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

func min(a, b int) int {
	if (a < b) {
		return a
	}
	return b
}

func max(a, b int) int {
	if (a > b) {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if (a < b) {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if (a > b) {
		return a
	}
	return b
}
