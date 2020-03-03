package main

import (
	"fmt"
)

func main() {
	x := make([]int, 1, 5)
	x[0] = 9
	fmt.Println("x initially: %v", x)
	
	appendOne(x)
	fmt.Println("x after calling appendOne(): %v", x)
	fmt.Println("len(x): %d", len(x))

	rewriteFirst(x)
	fmt.Println("x after calling rewriteFirst(): %v", x)
	fmt.Println("len(x): %d", len(x))
}

func appendOne(s []int) {
	s = append(s, 1)
	fmt.Println("s after calling append(): %v", s)
	fmt.Println("len(s): %d", len(s))
}

func rewriteFirst(s []int) {
	s[0] = 99
	fmt.Println("s after rewriting first element: %v", s)
	fmt.Println("len(s): %d", len(s))
}
