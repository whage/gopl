package main

import (
	"fmt"
)

var count = 1

func fib(n int) int {
	count += 1
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(30))
	fmt.Println(count)
}
