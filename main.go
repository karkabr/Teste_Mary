package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World (pra não dá azar kkk)!")
	fmt.Fprintf(w, "Forçar conflito na main!")
	fmt.Fprintf(w, "Teste modificação!")
	fmt.Fprintf(w, "Teste sh!")
}

func main() {
	fmt.Println("Teste modificação! AGAIN")
	fmt.Println("Teste sh!")
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
