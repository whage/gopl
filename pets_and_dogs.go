package main

import (
	"fmt"
)

type Pet interface {
	Name() string
	Speak() string
	Play()
}

type pet struct {
	name string
	speaker func() string
}

type Dog interface {
	Pet
	Breed()
}

type dog struct {
	pet
	breed string
}

func NewPet(name string) *pet {
	p := &pet{
		name: name,
	}
	p.speaker = p.speak
	return p
}

func (p *pet) Play() {
	fmt.Println(p.Speak())
}

func (p *pet) Speak() string {
	return p.speaker()
}

func (p *pet) speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *pet) Name() string {
	return p.name
}

func NewDog(name, breed string) *dog {
	d := &dog{
		pet: pet{name:name},
		breed: breed,
	}
	d.speaker = d.speak
	return d
}

func (d *dog) speak() string {
	return fmt.Sprintf("%v and I'm a %v", d.pet.speak(), d.breed)
}

func Play(p Pet) {
	p.Play()
}

func main() {
	d := NewDog("Martin", "pointer")
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
	Play(d)
}
