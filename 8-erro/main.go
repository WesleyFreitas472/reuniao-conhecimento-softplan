package main

import (
	"errors"
	"fmt"
)

func connectWithDatabase(uri string) (bool, error) {
	return false, errors.New("falha na conexão com o banco de dados")
}

func main() {
	if connection, err := connectWithDatabase("xpto"); err != nil {
		panic(err)	
	} else {
		fmt.Println(connection)
	}
}
