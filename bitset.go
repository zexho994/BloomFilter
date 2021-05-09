package BloomFilter

// the wordSize of a bit set
const wordSize = uint(64)
const mask = wordSize - 1

// log2WordSize is log(wordSize) => log(64) = 6
const log2WordSize = uint(6)

// MaxCapacity Cap returns the total possible capacity, or number of bits
const MaxCapacity = ^uint(0)

type BitSet struct {
	set    []uint64
	length uint
}

// From is a constructor used to create a BitSet from an array of integers
func From(buf []uint64) *BitSet {
	return &BitSet{set: buf, length: uint(len(buf)) * 64}
}

func New(l uint) (bset *BitSet) {
	defer func() {
		if r := recover(); r != nil {
			bset = &BitSet{length: 0, set: make([]uint64, 0)}
		}
	}()
	bset = &BitSet{set: make([]uint64, l/64), length: l}
	return bset
}

// Set bit num to 1
// "num >> log2WordSize" is to find the Set index from Set
// "1 << (num & (wordSize - 1 ))" is to find the index from uint(64)
func (bs *BitSet) Set(num uint) {
	setIdx := num >> log2WordSize
	if num > MaxCapacity || setIdx >= bs.length {
		panic("You are exceeding the capacity")
	}
	bs.set[setIdx] |= 1 << (num & mask)
}

func (bs *BitSet) Get(num uint) bool {
	return bs.set[num>>log2WordSize]&(1<<(num&mask)) > 0
}
