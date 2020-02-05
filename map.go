package main

import "fmt"

func main() {
	var y *map[string]int
	x := make(map[string]int)
	y = &x
	fmt.Println(y)
}
