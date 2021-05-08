package BloomFilter

// the wordSize of a bit set
const wordSize = uint(64)

// log2WordSize is lg(wordSize)
const log2WordSize = uint(6)

type BitSet struct {
	data   []uint64
	length uint
}

// From is a constructor used to create a BitSet from an array of integers
func From(buf []uint64) *BitSet {
	return &BitSet{data: buf, length: uint(len(buf)) * 64}
}

func New(l uint) *BitSet {
	return &BitSet{data: make([]uint64, l/65), length: l}
}

// wordsNeeded calculates the number of words needed for i bits
func wordsNeeded(i uint) int {
	if i > (Cap() - wordSize + 1) {
		return int(Cap() >> log2WordSize)
	}
	return int((i + (wordSize - 1)) >> log2WordSize)
}

// Cap returns the total possible capacity, or number of bits
func Cap() uint {
	return ^uint(0)
}
