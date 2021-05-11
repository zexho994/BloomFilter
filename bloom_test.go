package BloomFilter

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBloomFilter_add(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)
	bf.addString("hello")
	bf.addString("zexho")
	bf.addString("lishan")
	bf.addString("caohen")
	bf.addString("liangzai")
	bf.add([]byte{byte(1)})
	bf.add([]byte{byte(2)})
	bf.add([]byte{byte(3)})
	bf.add([]byte{byte(4)})
	bf.add([]byte{byte(5)})

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

	b11 := bf.existBts([]byte{byte(1)})
	b12 := bf.existBts([]byte{byte(2)})
	b13 := bf.existBts([]byte{byte(3)})
	b14 := bf.existBts([]byte{byte(4)})
	b15 := bf.existBts([]byte{byte(5)})

	b16 := bf.existBts([]byte{byte(6)})
	b17 := bf.existBts([]byte{byte(7)})
	b18 := bf.existBts([]byte{byte(8)})
	b19 := bf.existBts([]byte{byte(9)})
	b20 := bf.existBts([]byte{byte(10)})

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

	isTure(b11)
	isTure(b12)
	isTure(b13)
	isTure(b14)
	isTure(b15)

	isTure(!b16)
	isTure(!b17)
	isTure(!b18)
	isTure(!b19)
	isTure(!b20)
}

func TestBloomFilter_StringExist(t *testing.T) {
	total := 1000000
	bf := NewBloomFilter(uint(total), 0.05)
	for i := 0; i < total; i++ {
		bf.add([]byte(strconv.Itoa(i)))
	}

	fail := 0
	for i := 0; i < total; i++ {
		if !bf.exist(strconv.Itoa(i)) {
			fail++
		}
	}
	fmt.Println("fail count = ", fail)

	fail = 0
	for i := total; i < total+10000; i++ {
		if bf.exist(strconv.Itoa(i)) {
			fail++
		}
	}
	fmt.Println("fail count = ", fail)
}
