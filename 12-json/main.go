package main

import (
	"encoding/json"
	"fmt"
)

//Pessoa struct
type Pessoa struct {
	Nome string `json:"nome"`
	Idade int `json:"idade"`
}

func main() {
	p := Pessoa {
		Nome: "Wesley",
		Idade: 26,
	}

	pessoaBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Erro:", err.Error())
	}

	fmt.Println(pessoaBytes)
	fmt.Println(string(pessoaBytes))
}
