package algo

type SubsetFlags []bool

func SubsetFlagsNew(size int) SubsetFlags {
	flags := make([]bool, size)
	return flags
}

func (sf SubsetFlags) Next() bool {
	rollPos := len(sf) - 1
	for rollPos >= 0 && sf[rollPos] {
		rollPos -= 1
	}
	if rollPos < 0 {
		return false
	}

	sf[rollPos] = true
	for pos := rollPos + 1; pos < len(sf); pos++ {
		sf[pos] = false
	}
	return true
}
