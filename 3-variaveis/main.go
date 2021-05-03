package main

import "fmt"

func main() {

	var num int // uint8 uint16 uint32 uint64 int8 int16 int32 int64
	num = 10
	fmt.Println(num)

	var altura float32 //float64 complex64 complex128
	altura = 1.77
	fmt.Println(altura)

	var nome string
	nome = "wesley"
	fmt.Println(nome)

	var pointer *string
	pointer = &nome
	fmt.Println(pointer)

	foo, bar := "foo", "bar"
	fmt.Println(foo, bar)
}
