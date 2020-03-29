package main

import (
	"fmt"
)

type Breeder interface {
	Mate(other Breeder) Breeder
}

type Dog struct{}

func (d Dog) Mate(other Dog) Dog {
	return Dog{}
}

type Cat struct{}

func (c Cat) Mate(other Cat) Cat {
	return Cat{}
}

func main() {
	breeders := []Breeder{
		Dog{},
		Dog{},
		Cat{},
		Dog{},
	}

	fmt.Println(breeders)
}
