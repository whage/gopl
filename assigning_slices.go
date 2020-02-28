package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3,4,5}
	b := []int{9,8,7}
	c := b
	x := []int{}

	c[2] = 100
	
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("x", x)
}
