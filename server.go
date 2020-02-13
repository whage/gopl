package main

import (
	"log"
	"net/http"

	"github.com/whage/gopl/surface"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface.Draw(w)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
