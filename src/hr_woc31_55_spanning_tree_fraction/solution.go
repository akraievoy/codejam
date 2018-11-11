package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	/*
	"sort"
	*/
	"container/heap"
)

type Link struct {
	from  uint32
	into  uint32
	a     int8
	b     int8
	ratio float64
}

/*
type LinkSort []Link

func (ls LinkSort) Len() int {
	return len(ls)
}
func (ls LinkSort) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}
func (ls LinkSort) Less(i, j int) bool {
	lsi, lsj := ls[i], ls[j]
	if (lsi.from == lsj.from) {
		if (lsi.into == lsj.into) {
			if (lsi.b == lsj.b) {
				return lsi.a > lsj.a
			}
			return lsi.b < lsj.b
		}
		return lsi.into < lsj.into
	}
	return lsi.from < lsj.from
}
*/

type LinkHeap struct {
	links  []Link
	sumSet bool
	sumA   uint32
	sumB   uint32
}

func (lh *LinkHeap) Len() int {
	return len(lh.links)
}
func (lh *LinkHeap) Less(i, j int) bool {
	if (lh.sumSet) {
		optA := int64(lh.sumA)
		optB := int64(lh.sumB)
		li := &(lh.links[i])
		lj := &(lh.links[j])
		lia := int64(li.a)
		lib := int64(li.b)
		lja := int64(lj.a)
		ljb := int64(lj.b)
		return lia * optB + ljb * optA > lja * optB + lib*optA
	}
	return lh.links[i].ratio > lh.links[j].ratio
}
func (lh *LinkHeap) Swap(i, j int) {
	lh.links[i], lh.links[j] = lh.links[j], lh.links[i]
}
func (lh *LinkHeap) Push(x interface{}) {
	lh.links = append(lh.links, x.(Link))
}
func (lh *LinkHeap) Pop() interface{} {
	n := len(lh.links)
	x := lh.links[n - 1]
	lh.links = lh.links[0 : n - 1]
	return x
}

type In struct {
	n       int
	m       int
	/*
	fingers []int
	*/
	links   []Link
}

type Out struct {
	sumA uint32
	sumB uint32
}

func solve(in In) Out {
	uf := UnionFindNew(in.n)
	linkHeap := &LinkHeap{make([]Link, 0, len(in.links)), false, 0, 0}
	heap.Init(linkHeap)
	for _, l := range in.links {
	  heap.Push(linkHeap, l)
	}

	sumA, sumB := uint32(0), uint32(0)
	for uf.Count() > 1 && linkHeap.Len() > 0 {
		/*
		fmt.Println(fmt.Sprintf("NOSUM iter %v", linkHeap.links))
		*/
		link := heap.Pop(linkHeap).(Link)
		if (uf.Connected(link.from, link.into)) {
			/*
			fmt.Println(fmt.Sprintf("skipping %d -> %d with %d / %d", link.from, link.into, link.a, link.b))
			*/
			continue
		}
		/*
		fmt.Println(fmt.Sprintf("ADDING %d -> %d with %d / %d", link.from, link.into, link.a, link.b))
		*/
		uf.Union(link.from, link.into)
		sumA += uint32(link.a)
		sumB += uint32(link.b)
	}

	/*
	fmt.Println(fmt.Sprintf("target function INIT to %d/%d=%f",
		sumA,
		sumB,
		float64(sumA)/float64(sumB)))
	*/

	for true {
		uf.Reset()
		linkHeap.sumSet, linkHeap.sumA, linkHeap.sumB, linkHeap.links = true, sumA, sumB, linkHeap.links[:0]
		heap.Init(linkHeap)
		for _, l := range in.links {
		  heap.Push(linkHeap, l)
		}

		sumA, sumB = 0, 0
		for uf.Count() > 1 && linkHeap.Len() > 0 {
			/*
			fmt.Println(fmt.Sprintf("SUM iter %v", linkHeap.links))
			*/
			link := heap.Pop(linkHeap).(Link)
			if (uf.Connected(link.from, link.into)) {
				/*
				fmt.Println(fmt.Sprintf("skipping %d -> %d with %d / %d", link.from, link.into, link.a, link.b))
				*/
				continue
			}
			/*
			fmt.Println(fmt.Sprintf("ADDING %d -> %d with %d / %d", link.from, link.into, link.a, link.b))
			*/
			uf.Union(link.from, link.into)
			sumA += uint32(link.a)
			sumB += uint32(link.b)
		}

		if (uint64(sumA) * uint64(linkHeap.sumB) < uint64(linkHeap.sumA) * uint64(sumB)) {
			panic(
				fmt.Sprintf("target function went DOWN from %d/%d=%f to %d/%d=%f",
					linkHeap.sumA,
					linkHeap.sumB,
					float64(linkHeap.sumA)/float64(linkHeap.sumB),
					sumA,
					sumB,
					float64(sumA)/float64(sumB) ))
		}
		if (sumA == linkHeap.sumA && sumB == linkHeap.sumB) {
			break;
		} else {
			/*
			fmt.Println(fmt.Sprintf("target function went UP from %d/%d=%f to %d/%d=%f",
				sumA,
				sumB,
				float64(sumA)/float64(sumB),
				linkHeap.sumA,
				linkHeap.sumB,
				float64(linkHeap.sumA)/float64(linkHeap.sumB)))
			*/
		}
	}

	return Out{sumA, sumB}
}

