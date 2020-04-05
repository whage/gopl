package main

import (
	"fmt"
	"os"
)

type Actor interface {
	Act()
}

type MaleActor struct {}

func (m MaleActor) Act() {
	fmt.Println("MaleActor.Act()")
}

type Clooney struct {
	MaleActor
}

func (c Clooney) Act() {
	fmt.Println("Clooney.Act()")
}

func main() {
	var c Clooney = Clooney{}
	c.Act()

	var a Actor
	if len(os.Args) <= 1 {
		a = Clooney{}
	} else {
		a = MaleActor{}
	}

	a.Act()
}
