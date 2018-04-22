package algo

import "fmt"

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
	if p >= n {
		panic(fmt.Sprintf("index %d is not between 0 and %d", p, n - 1))
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
