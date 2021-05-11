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

func New(l uint) (bs *BitSet) {
	defer func() {
		if r := recover(); r != nil {
			bs = &BitSet{length: 0, set: make([]uint64, 0)}
		}
	}()
	bs = &BitSet{set: make([]uint64, l/64), length: l}
	return bs
}

// Set bit num to 1
// "num >> log2WordSize" is to find the Set index from Set
// "1 << (num & (wordSize - 1 ))" is to find the index from uint(64)
func (bs *BitSet) Set(num uint) {
	bs.extendSetMaybe(num)
	bs.set[num>>log2WordSize] |= 1 << (num & mask)
}

func (bs *BitSet) Get(num uint) bool {
	return bs.set[num>>log2WordSize]&(1<<(num&mask)) > 0
}

// extendSetMaybe adds additional words to incorporate new bits if needed
func (bs *BitSet) extendSetMaybe(num uint) {
	if num < bs.length {
		return
	}
	if num >= MaxCapacity {
		panic("You are exceeding the capacity")
	}
	nSize := bs.wordsNeeded(num)
	if bs.set == nil {
		bs.set = make([]uint64, nSize)
	} else if cap(bs.set) >= nSize {
		bs.set = bs.set[:nSize] // fast resize
	} else if len(bs.set) < nSize {
		newSet := make([]uint64, nSize, 2*nSize) // increase capacity 2x
		copy(newSet, bs.set)                     // TODO Concurrency issues
		bs.set = newSet
	}
	bs.length = num + 1
}

// wordsNeeded calculates the number of words needed for num bits
func (bs *BitSet) wordsNeeded(num uint) int {
	if num > (MaxCapacity - wordSize) {
		return int(MaxCapacity >> log2WordSize)
	}
	return int((num + wordSize) >> log2WordSize)
}
