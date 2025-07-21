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
	fmt.Fprintf(w, "Teste A -> B!")
	fmt.Fprintf(w, "Teste B -> A!")
	fmt.Fprintf(w, "Teste B -> A!")
}

func main() {
	fmt.Println("Teste modificação! AGAIN")
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
