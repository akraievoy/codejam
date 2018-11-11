package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type In struct {
	n, m  uint32
	flags [][]bool
}

type Out struct {
	largest uint32
}

func solve(in In) Out {
	uf := UnionFindNew(int(in.n * in.m))
	largest := uint32(0)
	for i := uint32(0); i < in.n; i++ {
		for j := uint32(0); j < in.m; j++ {
			if !in.flags[i][j] {
				continue
			}
			idx := i * in.m + j
			if i > 0 {
				if j > 0 &&  in.flags[i - 1][j - 1] {
					uf.Union(idx, idx - in.m - 1)
				}
				if in.flags[i - 1][j] {
					uf.Union(idx, idx - in.m)
				}
				if j + 1 < in.m && in.flags[i - 1][j + 1] {
					uf.Union(idx, idx - in.m + 1)
				}
			}
			if j > 0 &&  in.flags[i][j - 1] {
				uf.Union(idx, idx - 1)
			}
			maybeLargest := uf.Size(idx)
			if (maybeLargest > largest) {
				largest = maybeLargest
			}
		}
	}

	return Out{largest}
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

	n := ReadInt32(scanner)
	m := ReadInt32(scanner)
	flags := make([][]bool, n)
	for i := range flags {
		flags[i] = make([]bool, m)
		for j := range flags[i] {
			flags[i][j] = ReadInt32(scanner) > 0
		}
	}
	in := In{uint32(n), uint32(m), flags}

	out := solve(in)

	Writef(writer, "%d\n", out.largest)
}

//	boring IO

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

func (ufp *UnionFind) Size(p uint32) uint32 {
	return ufp.size[ufp.Find(p)]
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
