package main

import "fmt"

func rotina(ch chan int) {
	fmt.Println("Executou")
	ch <- 1 //insere no canal
	ch <- 2
	ch <- 3
	ch <- 4
}

func main() {
	ch := make(chan int, 4)

	rotina(ch) //se todas as goroutines dormirem -> deadlock

	fmt.Println(<-ch) //remove do canal -> bloquante
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) //se comentar não dá erro
}