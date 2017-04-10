package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
)

type In struct {
	nums  []uint32
}

type Out struct {
	primes   []bool
}

func solve(in In) (out Out) {
	primes := make([]bool, len(in.nums))
	for i, num := range in.nums {
		if (num < 4) {
			primes[i] = num > 1;
			continue
		}

		if num % 2 == 0 {
			primes[i] = false;
			continue
		}

		prime := true
		limit := uint32(math.Ceil(math.Sqrt(float64(num))))
		for div := uint32(3); prime && div <= limit; div += 2 {
			if num % div == 0 {
				prime = false
			}
		}

		primes[i] = prime
	}
	
	return Out{primes}
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

	size := ReadInt16(scanner)
	nums := make([]uint32, size)
	for i := range nums {
		nums[i] = uint32(ReadInt64(scanner))
	}
	in := In{nums}

	out := solve(in)

	for _, prime := range out.primes {
		res := "Not prime"
		if (prime) {
			res = "Prime"
		}
		Writef(writer, "%s\n", res)

	}
}

//	boring IO
func ReadInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return int64(res)
}

func ReadInt32(sc *bufio.Scanner) int32 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(res)
}

func ReadInt16(sc *bufio.Scanner) int16 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 16)
	if err != nil {
		panic(err)
	}
	return int16(res)
}

func ReadInt8(sc *bufio.Scanner) int8 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 8)
	if err != nil {
		panic(err)
	}
	return int8(res)
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
