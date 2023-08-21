package main

import (
	"fmt"
	"net/http"
)

func say_hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server is runngin!")
}

func main() {
	http.HandleFunc("/", say_hello)
	http.ListenAndServe(":8080", nil)
}
