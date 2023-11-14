package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", helloFunc)

	fmt.Println("Running server ...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
