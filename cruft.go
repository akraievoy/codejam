package main

import (
	"bufio"
	"strconv"
	"fmt"
	"os"
)


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

func minInt32(a, b int32) int32 {
	if (a < b) {
		return a
	}
	return b
}

func maxInt32(a, b int32) int32 {
	if (a > b) {
		return a
	}
	return b
}

func minUint32(a, b uint32) uint32 {
	if (a < b) {
		return a
	}
	return b
}

func maxUint32(a, b uint32) uint32 {
	if (a > b) {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if (a < b) {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if (a > b) {
		return a
	}
	return b
}

func minUint64(a, b uint64) uint64 {
	if (a < b) {
		return a
	}
	return b
}

func maxUint64(a, b uint64) uint64 {
	if (a > b) {
		return a
	}
	return b
}

func minInt16(a, b int16) int16 {
	if (a < b) {
		return a
	}
	return b
}

func maxInt16(a, b int16) int16 {
	if (a > b) {
		return a
	}
	return b
}

func minUint16(a, b uint16) uint16 {
	if (a < b) {
		return a
	}
	return b
}

func maxUint16(a, b uint16) uint16 {
	if (a > b) {
		return a
	}
	return b
}

func minInt8(a, b int8) int8 {
	if (a < b) {
		return a
	}
	return b
}

func maxInt8(a, b int8) int8 {
	if (a > b) {
		return a
	}
	return b
}

func minUint8(a, b uint8) uint8 {
	if (a < b) {
		return a
	}
	return b
}

func maxUint8(a, b uint8) uint8 {
	if (a > b) {
		return a
	}
	return b
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

	solveInput(scanner, writer)
}

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
