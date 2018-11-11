package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	flags []bool
}

type Out struct {
	deletions int32
}

func solve(in In) (out Out) {
	flags := []bool{true}
	flags = append(flags, in.flags...)
	flags = append(flags, true)

	for i := 1; i < len(flags) - 1; i++ {
		if (flags[i - 1] || flags[i + 1]) {
			continue;
		}
		flags[i] = false
	}
	deletions := int32(0)
	for i := 1; i < len(flags) - 1; i++ {
		if (flags[i - 1] || flags[i + 1]) {
			continue;
		}
		deletions += 1
	}

	return Out{deletions}
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

	games := ReadInt16(scanner)

	for g := int16(0) ; g < games; g++ {
		size := ReadInt32(scanner)
		flags := make([]bool, size)
		for i := range flags {
			num_i := ReadInt8(scanner)
			flags[i] = num_i == 1
		}

		out := solve(In{flags})

		winner := "Bob"
		if (out.deletions % 2 == 1) {
			winner = "Alice"
		}
		Writef(writer, "%s\n", winner)
	}
}

//	boring IO
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
