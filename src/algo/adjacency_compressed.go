package algo

import (
	"fmt"
	"sort"
)

/* BEWARE: this does not what you'd expect for the vertex with index zero */

type Edge struct {
	Source, Target uint32
	Label          uint32
}

type EdgeList struct {
	Edges             []Edge
	Directed          bool
	SelfLoopsAllowed  bool
	MultiEdgesAllowed bool
}

func EdgeListNew(
	numEdges uint32,
	directed bool,
	selfLoopsAllowed bool,
	MultiEdgesAllowed bool,
) *EdgeList {
	capacity := numEdges
	if !directed {
		capacity *= 2
	}
	list := make([]Edge, 0, capacity)
	return &EdgeList{
		list,
		directed,
		selfLoopsAllowed,
		MultiEdgesAllowed,
	}
}

func (el *EdgeList) AddEdge(source, target uint32, label uint32) {
	if source == target && !el.SelfLoopsAllowed {
		panic(
			fmt.Sprintf(
				"self-loop from %d to %d with label %v - not allowed", source, target, label,
			),
		)
	}

	el.Edges = append(el.Edges, Edge{source, target, label})
	if !el.Directed {
		el.Edges = append(el.Edges, Edge{target, source, label})
	}
}

type EdgeSort []Edge

func (es EdgeSort) Len() int      { return len(es) }
func (es EdgeSort) Swap(i, j int) { es[i], es[j] = es[j], es[i] }
func (es EdgeSort) Less(i, j int) bool {
	lsi, lsj := es[i], es[j]
	return lsi.Source == lsj.Source && lsi.Target < lsj.Target || lsi.Source < lsj.Source
}

func (el *EdgeList) Compress(numNodes uint32) *AdjacencyCompressed {
	sort.Sort(EdgeSort(el.Edges))
	if !el.MultiEdgesAllowed {
		var prev Edge
		for i, e := range el.Edges {
			if i > 0 && prev.Source == e.Source && prev.Target == e.Target {
				panic(
					fmt.Sprintf(
						"self-loop from %d to %d with labels %v and %v - not allowed",
						e.Source, e.Target, e.Label, prev.Label,
					),
				)
			}
			prev = e
		}
	}

	fingers := make([]uint32, numNodes+1)
	from := uint32(0)
	for edgeIdx, edge := range el.Edges {
		for from < edge.Source {
			from++
			fingers[from] = uint32(edgeIdx)
		}
	}
	for from < numNodes {
		from++
		fingers[from] = uint32(len(el.Edges))
	}

	return &AdjacencyCompressed{el.Edges, fingers}
}

type AdjacencyCompressed struct {
	Edges []Edge
	Fingers []uint32
}

func (ac *AdjacencyCompressed) EdgesFor(source uint32) []Edge {
	return ac.Edges[ac.Fingers[source]:ac.Fingers[source+1]]
}
