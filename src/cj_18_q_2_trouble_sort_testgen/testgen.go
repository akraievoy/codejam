package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
)

func main() {
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	r := rand.New(rand.NewSource(99))

	testCount := 100

	Writef(writer, "%d\n", testCount)

	for test := 0; test < testCount; test++ {
		size := int(r.Int31n(100001-3)+3)
		Writef(writer, "%d\n", size)
		for i := 0; i < size; i++ {
			Writef(writer, "%d", r.Int31n(1000000001))
			if i + 1 < size {
				Writef(writer, " ")
			} else {
				Writef(writer, "\n")
			}
		}
	}
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
