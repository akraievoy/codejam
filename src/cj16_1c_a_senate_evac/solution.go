package main

import (
	"bufio"
	"runtime"
	"sort"
	"strconv"
	"fmt"
	"os"
)


type In struct {
	index int
	nums  []int
}

type Out struct {
	index int
	evac  [][]rune
}

type Party struct {
	name  rune
	count int
}

type PartySort []Party

func (party_sort PartySort) Len() int {
	return len(party_sort)
}
func (party_sort PartySort) Swap(i, j int) {
	party_sort[i], party_sort[j] = party_sort[j], party_sort[i]
}
func (party_sort PartySort) Less(i, j int) bool {
	party_sorti, party_sortj := party_sort[i], party_sort[j]
	return party_sorti.count > party_sortj.count
}

func solve(in In) Out {
	parties := make([]Party, len(in.nums))
	totalLeft := 0
	for p := range parties {
		parties[p] = Party{'A' + int32(p), in.nums[p]}
		totalLeft += in.nums[p]
	}
	partiesLeft := len(in.nums)

	evac := make([][]rune, 0)
	for totalLeft > 0 {
		sort.Sort(PartySort(parties))
		 
		if partiesLeft > 2 || parties[0].count > parties[1].count {
			evac = append(evac, []rune{parties[0].name })
			(&parties[0]).count--
			if parties[0].count == 0 {
				partiesLeft -= 1
			}
			totalLeft--
		} else {
			evac = append(evac, []rune{parties[0].name, parties[1].name })
			(&parties[0]).count--
			if parties[1].count == 0 {
				partiesLeft -= 1
			}
			(&parties[1]).count--
			if parties[1].count == 0 {
				partiesLeft -= 1
			}
			totalLeft -= 2
		}
	}

	return Out{in.index, evac}
}

func solveInput(scanner *bufio.Scanner, writer *bufio.Writer) {
	caseCount := ReadInt(scanner)

	cores := runtime.NumCPU()
	var ins = make(chan In, cores)
	var outs = make(chan Out, caseCount)
	for t := 0; t < cores; t++ {
		go solveChannel(ins, outs)
	}

	outsSlice := make([]Out, caseCount)
	for index := 0; index < caseCount; index++ {
		size := int(ReadInt(scanner))
		nums := make([]int, size)
		for i := range nums {
			nums[i] = ReadInt(scanner)
		}
		in := In{index, nums}
		ins <- in
	}
	close(ins)

	for index := 0; index < caseCount; index++ {
		out := <-outs
		outsSlice[out.index] = out
	}
	close(outs)

	for _, out := range outsSlice {
		Writef(writer, "Case #%d:", 1 + out.index)

		for _,e := range out.evac {
			Writef(writer, " %s", string(e))
		}

		Writef(writer, "\n")
	}
}

func solveChannel(ins <-chan In, outs chan <- Out) {
	for in := range ins {
		outs <- solve(in)
	}
}

func min(a, b int) int {
	if (a < b) {
		return a
	}
	return b
}

func max(a, b int) int {
	if (a > b) {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if (a < b) {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if (a > b) {
		return a
	}
	return b
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

	solveInput(scanner, writer)
}

func ReadInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	res, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}

func ReadInt(sc *bufio.Scanner) int {
	return int(ReadInt64(sc))
}

func Writef(writer *bufio.Writer, formatStr string, values ...interface{}) {
	out := fmt.Sprintf(formatStr, values...)
	_, err := writer.WriteString(out)
	if err != nil {
		panic(err)
	}
}
