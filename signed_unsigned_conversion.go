package main

import "fmt"

func main() {
	var x uint8 = 255
	fmt.Println(int8(x)) // -1

	var y int8 = -128
	fmt.Println(uint8(y)) // 128
}
