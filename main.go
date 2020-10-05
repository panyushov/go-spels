package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/goodbye", GoodbyeServer)
	http.HandleFunc("/test", TestServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func GoodbyeServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye, %s!", r.URL.Path[1:])
}

func TestServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is test")
}
