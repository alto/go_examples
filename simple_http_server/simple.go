package main

import (
	"log"
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Dude")
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
