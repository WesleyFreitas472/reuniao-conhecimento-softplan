package main

import (
	"fmt"
	"time"
)

//Função retorna por meio de um canal
func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //envia dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	ch := make(chan int)

	go doisTresQuatroVezes(2, ch)

	a, b := <-ch, <-ch //recebendo o valor do canal - bloqueante
	fmt.Println(a, b)

	fmt.Println(<-ch)
}