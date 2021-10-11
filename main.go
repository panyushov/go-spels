package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/goodbye", GoodbyeServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	if err != nil {
		panic(err)
	}
}

func GoodbyeServer(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Goodbye, %s!", r.URL.Path[1:])
	if err != nil {
		panic(err)
	}
}
