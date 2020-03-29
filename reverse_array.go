package main

import (
	"fmt"
)

func reverse (s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{1,2,3,4,5}
	
	// Reverse accepts a slice, not an array!
	// Note that even though `a` is an array, we have a reference to it
	// in `s` and so the function modifies the array!

	s := a[:]
	reverse(s)
	fmt.Println(a)
}
