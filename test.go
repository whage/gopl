package main

import "fmt"

func main() {
	m := make(map[string]int)
	x, ok := m["alma"]
	if !ok {
		fmt.Println("key does not exist")
	}
	fmt.Println(x)
}
