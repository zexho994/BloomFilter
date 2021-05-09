package BloomFilter

import "testing"

func TestBloomFilter_add(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)
	bf.addString("hello")
	bf.addString("zexho")
	bf.addString("lishan")
	bf.addString("caohen")
	bf.addString("liangzai")

	b1 := bf.exist("hello")
	b2 := bf.exist("zexho")
	b3 := bf.exist("lishan")
	b4 := bf.exist("caohen")
	b5 := bf.exist("liangzai")
	b6 := bf.exist("hell")
	b7 := bf.exist("zeho")
	b8 := bf.exist("lshan")
	b9 := bf.exist("liazai")
	b10 := bf.exist("caahen")

	isTure(b1)
	isTure(b2)
	isTure(b3)
	isTure(b4)
	isTure(b5)
	isTure(!b6)
	isTure(!b7)
	isTure(!b8)
	isTure(!b9)
	isTure(!b10)
}