/*
func fingerLinks(n int, links []Link) ([]Link, []int)  {
	sortedLinks := make([]Link, len(links), len(links))
	for idx, link := range links {
		sortedLinks[idx] = link
	}
	sort.Sort(LinkSort(sortedLinks))

	resLinks := append(make([]Link, 0, len(links)), links[0])
	fingers := make([]int, n + 1)
	for i := range sortedLinks {
		if (i == 0) {
			continue
		}
		lp, lc := sortedLinks[i - 1], sortedLinks[i]
		if (lc.from == lc.into) {
			continue	//	do not add self-loops
		}
		if (lc.from == lp.from) {
			if (lc.into == lp.into) {
				if (lc.b == lp.b && lc.a < lp.a) {
					continue	//	do not add suboptimal multi-edge if better is present
				}
				if (lc.a == lp.a && lc.b > lp.b) {
					continue	//	do not add suboptimal multi-edge if better is present
				}
			}
			resLinks = append(resLinks, lc)
		} else {
			fingers[lc.from] = len(resLinks)
			resLinks = append(resLinks, lc)
		}
	}
	fingers[n] = len(resLinks)

	return resLinks, fingers
}
*/

func gcd(x, y uint32) uint32 {
	for y != 0 {
		x, y = y, x % y
	}
	return x
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

	n := int(ReadInt32(scanner))
	m := int(ReadInt32(scanner))
	links := make([]Link, m)
	for i := range links {
		from := uint32(ReadInt32(scanner))
		into := uint32(ReadInt32(scanner))
		a := ReadInt8(scanner)
		b := ReadInt8(scanner)
		ratio := float64(a) / float64(b)
		links[i] = Link{from, into, a, b, ratio }
	}
	/*
	cleanLinks, fingers := fingerLinks(n, links)
	*/
	in := In{n, m, /*fingers, cleanLinks*/ links}

	out := solve(in)

	sumGCD := gcd(out.sumA, out.sumB)
	Writef(writer, "%d/%d\n", out.sumA / sumGCD, out.sumB / sumGCD)
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

//	http://stackoverflow.com/questions/30957644
//	http://www.math.tau.ac.il/~michas/ufind.pdf
//	http://algs4.cs.princeton.edu/15uf/WeightedQuickUnionPathCompressionUF.java.html
type UnionFind struct {
	parent []uint32
	size   []uint32
	count  uint32
}

func UnionFindNew(count int) *UnionFind {
	parent := make([]uint32, count, count)
	size := make([]uint32, count, count)
	lim := uint32(count)
	for i := uint32(0); i < lim; i++ {
		parent[i], size[i] = i, 1
	}
	return &UnionFind{parent, size, lim}
}

func (ufp *UnionFind) Reset() {
	lim := uint32(len(ufp.parent))
	for i := uint32(0); i < lim; i++ {
		ufp.parent[i], ufp.size[i] = i, 1
	}
	ufp.count = lim
}

func (ufp *UnionFind) validate(p uint32) {
	n := uint32(len(ufp.parent))
	if (p >= n) {
		panic(fmt.Sprintf("index %d is not between 0 and %d", p, (n - 1)))
	}
}

func (ufp *UnionFind) Count() uint32 {
	return ufp.count;
}

func (ufp *UnionFind) Find(p uint32) uint32 {
	ufp.validate(p)

	root := p
	for root != ufp.parent[root] {
		root = ufp.parent[root]
	}
	for p != root {
		newp := ufp.parent[p]
		ufp.parent[p] = root
		p = newp
	}

	return root;
}

func (ufp *UnionFind) Connected(p, q uint32) bool {
	return ufp.Find(p) == ufp.Find(q);
}

func (ufp *UnionFind) Union(p, q uint32) {
	rootP := ufp.Find(p)
	rootQ := ufp.Find(q)

	if (rootP == rootQ) {
		return
	}

	if (ufp.size[rootP] < ufp.size[rootQ]) {
		ufp.parent[rootP] = rootQ
		ufp.size[rootQ] += ufp.size[rootP]
	} else {
		ufp.parent[rootQ] = rootP
		ufp.size[rootP] += ufp.size[rootQ]
	}

	ufp.count -= 1;
}
