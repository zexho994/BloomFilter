package BloomFilter

type BloomFilter struct {
	m  uint // the numbers of val
	k  uint // the maximum tolerable error
	bs *BitSet
}

func NewBloomFilter(m, k uint) BloomFilter {
	return BloomFilter{m: max(1, m), k: max(1, k), bs: New(max(1, m))}
}

func EstimateParameters(n uint, p float64) (m, k uint) {
	m = uint(math.Ceil(-1 * float64(n) * math.Log(p) / math.Pow(math.Ln2, 2)))
	k = uint(math.Ceil(math.Ln2 * float64(m) / float64(n)))
	return
}

func (bf *BloomFilter) add(s string) {

}

func baseHashes(data []byte) [4]uint64 {
	var d digest128
	hash1, hash2, hash3, hash4 := d.sum256(data)
	return [4]uint64{
		hash1, hash2, hash3, hash4,
	}
}

func (bf *BloomFilter) exist(s string) bool {
	return false
}

func max(n1, n2 uint) uint {
	if n1 > n2 {
		return n1
	}
	return n2
}
