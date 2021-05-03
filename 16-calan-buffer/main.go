package main

import (
	"fmt"
	"time"
)

func getCincoItens(ch chan int) {
	ch <- 1
	ch <- 2
	time.Sleep(time.Second * 2)
	ch <- 3
	ch <- 4
	time.Sleep(time.Second * 2)
	ch <- 5
}

func main() {
	ch := make(chan int, 2)
	go getCincoItens(ch)
	
	fmt.Print(<-ch)
	fmt.Println(<-ch)

	fmt.Print(<-ch)
	fmt.Println(<-ch)

	fmt.Println(<-ch)
}