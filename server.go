package main

import (
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	Lissajous(w)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}
