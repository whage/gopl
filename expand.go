package main

import (
	"fmt"
)

func replacer(string) string {
	return "x"
}

func expand(s string, f func(string) string) string {
	// TODO	
}

func main() {
	fmt.Println(expand("this is a $foo message", replacer))
}
