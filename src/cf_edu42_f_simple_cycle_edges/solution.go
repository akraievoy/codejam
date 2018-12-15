package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

const (
	HOUSTON = false
)

type caseInput struct {
	n, m uint32
	ac   *EdgeList
}

func readCaseInput(scanner *bufio.Scanner) caseInput {
	n := uint32(readInt(scanner))
	m := uint32(readInt(scanner))

	el := EdgeListNew(n + 1)
	for label := uint32(1); label <= m; label++ {
		source := uint32(readInt(scanner))
		target := uint32(readInt(scanner))
		el.AddEdge(source, target, label)
	}

	return caseInput{n, m, el}
}

type caseOutput struct {
	simpleCycleLabels []uint32
}

func writeCaseOutput(writer *bufio.Writer, out caseOutput) {
	writef(writer, "%d\n", len(out.simpleCycleLabels))
	for _, l := range out.simpleCycleLabels {
		writef(writer, "%d ", l)
	}
}

//noinspection GoBoolExpressions
func solveCase(in caseInput) caseOutput {
	cycleLabelUF := UnionFindNew(in.m + 1) //	which cycles intersect

	dfsLevels := make([]uint32, in.n+1)
	cycleTops := make([]uint32, in.m+1)
	dfsEdges := make([]*Edge, in.n+1)

	vertexQ := make([]*Edge, 0, in.n)
	cycleEdges := make([]*Edge, 0, in.m)
	dfsStart := uint32(1)
	for dfsStart <= in.n {
		if dfsEdges[dfsStart] != nil {
			dfsStart += 1
			continue
		}

		vertexQ = append(vertexQ[:0], &Edge{0, dfsStart, 0, 0, nil})
		cycleEdges = cycleEdges[:0]
		dfsLevels[dfsStart] = 1

		for len(vertexQ) > 0 {
			edgeToSource := vertexQ[len(vertexQ)-1]
			source := edgeToSource.Target
			sourceLevel := dfsLevels[source]
			vertexQ = vertexQ[:len(vertexQ)-1]

			if HOUSTON {
				fmt.Printf("edgesFor %d (%d)\n", source, len(in.ac.EdgesFor(source)))
			}
			for _, e := range in.ac.EdgesFor(source) {

				if e.Target == edgeToSource.Source {
					if HOUSTON {
						fmt.Printf("\tBACKSTOP\t%d -> %d\t%d\t%d\n", e.Source, e.Target, e.Label, e.CycleLabel)
					}
					continue // same link we used for DFS, but seen from the other side
				}

				if dfsLevels[e.Target] == 0 {
					if HOUSTON {
						fmt.Printf("\tPROPAGATE\t%d -> %d\t%d\t%d\n", e.Source, e.Target, e.Label, e.CycleLabel)
					}
					//	add all new vertexes to spanning tree
					vertexQ = append(vertexQ, e)
					dfsEdges[e.Target] = e
					dfsLevels[e.Target] = sourceLevel + 1
				} else {
					if e.Source < e.Target {
						if HOUSTON {
							fmt.Printf("\tCYCLE\t%d -> %d\t%d\t%d\n", e.Source, e.Target, e.Label, e.CycleLabel)
						}
						//	this edge forms a cycle
						cycleEdges = append(cycleEdges, e)
					} else {
						//	do nothing, we'll see this edge from the other side
						if HOUSTON {
							fmt.Printf("\tCYCLE DUPE\t%d -> %d\t%d\t%d\n", e.Source, e.Target, e.Label, e.CycleLabel)
						}
					}
				}
			}
		}

		if HOUSTON {
			fmt.Printf("dfsLevels = %v\n", dfsLevels)
			fmt.Printf("cycleEdges (%d)\n", len(cycleEdges))
			for _, e := range cycleEdges {
				fmt.Printf("%d -> %d\t%d\t%d\n", e.Source, e.Target, e.Label, e.CycleLabel)
			}
			fmt.Printf("dfsEdges\n")
			for i, e := range dfsEdges {
				if e != nil {
					fmt.Printf("\t%d %d -> %d\t%d\t%d\n", i, e.Source, e.Target, e.Label, e.CycleLabel)
				} else {
					fmt.Printf("\t%d NIL\n", i)
				}
			}
		}

		for _, ce := range cycleEdges {
			cycleLabel := ce.Label
			ce.CycleLabel = cycleLabel
			ce.Reverse.CycleLabel = cycleLabel
			s := ce.Source
			t := ce.Target
			if HOUSTON {
				fmt.Printf(
					"cycleEdge\t%d -> %d\t%d\t%d\n",
					ce.Source, ce.Target, ce.Label, ce.CycleLabel,
				)
			}

			for s != t {
				if dfsLevels[s] >= dfsLevels[t] {
					if HOUSTON {
						fmt.Printf(
							"\tdfsEdges[s=%d at level %d]\t%d -> %d\t%d\t%d\n",
							s, dfsLevels[s], dfsEdges[s].Source, dfsEdges[s].Target, dfsEdges[s].Label, dfsEdges[s].CycleLabel,
						)
					}
					if dfsEdges[s].CycleLabel == cycleLabel {
						panic(fmt.Sprintf("seen same cycle with label %d", cycleLabel))
					} else if dfsEdges[s].CycleLabel > 0 {
						if HOUSTON {
							fmt.Printf("\t\tUNION %d with %d\n", cycleLabel, dfsEdges[s].CycleLabel)
							fmt.Printf("\t\ts <== %d\n", cycleTops[dfsEdges[s].CycleLabel])
						}
						cycleLabelUF.Union(dfsEdges[s].CycleLabel, cycleLabel)
						s = cycleTops[dfsEdges[s].CycleLabel]
					} else {
						if HOUSTON {
							fmt.Printf("\t\tLABEL with %d\n", cycleLabel)
							fmt.Printf("\t\ts <== %d\n", dfsEdges[s].Source)
						}
						dfsEdges[s].CycleLabel = cycleLabel
						dfsEdges[s].Reverse.CycleLabel = cycleLabel
						s = dfsEdges[s].Source
					}
				} else {
					if HOUSTON {
						fmt.Printf(
							"\tdfsEdges[t=%d at level %d]\t%d -> %d\t%d\t%d\n",
							t, dfsLevels[t], dfsEdges[t].Source, dfsEdges[t].Target, dfsEdges[t].Label, dfsEdges[t].CycleLabel,
						)
					}
					if dfsEdges[t].CycleLabel == cycleLabel {
						panic(fmt.Sprintf("seen same cycle with label %d", cycleLabel))
					} else if dfsEdges[t].CycleLabel > 0 {
						if HOUSTON {
							fmt.Printf("\t\tUNION %d with %d\n", cycleLabel, dfsEdges[t].CycleLabel)
							fmt.Printf("\t\tt = %d\n", cycleTops[dfsEdges[t].CycleLabel])
						}
						cycleLabelUF.Union(dfsEdges[t].CycleLabel, cycleLabel)
						t = cycleTops[dfsEdges[t].CycleLabel]
					} else {
						if HOUSTON {
							fmt.Printf("\t\tLABEL with %d\n", cycleLabel)
							fmt.Printf("\t\tt = %d\n", dfsEdges[t].Source)
						}
						dfsEdges[t].CycleLabel = cycleLabel
						dfsEdges[t].Reverse.CycleLabel = cycleLabel
						t = dfsEdges[t].Source
					}
				}
			}
			cycleTops[cycleLabel] = s
		}

		dfsStart += 1
	}

	simpleCycleLabels := make([]uint32, 0, in.m)

	if HOUSTON {
		fmt.Printf("Final edge filter\n")
	}
	for _, e := range in.ac.Added {
		if HOUSTON {
			fmt.Printf(
				"\t%d -> %d\t%d\t%d\n",
				e.Source, e.Target, e.Label, e.CycleLabel,
			)
		}

		if e.CycleLabel > 0 {
			ufSize := cycleLabelUF.Size(e.CycleLabel)
			if HOUSTON {
				fmt.Printf(
					"\t\tnumber of intersecting cycles: %d\n",
					ufSize,
				)
			}
			if ufSize == 1 {
				simpleCycleLabels = append(simpleCycleLabels, e.Label)
			}
		}
	}

	// sort.Sort(Uint32Sort(simpleCycleLabels))

	return caseOutput{simpleCycleLabels}
}

