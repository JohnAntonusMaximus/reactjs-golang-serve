package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./build")))
	fmt.Printf("Serving on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
