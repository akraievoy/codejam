package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	basePath := "src/a/"

	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	caseCount := 10
	testIndex, testInput, testOutput, logWriter, closeFunc := probeAndOpen(basePath)
	defer closeFunc()

	writef(logWriter, "test# = %02d caseCount = %d seed = %d\n", testIndex, caseCount, seed)

	writef(testInput, "%d\n", caseCount)
	for caseI := 1; caseI <= caseCount; caseI++ {

		writef(logWriter, "Case #%d -- input generation...\n", caseI)
		l := r.Intn(1000000 * 1000)
		r := r.Intn(1000000 * 1000)

		writef(testInput, "%d %d\n", l, r)

		writef(logWriter, "Case #%d -- solution...\n", caseI)
		i := 1
		for ; i <= l || i <= r; i++ {
			if l >= r {
				l -= i
			} else {
				r -= i
			}
		}

		writef(logWriter, "Case #%d -- output generation...\n", caseI)
		writef(testOutput, "Case #%d: %d %d %d\n", caseI, i-1, l, r)
	}
}

func probeAndOpen(basePath string) (int, *bufio.Writer, *bufio.Writer, *bufio.Writer, func()) {
	testIndex := 0
	for {
		if _, err := os.Stat(fmt.Sprintf("%s/%02d.in", basePath, testIndex)); err != nil {
			break
		}
		testIndex += 1
	}

	testInput, err := os.Create(fmt.Sprintf("%s/%02d.in", basePath, testIndex))
	if err != nil {
		panic(err)
	}
	testOutput, err := os.Create(fmt.Sprintf("%s/%02d.out", basePath, testIndex))
	if err != nil {
		panic(err)
	}
	testInputWriter := bufio.NewWriter(testInput)
	testOutputWriter := bufio.NewWriter(testOutput)
	logWriter := bufio.NewWriterSize(os.Stderr, 1)

	// close on exit and check for returned errors
	closeFunc := func() {
		if err := testInputWriter.Flush(); err != nil {
			panic(err)
		}
		if err := testInput.Close(); err != nil {
			panic(err)
		}
		if err := testOutputWriter.Flush(); err != nil {
			panic(err)
		}
		if err := testOutput.Close(); err != nil {
			panic(err)
		}
		if err := logWriter.Flush(); err != nil {
			panic(err)
		}
	}
	return testIndex, testInputWriter, testOutputWriter, logWriter, closeFunc
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	_, err := fmt.Fprintf(writer, formatStr, values...)
	if err != nil {
		panic(err)
	}
}
