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
