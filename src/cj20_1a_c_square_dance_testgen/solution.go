package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	r := rand.New(rand.NewSource(99))
	size := 5

	Writef(writer, "1\n")
	Writef(writer, "%d %d\n", size, size)

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			Writef(writer, "%d", 1 + r.Int31n(100000))
			if col+1 < size {
				Writef(writer, " ")
			}
		}
		Writef(writer, "\n")
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

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
