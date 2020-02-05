package main

import "fmt"

func main() {
	n := 1

	switch n {
	case 1:
		fmt.Println("heads")
	case 0:
		fmt.Println("tails")
	default:
		fmt.Println("Landed on edge!")
	}
}
