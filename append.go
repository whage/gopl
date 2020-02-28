package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	fmt.Println("zlen", zlen)
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2*len(x)
		}
		fmt.Println("zcap", zcap)
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	fmt.Println("z inside appendInt", z)
	return z

}

func main() {
	//x := []int{}
	var x []int
	x = appendInt(x, 123)
	x = appendInt(x, 99)
	fmt.Println(x)
}
