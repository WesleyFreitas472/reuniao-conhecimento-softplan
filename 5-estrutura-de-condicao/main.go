package main

import "fmt"

func main() {
	idade := 18

	if idade >= 18 {
		fmt.Println("Pode dirigir")
	} else {
		fmt.Println("Não pode dirigir")
	}

	fruta := "Uva"
	switch fruta {
	case "Uva":
		fmt.Println("Minha fruta favorita é Uva")
	case "Morango":
		fmt.Println("Morango não é uma fruta")
	default:
		fmt.Println("A fruta digitada foi: ", fruta)
	}
}
