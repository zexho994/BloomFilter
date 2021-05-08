package BloomFilter

import (
	"fmt"
	"testing"
)

func TestCap(t *testing.T) {
	u := Cap()
	fmt.Print(u)
}

func TestWordSize(t *testing.T) {
	s := uint(64)
	fmt.Print(s)
}
