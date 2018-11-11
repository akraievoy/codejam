package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"container/heap"
)

type hop struct {
	length  uint32
	opStamp uint32
}

func (h *hop) String() string {
	return fmt.Sprintf("{%d@%d}", h.length, h.opStamp)
}

type distance struct {
	sum     uint64
	opStamp uint32
}

func (d *distance) String() string {
	return fmt.Sprintf("{%d@%d}", d.sum, d.opStamp)
}

func (d *distance) plus(prefix uint64) distance {
	return distance{d.sum + prefix, d.opStamp}
}

func (d *distance) longer(that *distance) bool {
	if d.sum == that.sum {
		return d.opStamp > that.opStamp
	}
	return d.sum > that.sum
}

type route struct {
	hops []hop
	dist distance
}

func (r *route) String() string {
	return fmt.Sprintf("{%d@%d[%d]}", r.dist.sum, r.dist.opStamp, len(r.hops))
}

func (r *route) add(h hop) {
	if r.hops == nil {
		r.hops = make([]hop, 0)
	}
	r.hops = append(r.hops, h)
	r.dist.sum += uint64(h.length)
	r.dist.opStamp = h.opStamp
}

func (r *route) remove() {
	if r.hops == nil || len(r.hops) == 0 {
		panic("trying to remove way too much")
	}

	r.dist.sum -= uint64(r.hops[len(r.hops)-1].length)
	r.hops = r.hops[:len(r.hops)-1]
	if len(r.hops) > 0 {
		r.dist.opStamp = r.hops[len(r.hops) - 1].opStamp
	} else {
		r.dist.opStamp = 0
	}
}

type routeHeap []route

