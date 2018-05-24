package algo

type CapacityMatrix [][]float64

func CapacityMatrixNew(n uint32) CapacityMatrix {
	outerSlice := make([][]float64, n)
	for i := range outerSlice {
		outerSlice[i] = make([]float64, n)
	}
	return outerSlice
}

func (cm *CapacityMatrix) Size() int {
	return len(*cm)
}

func (cm *CapacityMatrix) AddLink(i, j uint32, capacity float64) {
	(*cm)[i][j] += capacity
	(*cm)[j][i] += capacity
}

func (cm *CapacityMatrix) Clone() CapacityMatrix {
	n := cm.Size()

	outerSlice := make([][]float64, n)

	for i := range outerSlice {
		outerSlice[i] = make([]float64, n)
		copy(outerSlice[i], (*cm)[i])
	}

	return outerSlice
}
