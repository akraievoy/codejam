package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math/rand"
)

func main() {
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	r := rand.New(rand.NewSource(99))
	size := 100000

	Writef(writer, "%d\n", size)

	Writef(writer, "%d\n", r.Int31n(1000000001))

	for i := 1; i < size; i++ {
		Writef(writer, "%d", r.Int31n(1000000001))
		if i + 1 < size {
			Writef(writer, " ")
		}
	}
	Writef(writer, "\n")

	for i := 1; i < size; i++ {
		Writef(writer, "%d", r.Int31n(2000000001) - 1000000000)
		if i + 1 < size {
			Writef(writer, " ")
		}
	}
	Writef(writer, "\n")
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
