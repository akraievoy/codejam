package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
)

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	n := readFloat64(scanner)
	r := readFloat64(scanner)

	angle := 2 * math.Pi / n

	z := math.Sqrt(math.Pow(1 - math.Cos(angle), 2) + math.Pow(math.Sin(angle), 2))

	x := (z * r) / (2 - z)

	writef(writer, "%.9f\n", x)
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
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solveSequential(scanner, writer)
}

func readFloat64(sc *bufio.Scanner) float64 {
	sc.Scan()
	res, err := strconv.ParseFloat(sc.Text(), 64)
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