//type Uint32Sort []uint32
//
//func (s Uint32Sort) Len() int           { return len(s) }
//func (s Uint32Sort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
//func (s Uint32Sort) Less(i, j int) bool { return s[i] < s[j] }

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
	scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	writeCaseOutput(writer, solveCase(readCaseInput(scanner)))
}

func readInt64(sc *bufio.Scanner) int64 {
	if !sc.Scan() {
		panic("failed to scan next token")
	}

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

type Edge struct {
	Source, Target uint32
	Label          uint32
	CycleLabel     uint32
	Reverse        *Edge
}

type EdgeList struct {
	Added []*Edge
	Edges [][]*Edge
}

func EdgeListNew(
	numNodes uint32,
) *EdgeList {
	edges := make([][]*Edge, numNodes+1, numNodes+1)
	added := make([]*Edge, 0, numNodes+1)
	for i := range edges {
		edges[i] = make([]*Edge, 0, 32)
	}
	return &EdgeList{
		added,
		edges,
	}
}

func (el *EdgeList) AddEdge(source, target uint32, label uint32) {
	added := &Edge{source, target, label, 0, nil}
	reverse := &Edge{target, source, label, 0, nil}

	added.Reverse = reverse
	reverse.Reverse = added

	el.Added = append(el.Added, added)
	el.Edges[source] = append(el.Edges[source], added)
	el.Edges[target] = append(el.Edges[target], reverse)
}

func (el *EdgeList) EdgesFor(source uint32) []*Edge {
	return el.Edges[source]
}

type UnionFind struct {
	parent []uint32
	size   []uint32
	count  uint32
}

func UnionFindNew(count uint32) *UnionFind {
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
