package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solveSequential(j Jam) {
	T, _, M := j.Ri(), j.Ri(), j.Ri()
	modulos := []int64{17,16,13,11,9,7,5}
	remainders := make([]int64, len(modulos))
	for t := int64(0); t < T; t++ {
		for moduloI, m := range modulos {
			for i := 0; i < 18; i++ {
				if i == 0 {
					j.W("%d", m)
				} else {
					j.W(" %d", m)
				}
			}
			j.W("\n")
			j.Fl()

			remRes := int64(0)
			for i := 0; i < 18; i++ {
				rem := j.Ri()
				if rem < 0 {
					panic("judge does not want us anymore")
				}
				remRes = ( remRes + rem ) % m
			}
			remainders[moduloI] = remRes
		}

		foundM := false
		for m := remainders[0]; m <= M; m += modulos[0] {
			valid := true
			for moduloI, modulo := range modulos {
				if moduloI == 0 {
					continue
				}
				if m % modulo != remainders[moduloI] {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}
			foundM = true

			j.W("%d\n", m)
			j.Fl()
		}

		if !foundM {
			panic("failed to recover m")
		}

		testCaseReply := j.Rs()
		if testCaseReply == "1" {
			// cool we're still on track
		} else {
			panic(testCaseReply)
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
		defer func() {
			if err := reader.Close(); err != nil {
				panic(err)
			}
		}()
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	defer func() {
		if err := writer.Flush(); err != nil {
			panic(err)
		}
	}()

	solveSequential(&jam{scanner, writer})
}

type Jam interface {
	Sc() *bufio.Scanner
	Wr() *bufio.Writer
	Fl()

	Rs() string
	Ri() int64
	Rf() float64

	W(format string, values ...interface{})
}

type jam struct {
	sc *bufio.Scanner
	wr *bufio.Writer
}

func (j *jam) Sc() *bufio.Scanner {
	return j.sc
}

func (j *jam) Wr() *bufio.Writer {
	return j.wr
}

func (j *jam) Fl() {
	if err := j.wr.Flush(); err != nil {
		panic(err)
	}
}

func (j *jam) Rs() string {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	return j.sc.Text()
}

func (j *jam) Ri() int64 {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	res, err := strconv.ParseInt(j.sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	return res
}

func (j *jam) Rf() float64 {
	j.sc.Scan()
	res, err := strconv.ParseFloat(j.sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (j *jam) W(format string, values ...interface{}) {
	out := fmt.Sprintf(format, values...)
	_, err := j.wr.WriteString(out)
	if err != nil {
		panic(err)
	}
}

//	GoLang shorthand methods for math go below
//	TODO wipe unused methods before submitting

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
