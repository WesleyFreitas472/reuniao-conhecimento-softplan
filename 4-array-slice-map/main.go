package main

import "fmt"

func main() {

	//Array
	var a [5]int
	a[0] = 20
	fmt.Println(a[0]) // 20

	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println(b) // [0 1 2 3 4]

	//slice - é um pedaço de um array
	nomes := []string{"Wesley", "Freitas"}
	fmt.Println(nomes) //[Wesley Freitas]

	c := b[0:2]
	fmt.Println(c) // [0 1]
	
	c[0] = 10
	fmt.Println(b) //[10 1 2 3 4]

	c = b[2:]
	fmt.Println(c) //[2 3 4]

	c = b[:3]
	fmt.Println(c) //[10 1 2]

	slice := make([]int, 2) //cria array interno de string, tamanho do slice e capacidade do array interno
	fmt.Println(slice, len(slice), cap(slice)) //[         ] 2, 4

	slice = append(slice, 0, 1)
	fmt.Println(slice, len(slice), cap(slice)) //[0 0 0 1] 4 4

	slice = append(slice, 3)
	fmt.Println(slice, len(slice), cap(slice)) //[0 0 0 1 3] 5 8

	about := map[string]string{
		"Nome": "Wesley",
		"Time": "Mercúrio - Devops",
	}
	about["Cargo"] = "Software Developer"
	fmt.Println(about) // map[Cargo:Software Developer Nome:Wesley Time:Mercúrio - Devops]

	fmt.Println("Meu nome é:", about["Nome"])
}
