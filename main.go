package main

import (
	"fmt"
	"log"
	"net/http"
)

func gwpHandler(msg string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, msg)
	})
}

func shortHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact me at ardeshir.org at gmail.com")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", gwpHandler("Hi from gwp! Try /about"))
	mux.Handle("/about", gwpHandler("About gwp.."))
	mux.HandleFunc("/contact", shortHandler)

	log.Println("We're up on 9090...")
	http.ListenAndServe(":9090", mux)
}
