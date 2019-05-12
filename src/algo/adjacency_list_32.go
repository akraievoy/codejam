package algo

type AdjacencyListEntry32 struct {
	Target uint32
	One    bool
}

type AdjacencyList32 [][]AdjacencyListEntry32

func AdjacencyList32New(size uint32, avgDegree uint32) AdjacencyList32 {
	list := make([][]AdjacencyListEntry32, size)
	for from := range list {
		list[from] = make([]AdjacencyListEntry32, 0, avgDegree)
	}
	return list
}

func (alp *AdjacencyList32) Size() int {
	return len(*alp)
}

func (alp *AdjacencyList32) Links(from uint32) []AdjacencyListEntry32 {
	return (*alp)[from]
}

func (alp *AdjacencyList32) AddEdge(from, into uint32, one bool) {
	al := *alp

	al[from] =
		append(
			al[from],
			AdjacencyListEntry32{
				into,
				one,
			},
		)
	al[into] =
		append(
			al[into],
			AdjacencyListEntry32{
				from,
				one,
			},
		)
}
