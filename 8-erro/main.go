package main

import (
	"errors"
	"fmt"
)

func connectWithDatabase(uri string) (bool, error) {
	return false, errors.New("falha na conexão com o banco de dados")
}

func main() {
	connection, err := connectWithDatabase("xpto")
	if err != nil {
		panic(err)
	}
	fmt.Println(connection)
}