package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
)

type caseInput struct {
	index     int
	lollipops int
	scanner   *bufio.Scanner
	writer    *bufio.Writer
}

func readCaseInput(scanner *bufio.Scanner, writer *bufio.Writer, index int) caseInput {
	return caseInput{index, readInt(scanner), scanner, writer}
}

type state struct {
	dislikeStats []int
	sold         []bool
	sellQueue    []int
}

func (s state) Len() int           { return len(s.sold) }
func (s state) Swap(i, j int)      { s.sellQueue[i], s.sellQueue[j] = s.sellQueue[j], s.sellQueue[i] }
func (s state) Less(i, j int) bool {
	if s.sold[s.sellQueue[i]] {
		if s.sold[s.sellQueue[j]] {
			return false
		} else {
			return false
		}
	} else {
		if s.sold[s.sellQueue[j]] {
			return true
		} else {
			return s.dislikeStats[s.sellQueue[i]] > s.dislikeStats[s.sellQueue[j]]
		}
	}
}

func solveCase(in caseInput) bool {
	L := in.lollipops

	s := state{
		make([]int, L, L),
		make([]bool, L, L),
		make([]int, L, L),
	}
	for l := range s.sellQueue {
		s.sellQueue[l] = l;
	}
	custLikes := make([]int, 0, L)
	custLikeFlags := make([]bool, L)

	for cust := 0; cust < L; cust++ {
		custLikes = custLikes[:0]
		for l := range custLikeFlags {
			custLikeFlags[l] = false
		}
		d := readInt(in.scanner)
		if d < 0 {
			return false;
		}
		for l := 0; l < d; l++ {
			liked := readInt(in.scanner)
			custLikes = append(custLikes, liked)
			custLikeFlags[liked] = true
		}
		for d := range s.dislikeStats {
			s.dislikeStats[d] += 1
		}
		for _, l := range custLikes {
			s.dislikeStats[l] -= 1
		}
		sort.Sort(s)
		sold := false
		// fmt.Fprintf(os.Stderr, "\n\nqueue\n\t%v\nsold\n\t%v\nstats\n\t%v\ncustLikeFlags\n\t%v\n", s.sellQueue, s.sold, s.dislikeStats, custLikeFlags)
		for _, q := range s.sellQueue {
			if custLikeFlags[q] && !s.sold[q] {
				sold = true
				s.sold[q] = true
				// fmt.Fprintf(os.Stderr, "SOLD\t%d\n", q)
				writef(in.writer, "%d\n", q)
				in.writer.Flush()
				break;
			}
		}
		if !sold {
			// fmt.Fprintf(os.Stderr, "SOLD NONE\t-1\n")
			writef(in.writer, "-1\n")
			in.writer.Flush()
		}
	}

	return true
}

//	everything below is reusable boilerplate
func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseCount := readInt(scanner)
	for index := 0; index < caseCount; index++ {
		if !solveCase(readCaseInput(scanner, writer, index)) {
			break
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
		defer reader.Close()
		scanner = bufio.NewScanner(reader)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Split(bufio.ScanWords)

	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	solveSequential(scanner, writer)
}

func readInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func readInt(sc *bufio.Scanner) int {
	return int(readInt64(sc))
}

func writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}