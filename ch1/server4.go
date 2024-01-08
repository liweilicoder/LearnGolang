package main

import (
	"log"
	"net/http"
)

func server4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8002", nil))
}
