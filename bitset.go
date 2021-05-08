package BloomFilter

type BitSet struct {
	data   []uint64
	length uint
}

// From is a constructor used to create a BitSet from an array of integers
func From(buf []uint64) *BitSet {
	return &BitSet{data: buf, length: uint(len(buf)) * 64}
}
