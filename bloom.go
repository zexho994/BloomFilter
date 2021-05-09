package BloomFilter

import "math"

type BloomFilter struct {
	m  uint // count of val
	k  uint // count of hash
	bs *BitSet
}

func NewBloomFilter(m uint, fp float64) *BloomFilter {
	n, k := EstimateParameters(m, fp)
	return &BloomFilter{m: n, k: k, bs: New(n)}
}

// EstimateParameters estimates requirements for m and k.
// Based on https://www.cnblogs.com/allensun/archive/2011/02/16/1956532.html
// used with permission.
func EstimateParameters(n uint, p float64) (m, k uint) {
	m = uint(math.Ceil(-1 * float64(n) * math.Log(p) / math.Pow(math.Ln2, 2)))
	k = uint(math.Ceil(math.Ln2 * float64(m) / float64(n)))
	return
}

func (bf *BloomFilter) addString(s string) {
	bf.add([]byte(s))
}

func (bf *BloomFilter) add(s []byte) {
	h := baseHashes(s)
	for i := uint(0); i < bf.k; i++ {
		bf.bs.Set(bf.location(h, i))
	}
}

// location returns the ith hashed location using the four base hash values
func (bf *BloomFilter) location(h [4]uint64, i uint) uint {
	return uint(location(h, i) % uint64(bf.m))
}

// location returns the ith hashed location using the four base hash values
func location(h [4]uint64, i uint) uint64 {
	ii := uint64(i)
	return h[ii%2] + ii*h[2+(((ii+(ii%2))%4)/2)]
}

func baseHashes(data []byte) [4]uint64 {
	var d digest128
	hash1, hash2, hash3, hash4 := d.sum256(data)
	return [4]uint64{
		hash1, hash2, hash3, hash4,
	}
}

func (bf *BloomFilter) exist(s string) bool {
	bt := []byte(s)
	h := baseHashes(bt)
	for i := uint(0); i < bf.k; i++ {
		if !bf.bs.Get(bf.location(h, i)) {
			return false
		}
	}
	return true
}
