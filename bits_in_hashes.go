package main

import(
	"fmt"
	"crypto/sha256"
	"github.com/whage/gopl/popcount"
)

func countDifferentBits(a, b [32]byte) int {
	sum := 0

	for i, _ := range a {
		sum += popcount.PopCount(uint64(a[i]) ^ uint64(b[i]))
	}

	return sum
}

func main() {
	a := sha256.Sum256([]byte("A"))
	b := sha256.Sum256([]byte("B"))
	fmt.Printf("Number of different bits: %d\n", countDifferentBits(a, b))
}
