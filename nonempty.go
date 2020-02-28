package main

import (
	"fmt"
)

// Modifies the input string in place and returns
// a slice of the modified string
func nonEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	s := []string{"alma", "korte", "", "banan"}
	fmt.Println(nonEmpty(s))
	fmt.Println(s)
}
