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
		n := 2 + r.Intn(10)
		k := 2 + r.Intn(3)
		writef(testInput, "%d %d\n", n, k)
		a := make([]int, n, n)
		for i := range a {
			a[i] = 1 + r.Intn(5)
			writef(testInput, "%d ", a[i])
		}
		writef(testInput, "\n")

		writef(logWriter, "Case #%d -- solution...\n", caseI)
		countdowns := 0
		for end := k - 1; end < n; end++ {
			countdown := true
			for offs := k; offs > 0; offs-- {
				if a[end-offs+1] != offs {
					countdown = false
					break
				}
			}
			if countdown {
				countdowns++
			}
		}

		writef(logWriter, "Case #%d -- output generation...\n", caseI)
		writef(testOutput, "Case #%d: %d\n", caseI, countdowns)
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
