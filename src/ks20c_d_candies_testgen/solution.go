package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	basePath := "src/d/"

	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	caseCount := 1
	testIndex, testInput, testOutput, logWriter, closeFunc := probeAndOpen(basePath)
	defer closeFunc()

	writef(logWriter, "test# = %02d caseCount = %d seed = %d\n", testIndex, caseCount, seed)

	writef(testInput, "%d\n", caseCount)
	for caseI := 1; caseI <= caseCount; caseI++ {

		writef(logWriter, "Case #%d -- input generation...\n", caseI)
		n := 200000
		q := 100000
		writef(testInput, "%d %d\n", n, q)

		arr := make([]int64, n, n)
		for i := 0; i < n; i++ {
			arr[i] = 1+r.Int63n(100)
			writef(testInput, "%d ", arr[i])
		}
		writef(testInput, "\n")


		writef(logWriter, "Case #%d -- solution...\n", caseI)
		sweetenessSum := int64(0)
		for qi := 0; qi < q; qi++ {
			if qi == q/2 {
				idx := r.Intn(n)
				upd := 1+r.Int63n(100)
				arr[idx] = upd
				writef(testInput, "U %d %d\n", idx+1, upd)
			} else {
				querySweeteness, start, end := int64(0), r.Intn(n), r.Intn(n)
				if start > end {
					start, end = end, start
				}
				writef(testInput, "Q %d %d\n", start+1, end+1)
				mult := int64(1)
				for i := start; i <= end; i++ {
					querySweeteness += mult * arr[i]
					mult = TernI8(mult > 0, -(mult + 1), -(mult - 1))
				}
				sweetenessSum += querySweeteness
			}
		}

		writef(logWriter, "Case #%d -- output generation...\n", caseI)
		writef(testOutput, "Case #%d: %d\n", caseI, sweetenessSum)
	}
}

func TernI8(b bool, t, f int64) int64 {
	if b {
		return t
	}
	return f
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
