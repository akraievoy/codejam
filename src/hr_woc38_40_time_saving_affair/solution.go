package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
	"container/heap"
	"math"
)

type edge struct {
	into  uint16
	delay uint16
}

type caseInput struct {
	n uint16
	k uint32

	egressRangeStarts []uint32
	edges             []edge
}

type fullEdge struct {
	from  uint16
	into  uint16
	delay uint16
}

type BySource []fullEdge

func (a BySource) Len() int      { return len(a) }
func (a BySource) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySource) Less(i, j int) bool {
	if a[i].from == a[j].from {
		return a[i].delay < a[j].delay
	} else {
		return a[i].from < a[j].from
	}
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	n, k, m := uint16(readInt(scanner)), uint32(readInt(scanner)), readInt(scanner)

	fullEdges := make([]fullEdge, 0)
	for ei := 0; ei < m; ei++ {
		fe := fullEdge{
			uint16(readInt(scanner)) - 1,
			uint16(readInt(scanner)) - 1,
			uint16(readInt(scanner)),
		}
		if fe.from == fe.into {
			continue
		}
		fullEdges = append(fullEdges, fe)
		fullEdges = append(fullEdges, fullEdge{fe.into, fe.from, fe.delay})
	}

	sort.Sort(BySource(fullEdges))

	egressRangeStarts := make([]uint32, n+1)
	edges := make([]edge, len(fullEdges))
	egressIdx := uint16(0)

	for ei, e := range fullEdges {
		for egressIdx < e.from {
			egressIdx += 1
			egressRangeStarts[egressIdx] = uint32(ei)
		}
		edges[ei] = edge{e.into, e.delay}
	}
	for egressIdx < n {
		egressIdx += 1
		egressRangeStarts[egressIdx] = uint32(len(fullEdges))
	}

	return caseInput{n, k, egressRangeStarts, edges}
}

type caseOutput struct {
	minTime int
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", out.minTime)
}

type dijkstraDistance struct {
	finite bool
	time   uint32
}

//	TODO pointer receivers and other deep magick?
func (this dijkstraDistance) Less(that dijkstraDistance) bool {
	return this.finite && (!that.finite || this.time < that.time)
}

type dijkstraQueueElem struct {
	node uint16
	dist dijkstraDistance
}

type dijkstraQueue []dijkstraQueueElem

func (a dijkstraQueue) Len() int      { return len(a) }
func (a dijkstraQueue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a dijkstraQueue) Less(i, j int) bool {
	return a[i].dist.Less(a[j].dist)
}
func (pa *dijkstraQueue) Push(x interface{}) {
	*pa = append(*pa, x.(dijkstraQueueElem))
}
func (pa *dijkstraQueue) Pop() interface{} {
	old := *pa
	n := len(old)
	x := old[n-1]
	*pa = old[0 : n-1]
	return x
}

func solveCase(in caseInput) caseOutput {
	fromIdx := uint16(0)
	intoIdx := uint16(in.n - 1)

	targetDistances := make([]dijkstraDistance, in.n)
	for i := range targetDistances {
		targetDistances[i] = dijkstraDistance{false, math.MaxUint32}
	}

	queue := &dijkstraQueue{}
	heap.Init(queue)

	heap.Push(queue, dijkstraQueueElem{fromIdx, dijkstraDistance{true, 0}})

	for queue.Len() > 0 {
		qElem := heap.Pop(queue).(dijkstraQueueElem)
		if !qElem.dist.Less(targetDistances[qElem.node]) {
			continue
		}
		if !qElem.dist.Less(targetDistances[intoIdx]) {
			continue
		}
		targetDistances[qElem.node] = qElem.dist

		rangeStart := in.egressRangeStarts[qElem.node]
		rangeEnd := in.egressRangeStarts[qElem.node+1]
		for edgeIdx := rangeStart; edgeIdx < rangeEnd; edgeIdx++ {
			e := in.edges[edgeIdx]

			arrivalTime := qElem.dist.time + uint32(e.delay)

			var goodToGoTime uint32
			if (arrivalTime / in.k) % 2 == 0 || e.into == intoIdx {
				goodToGoTime = arrivalTime
			} else {
				goodToGoTime = (arrivalTime / in.k + 1) * in.k
			}

			newDistance := dijkstraDistance{true, goodToGoTime}
			if !newDistance.Less(targetDistances[e.into]) {
				continue
			}
			if !newDistance.Less(targetDistances[intoIdx]) {
				continue
			}
			heap.Push(queue, dijkstraQueueElem{e.into, newDistance})
		}
	}

	if targetDistances[intoIdx].finite {
		return caseOutput{int(targetDistances[intoIdx].time)}
	}
	return caseOutput{-1}
}

func solveSequential(scanner *bufio.Scanner, writer *bufio.Writer) {
	writeCaseOutput(writer, solveCase(readCaseInput(scanner)))
}

func main() {
	var scanner *bufio.Scanner
	if len(os.Getenv("CODEJAM_INPUT")) > 0 {
		reader, err := os.Open(os.Getenv("CODEJAM_INPUT"))
		if err != nil {
			panic(err)
		}
		defer reader.Close()
		scanner = bufio.NewScanner(reader)
	} else if len(os.Args) > 1 {
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