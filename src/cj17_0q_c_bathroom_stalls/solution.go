package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type In struct {
	index int32
	N     int64
	K     int64
}

type Out struct {
	index int32
	maxLR int64
	minLR int64
}

type R struct {
	count int64
	size  int64
}

func solve(in In) (out Out) {
	kLeft := in.K
	rqNext := 0
	rq := make([]R, 1, 10)
	rq[rqNext] = R{1, in.N}

	for rq[rqNext].count < kLeft {
		// fmt.Println(fmt.Sprintf("%d %d %+v", kLeft, rqNext, rq[rqNext]))
		rNext := rq[rqNext]
		//	what the fork, maaaaaan, the record is copied out from array by value to my reference??? wuuuuuuut!!! o_O
		rLast := &rq[len(rq) - 1]
		kLeft -= rNext.count
		if rNext.size % 2 == 1 {
			if rNext.size > 1 {
				if (rLast.size < rNext.size / 2) {
					panic("queue is not sorted?")
				} else if (rLast.size == rNext.size / 2) {
					// fmt.Println(fmt.Sprintf("  incrementing %+v by %d", rLast, rNext.count*2))
					rLast.count += rNext.count * 2
				} else {
					rq = append(rq, R{rNext.count * 2, rNext.size / 2})
				}
			}
		} else {
			if (rLast.size < rNext.size / 2) {
				panic("queue is not sorted?")
			} else if (rLast.size == rNext.size / 2) {
				// fmt.Println(fmt.Sprintf("  incrementing %+v by %d", rLast, rNext.count))
				rLast.count += rNext.count
			} else {
				rq = append(rq, R{rNext.count, rNext.size / 2})
			}
			if rNext.size > 0 {
				rq = append(rq, R{rNext.count, (rNext.size - 1) / 2})
			}
		}

		rqNext++
	}

	final := rq[rqNext]
	return Out{in.index, final.size / 2, (final.size - 1) / 2}
}

func solveChannel(ins <-chan In, outs chan <- Out) {
	for in := range ins {
		outs <- solve(in)
	}
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

	caseCount := ReadInt32(scanner)

	cores := runtime.NumCPU()
	var ins = make(chan In, cores)
	var outs = make(chan Out, caseCount)
	for t := 0; t < cores; t++ {
		go solveChannel(ins, outs)
	}

	for index := int32(0); index < caseCount; index++ {
		ins <- In{index, ReadInt64(scanner), ReadInt64(scanner)}
	}
	close(ins)

	outsSlice := make([]Out, caseCount)
	for index := int32(0); index < caseCount; index++ {
		out := <-outs
		outsSlice[out.index] = out
	}
	close(outs)

	for _, out := range outsSlice {
		Writef(writer, "Case #%d: %d %d\n", 1 + out.index, out.maxLR, out.minLR)
	}
}

//	boring IO
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

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
