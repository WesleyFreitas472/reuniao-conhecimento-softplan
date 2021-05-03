package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	numeros := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for indice, numero := range numeros {
		fmt.Printf("Indice %d valor %d\n", indice, numero)
	}

	filme := map[string]string{
		"Nome": "A Fantástica Fabrica de Chocolate",
		"Diretor": "Tim Burton",
	}

	for key, value := range filme {
		fmt.Printf("Key: %s Value: %s\n", key, value)	
	}
}
