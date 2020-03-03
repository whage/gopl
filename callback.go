package main

import (
	"fmt"
)

func f(cb func(a int) error) {
	fmt.Println(cb(2))
}

func alma(x int) error {
	return fmt.Errorf("hello")
}

func main() {
	f(alma)
}
