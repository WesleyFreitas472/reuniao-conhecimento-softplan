package main

import (
	"fmt"
	"time"
)

func printa(qtde int, msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {
	printa(10, "normal")
	go printa(10, "goroutine")
	time.Sleep(2 * time.Second) //goroutine morre após o fim da execução da rotina principal
	printa(10, "normal")
}