package main

import (
	"fmt"
)

func rotina(ch chan int) {
	ch <- 1
//	ch <- 2
}

func main() {
	ch := make(chan int, 1)

	go rotina(ch) //se todas as goroutines dormirem -> deadlock
	fmt.Println(<-ch) //remove do canal -> bloquante //1
	fmt.Println(<-ch)
	fmt.Print("Terminou")
}