package BloomFilter

import (
	"fmt"
	"testing"
)

func TestCap(t *testing.T) {
	u0 := ^uint(0)
	u1 := +uint(0)
	u2 := -uint(0)
	u3 := ^uint(1)
	fmt.Println(u0)
	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)
}

func TestWordSize(t *testing.T) {
	s := uint(64)
	fmt.Print(s)
}

func TestBitSet_Set(t *testing.T) {
	bset := New(640)
	bset.Set(64)
	bset.Set(131)
	//bset.Set(105)
	//bset.Set(132)
	//bset.Set(53)
	//bset.Set(159)

	b1 := bset.Get(64)
	b2 := bset.Get(131)
	b3 := bset.Get(132)
	b4 := bset.Get(63)

	isTure(b1)
	isTure(b2)
	isTure(!b3)
	isTure(!b4)
}
