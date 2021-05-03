package main

import "fmt"

//Pessoa struct
type Pessoa struct {
	Nome string
	Idade int
}

//Filho struct
type Filho struct {
	Pessoa
	NomePai string
}

func (p Pessoa) andar() {
	fmt.Printf("%s está andando", p.Nome)
}

func (p *Pessoa) setIdade(idade int) {
	p.Idade = idade
}

func main() {
	pessoa := Pessoa {
		Nome: "Wesley",Idade: 25,
	}
	fmt.Printf("Nome: %s Idade: %d\n", pessoa.Nome, pessoa.Idade)

	pessoa.setIdade(30)

	fmt.Printf("Idade: %d\n", pessoa.Idade) //Idade: 30

	filho := Filho {
		NomePai: "Jose",
	}
	filho.Nome = "Wesley"
	filho.Idade = 26
	filho.andar()
}