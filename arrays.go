package main

import (
	"fmt"
)

func giveMe(a [3]int) {
	a[0] = 99
}

var x [3]int = [3]int{1,2,3}

func main() {
	fmt.Println(x)
	giveMe(x)
	fmt.Println(x)
}
