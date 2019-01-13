package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	maxX := uint64(0)
	maxY := uint64(0)

	queries := readInt(scanner)
	for q := 0; q < queries; q++ {
		queryType := readString(scanner)
		if queryType == "+" {
			x := uint64(readInt64(scanner))
			y := uint64(readInt64(scanner))
			if y > x {
				maxX = max64(maxX, y)
				maxY = max64(maxY, x)
			} else {
				maxX = max64(maxX, x)
				maxY = max64(maxY, y)
			}
		} else if queryType == "?" {
			h := uint64(readInt64(scanner))
			w := uint64(readInt64(scanner))

			if h > w {
				h,w = w,h
			}

			if h >= maxY && w >= maxX {
				writef(writer, "YES\n")
			} else {
				writef(writer, "NO\n")
			}
		}
	}
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

func readInt(sc *bufio.Scanner) int {
	return int(readInt64(sc))
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}

func max64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}