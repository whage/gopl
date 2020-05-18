package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Man struct {
	name string
}

func (t Man) Speak() string {
	return "asdf"
}

func main() {
	var a Speaker = Man{"Tod"}
	s := a.(Speaker) // try to type assert to interface   
	fmt.Print(s)
}