func (h routeHeap) Len() int      { return len(h) }
func (h routeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h routeHeap) Less(i, j int) bool {
	return h[i].dist.longer(&h[j].dist)
}
func (h *routeHeap) Push(x interface{}) {
	*h = append(*h, x.(route))
}
func (h *routeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type routeGroup struct {
	nodeIdx uint32
	//	updateable
	heapIdx uint32
	routes  routeHeap
}

func (rg *routeGroup) longest() distance {
	if rg.routes == nil || len(rg.routes) == 0 {
		return distance{0, 0}
	}
	return rg.routes[0].dist
}

func (rg *routeGroup) addToLongest(h hop) {
	if rg.routes == nil {
		rg.routes = routeHeap(make([]route, 0))
	}
	if len(rg.routes) == 0 {
		rg.routes = append(rg.routes, route{})
	}
	rg.routes[0].add(h)
}

func (rg *routeGroup) add(h hop) {
	if rg.routes == nil {
		rg.routes = routeHeap(make([]route, 0))
	}
	newRoute := route{}
	newRoute.add(h)
	heap.Push(&rg.routes, newRoute)
}

func (rg *routeGroup) removeFromLongest() {
	if rg.routes == nil || len(rg.routes) == 0 {
		panic(fmt.Sprintf("trying to delete way too much from %d", rg.nodeIdx))
	}
	rg.routes[0].remove()
	if len(rg.routes[0].hops) == 0 {
		heap.Remove(&rg.routes, 0)
	} else {
		heap.Fix(&rg.routes, 0)
	}
}

type groupHeap []*routeGroup

func (h groupHeap) Len() int { return len(h) }
func (h groupHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = uint32(i)
	h[j].heapIdx = uint32(j)
}
func (h groupHeap) Less(i, j int) bool {
	iLongest := h[i].longest()
	jLongest := h[j].longest()
	return iLongest.longer(&jLongest)
}
func (h *groupHeap) Push(x interface{}) {
	panic("Push for routeGroup Heap not needed")
	//xRG := x.(routeGroup)
	//xRG.heapIdx = uint32(len(*h))
	//*h = append(*h, xRG)
}
func (h *groupHeap) Pop() interface{} {
	panic("Pop for routeGroup Heap not needed")
	//old := *h
	//n := len(old)
	//x := old[n-1]
	//x.heapIdx = math.MaxUint32
	//*h = old[0 : n-1]
	//return x
}

type caseInput struct {
	n                  uint32
	cycleLenPrefixSums []uint64
	groups             []routeGroup
	groupH             groupHeap
	m                  uint32
	scanner            *bufio.Scanner
}

func (pIn *caseInput) farthest(from uint32) (uint32, uint64) {
	in := *pIn
	maxGroupLen := in.groupH[0].longest()

	globalInto, globalLongest := from, in.groups[from].longest()

	//	route group heap is partially ordered: let's preheat with constant number of probes
	const heapProbes = uint32(15)
	for heapProbe := uint32(0); heapProbe < heapProbes && heapProbe < in.n; heapProbe++ {
		into := in.groupH[heapProbe].nodeIdx
		var cycleHops uint32
		if into >= from {
			cycleHops = into - from
		} else {
			cycleHops = in.n + into - from
		}
		cycleLen := in.cycleLenPrefixSums[from+cycleHops] - in.cycleLenPrefixSums[from]

		probeGroupLongest := in.groupH[heapProbe].longest()
		probeLongest := probeGroupLongest.plus(cycleLen)
		if probeLongest.longer(&globalLongest) {
			globalLongest = probeLongest
			globalInto = into
		}
	}

	//	this is potentially root-node-global scan
	for cycleHops := in.n - 1; cycleHops > 0; cycleHops-- {
		cycleLen := in.cycleLenPrefixSums[from+cycleHops] - in.cycleLenPrefixSums[from]
		into := (from + cycleHops) % in.n

		//	cycleLen values do always descend, so if we went under once we'll never recover to improve
		if cycleLen+maxGroupLen.sum < globalLongest.sum {
			break
		}

		currentGroupLongest := in.groups[into].longest()
		currentLongest := currentGroupLongest.plus(cycleLen)
		if currentLongest.longer(&globalLongest) {
			globalLongest = currentLongest
			globalInto = into
		}
	}

	return globalInto, globalLongest.sum
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	n := uint32(readInt(scanner))

	cycle := make([]uint64, n)
	for i := range cycle {
		cycle[i] = uint64(readInt(scanner))
	}

	cycleLenPrefixSums := make([]uint64, 2*n+1)
	for i := range cycle {
		cycleLenPrefixSums[i+1] = cycleLenPrefixSums[i] + cycle[i]
	}
	for i := range cycle {
		cycleLenPrefixSums[int(n)+i+1] = cycleLenPrefixSums[int(n)+i] + cycle[i]
	}

	groups := make([]routeGroup, n)
	groupH := groupHeap(make([]*routeGroup, n))
	for i := range cycle {
		groups[i] = routeGroup{uint32(i), uint32(i), nil}
		groupH[i] = &(groups[i])
	}
	m := uint32(readInt(scanner))

	return caseInput{n, cycleLenPrefixSums, groups, groupH, m, scanner}
}

type caseOutput struct {
	queries []uint64
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	for _, q := range out.queries {
		writef(writer, "%d\n", q)
	}
}

func solveCase(in caseInput) caseOutput {
	queries := make([]uint64, 0)

	for opStamp := uint32(1); opStamp <= in.m; opStamp++ {
		queryType := uint8(readInt(in.scanner))
		x := uint32(readInt(in.scanner) - 1)
		switch queryType {
		case 1:
			h := hop{uint32(readInt(in.scanner)), opStamp}
			y, _ := in.farthest(x)
			in.groups[y].addToLongest(h)
			heap.Fix(&in.groupH, int(in.groups[y].heapIdx))
		case 2:
			h := hop{uint32(readInt(in.scanner)), opStamp}
			in.groups[x].add(h)
			heap.Fix(&in.groupH, int(in.groups[x].heapIdx))
		case 3:
			y, _ := in.farthest(x)
			in.groups[y].removeFromLongest()
			heap.Fix(&in.groupH, int(in.groups[x].heapIdx))
		case 4:
			_, dist := in.farthest(x)
			queries = append(queries, dist)
		default:
			panic(fmt.Sprintf("opStamp %d queryType %d ", opStamp, queryType))
		}
	}

	return caseOutput{queries}
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