package main

import "fmt"

//Stringer interface
type Stringer interface {
	String() string
}

//Pessoa struct
type Pessoa struct {
	Nome string
	Idade int
}

//Carro struct
type Carro struct {
	Modelo string
	Ano int
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Meu nome é %s e minha idade é %d", p.Nome, p.Idade)
}

func (c Carro) String() string {
	return fmt.Sprintf("O carro é um %s ano %d", c.Modelo, c.Ano)
}

func printa(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	pessoa := Pessoa{
		Nome: "Wesley",
		Idade: 25,
	}

	carro := Carro{
		Modelo: "Celta",
		Ano: 2010,
	}

	printa(pessoa)
	printa(carro)
}



