package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math/big"
)

type In struct {
	K int
	A []int
}

type Out struct {
	MaxAnd int
	Count  *big.Int
}

func count(A []int, mask int) int {
	count := 0
	for _, a := range A {
		if a & mask == mask {
			count += 1
		}
	}

	return count
}

func solve(in In) (out Out) {
	maxMask := 0
	maxCount := len(in.A)
	bit := 1
	for bit <= 1000000000000000000 {
		countForBit := count(in.A, bit)
		if countForBit >= in.K {
			maxMask = bit
			maxCount = countForBit
		}
		bit *= 2
	}
	for bit > 0 {
		maskExt := maxMask | bit
		countForMaskExt := count(in.A, maskExt)
		if countForMaskExt >= in.K {
			maxMask = maskExt
			maxCount = countForMaskExt
		}
		bit /= 2
	}

	fullChooseCount :=
		big.NewInt(1).Div(
			big.NewInt(1).MulRange(int64(in.K+1), int64(maxCount)),
			big.NewInt(1).MulRange(1, int64(maxCount - in.K)))

	return Out{maxMask, fullChooseCount}
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

	N := ReadInt(scanner)
	K := ReadInt(scanner)
	A := make([]int, N)
	for i := range A {
		A[i] = ReadInt(scanner)
	}
	in := In{K, A}

	out := solve(in)

	Writef(writer, "%d\n%s\n", out.MaxAnd, big.NewInt(0).Mod(out.Count, big.NewInt(1000000007)))
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
