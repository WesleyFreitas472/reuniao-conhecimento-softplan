package main

import "fmt"

func soma(num1, num2 int) int {
	return num1 + num2
}

func somatorio(numeros ...int) (resultado int) { //função variática
	for _, numero := range numeros {
		resultado = resultado + numero
	}
	return
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fibonacci(num - 1) + fibonacci((num - 2))
}

func apresentacao(nome string, idade int) {
	fmt.Printf("Meu nome é %s e minha idade é %d\n", nome, idade)
}

func alteraValor(num *int) {
	*num = 10
}

func idadeDoJose() {
	idade := 30
	fmt.Println(idade)
}

func getIdadeAndName() (int, string) {
	return 20, "Maria"
}

func main() { //Função principal
	fmt.Println("soma", soma(10, 20))
	fmt.Println("somatorio", somatorio(1, 2, 3, 4, 5))
	fmt.Println("fibonnaci de 5 é", fibonacci(6))
	
	idade := 26
	apresentacao("Wesley", idade)

	alteraValor(&idade)
	fmt.Println("Nova idade:", idade)
	
	idadeDoJose()
	fmt.Println("Idade na função main:", idade)

	fmt.Println(getIdadeAndName())
}