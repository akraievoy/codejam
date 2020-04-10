package algo

// https://stackoverflow.com/a/11694507

type BitSetElem uint64
const BitSetElemBits = 64

// BitSet is a set of bits that can be set, cleared and queried.
type BitSet []BitSetElem

func NewBitSet(bitlen uint64) BitSet {
	return make([]BitSetElem, (bitlen + BitSetElemBits - 1)/BitSetElemBits)
}

// Set ensures that the given bit is set in the BitSet.
func (s *BitSet) Set(i uint64) {
	(*s)[i/BitSetElemBits] |= 1 << (i % BitSetElemBits)
}

// Clear ensures that the given bit is cleared (not set) in the BitSet.
func (s *BitSet) Clear(i uint64) {
	if len(*s) >= int(i/BitSetElemBits+1) {
		(*s)[i/BitSetElemBits] &^= 1 << (i % BitSetElemBits)
	}
}

// IsSet returns true if the given bit is set, false if it is cleared.
func (s *BitSet) IsSet(i uint64) bool {
	return (*s)[i/BitSetElemBits]&(1<<(i%BitSetElemBits)) != 0
}
