package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func main() {
	port := "8080"

	http.HandleFunc("/", rootHandler)
	fmt.Println("Ouvindo 0.0.0.0:" + port)
	http.ListenAndServe(":"+port, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requisição recebida")
	response := map[string]string{"msg": "Hello World"}
	responseBytes, err := json.Marshal(response)
	if err == nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(responseBytes)
}