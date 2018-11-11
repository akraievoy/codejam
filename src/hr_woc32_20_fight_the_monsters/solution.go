package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

type In struct {
	hit int
	t   int
	h   []int
}

type Out struct {
	maxKilledZombies int
}

func solve(in In) (out Out) {
	sort.Ints(in.h)

	i, t := 0, in.t
	for i < len(in.h) && t > 0 {
		toKill := (in.h[i] + in.hit - 1) / in.hit
		if toKill <= t {
			i += 1
			t -= toKill
		} else {
			break
		}
	}

	return Out{maxKilledZombies: i}
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

	n, hit, t := ReadInt(scanner), ReadInt(scanner), ReadInt(scanner)
	h := make([]int, n)
	for i := range h {
		h[i] = ReadInt(scanner)
	}
	in := In{hit, t, h}

	out := solve(in)

	Writef(writer, "%d\n", out.maxKilledZombies)
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

