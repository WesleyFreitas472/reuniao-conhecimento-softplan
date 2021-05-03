package main

import "fmt"

//Ave interface
type Ave interface {
	voar()
}

//Aquatico interface
type Aquatico interface {
	nadar()
}

//Animal interface
type Animal interface {
	Ave
	Aquatico
}

//Pato struuct
type Pato struct {
}

//Nadar func
func (p Pato) nadar() {
	fmt.Println("Nadando...")
}

//Voar func
func (p Pato) voar() {
	fmt.Println("Voando...")
}

func main() {
	p := Pato{}
	voarComAnimal(p)
	voarComAve(p)
	nadarComAquatico(p)
}

func voarComAve(a Ave) {
	a.voar()
}

func voarComAnimal(a Animal) {
	a.voar()
}

func nadarComAquatico(a Aquatico) {
	a.nadar()
}