package BloomFilter

type BloomFilter struct {
	m  uint // the numbers of val
	k  uint // the maximum tolerable error
	bs *BitSet
}

func NewBloomFilter(m, k uint) BloomFilter {
	return BloomFilter{m: max(1, m), k: max(1, k), bs: New(max(1, m))}
}

func (bf *BloomFilter) add(s string) {

}

func (bf *BloomFilter) exist(s string) bool {
	return false
}
