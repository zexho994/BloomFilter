package BloomFilter

type BloomFilter struct {
	m  uint // the numbers of val
	k  uint // the maximum tolerable error
	bs *BitSet
}

func NewBloomFilter() BloomFilter {
	return BloomFilter{}
}

func (bf *BloomFilter) add(s string) {

}

func (bf *BloomFilter) exist(s string) bool {
	return false
}
