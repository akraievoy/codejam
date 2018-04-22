package algo

type Perm []uint32

func PermNew(size int) Perm {
	perm := make([]uint32, size)
	for i := range perm {
		perm[i] = uint32(i)
	}
	return perm
}
//	https://www.nayuki.io/page/next-lexicographical-permutation-algorithm

func (perm Perm) Next() bool {
	rollPos := len(perm) - 2
	for rollPos >= 0 && perm[rollPos] > perm[rollPos + 1] {
		rollPos -= 1
	}
	if rollPos < 0 {
		return false
	}

	swapPos := rollPos+1
	for pos := rollPos + 2; pos < len(perm); pos++ {
		if perm[rollPos] < perm[pos] && perm[pos] < perm[swapPos]{
			swapPos = pos
		}
	}
	perm[rollPos], perm[swapPos] = perm[swapPos], perm[rollPos]

	for i, j := rollPos+1, len(perm)-1; i < j; i, j = i+1, j-1 {
		perm[i], perm[j] = perm[j], perm[i]
	}

	return true
}

