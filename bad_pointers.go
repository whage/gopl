package main

import "fmt"

func main() {
	a := 10
	var x *int
	x = &a
	if x+1 != nil {
		fmt.Println("not equal to nil")
	}
}
