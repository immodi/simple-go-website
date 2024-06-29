package main

import (
	"fmt"
	"immmodi/website/helpers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Currently listening on http://127.0.0.1:8000/")

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", helpers.HomeHandler)
	http.HandleFunc("/play", helpers.PlayHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
