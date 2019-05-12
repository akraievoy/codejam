package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(j Jam) {
	n := uint32(j.Int())
	zeroClusters, oneClusters := UnionFindNew(n), UnionFindNew(n)
	for l := uint32(1); l < n; l++ {
		x, y, c := uint32(j.Int())-1, uint32(j.Int()-1), j.Int() > 0
		if c {
			oneClusters.Union(x, y)
		} else {
			zeroClusters.Union(x, y)
		}
	}
	validPairs := int64(0)
	countedZeroComponents, countedOneComponents := make(map[uint32]bool, n), make(map[uint32]bool, n)
	for v := uint32(0); v < n; v++ {
		zeroClusterSize, oneClusterSize := zeroClusters.Size(v), oneClusters.Size(v)
		zeroClusterRoot, oneClusterRoot := zeroClusters.Find(v), oneClusters.Find(v)
		if zeroClusterSize > 1 && !countedZeroComponents[zeroClusterRoot] {
			validPairs += int64(zeroClusterSize) * int64(zeroClusterSize-1)
			countedZeroComponents[zeroClusterRoot] = true
		}
		if oneClusterSize > 1 && !countedOneComponents[oneClusterRoot] {
			validPairs += int64(oneClusterSize) * int64(oneClusterSize-1)
			countedOneComponents[oneClusterRoot] = true
		}
		if zeroClusterSize > 1 && oneClusterSize > 1 {
			validPairs += int64(zeroClusterSize-1) * int64(oneClusterSize-1)
		}
	}
	j.P("%d\n", validPairs)
}

func main() {
	jam, closeFunc := JamNew()
	defer closeFunc()
	solve(jam)
}

type Jam interface {
	Scanner() *bufio.Scanner
	Writer() *bufio.Writer
	Close()

	Str() string
	Int() int64
	Float() float64

	P(format string, values ...interface{})
	PF(format string, values ...interface{})
}

func JamNew() (Jam, func()) {
	if len(os.Args) > 1 {
		panic("running with input file path is not supported")
	}

	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	jam := &jam{scanner, writer}
	return jam, jam.Close
}

type jam struct {
	sc *bufio.Scanner
	wr *bufio.Writer
}

func (j *jam) Close() {
	if err := j.wr.Flush(); err != nil {
		panic(err)
	}
}

func (j *jam) Scanner() *bufio.Scanner {
	return j.sc
}

func (j *jam) Writer() *bufio.Writer {
	return j.wr
}

func (j *jam) Str() string {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	return j.sc.Text()
}

func (j *jam) Int() int64 {
	if !j.sc.Scan() {
		panic("failed to scan next token")
	}

	res, err := strconv.ParseInt(j.sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}

	return res
}

func (j *jam) Float() float64 {
	j.sc.Scan()
	res, err := strconv.ParseFloat(j.sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return res
}

func (j *jam) P(format string, values ...interface{}) {
	_, err := fmt.Fprintf(j.wr, format, values...)
	if err != nil {
		panic(err)
	}
}

func (j *jam) PF(format string, values ...interface{}) {
	_, err := fmt.Fprintf(j.wr, format, values...)
	if err != nil {
		panic(err)
	}
	if err = j.wr.Flush(); err != nil {
		panic(err)
	}
}

type UnionFind struct {
	parent []uint32
	size   []uint32
	count  uint32
}

func UnionFindNew(count uint32) *UnionFind {
	parent := make([]uint32, count, count)
	size := make([]uint32, count, count)
	lim := count
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
	if p >= n {
		panic(fmt.Sprintf("index %d is not between 0 and %d", p, n-1))
	}
}

func (ufp *UnionFind) Count() uint32 {
	return ufp.count
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

	return root
}

func (ufp *UnionFind) Size(p uint32) uint32 {
	return ufp.size[ufp.Find(p)]
}

func (ufp *UnionFind) Connected(p, q uint32) bool {
	return ufp.Find(p) == ufp.Find(q)
}

func (ufp *UnionFind) Union(p, q uint32) {
	rootP := ufp.Find(p)
	rootQ := ufp.Find(q)

	if rootP == rootQ {
		return
	}

	if ufp.size[rootP] < ufp.size[rootQ] {
		ufp.parent[rootP] = rootQ
		ufp.size[rootQ] += ufp.size[rootP]
	} else {
		ufp.parent[rootQ] = rootP
		ufp.size[rootP] += ufp.size[rootQ]
	}

	ufp.count -= 1
}
