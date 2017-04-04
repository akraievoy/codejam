package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readUint16(sc *bufio.Scanner) uint16 {
	sc.Scan()
	res, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return uint16(res)
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}

func main() {
	var scanner *bufio.Scanner;
	if len(os.Args) > 1 {
		reader, err := os.Open(os.Args[1])
		if (err != nil) {
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

}
