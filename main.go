package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, go"))
	})
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
