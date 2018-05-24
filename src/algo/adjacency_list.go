package algo

type AdjacencyListEntry struct {
	Target, IndexAtTarget uint32
	Capacity, Flow        float64
}

type AdjacencyList [][]AdjacencyListEntry

func AdjacencyListNew(size uint32) AdjacencyList {
	list := make([][]AdjacencyListEntry, size)
	for from := range list {
		list[from] = make([]AdjacencyListEntry, 0, size)
	}
	return list
}

func (alp *AdjacencyList) Size() int {
	return len(*alp)
}

func (alp *AdjacencyList) AddEdge(s, t uint32, capacity float64) {
	al := *alp

	al[s] =
		append(
			al[s],
			AdjacencyListEntry{
				t,
				uint32(len(al[t])),
				capacity,
				0,
			},
		)
	al[t] =
		append(
			al[t],
			AdjacencyListEntry{
				s,
				uint32(len(al[s])-1),
				0,
				0,
			},
		)
}
