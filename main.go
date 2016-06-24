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

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", gwpHandler("Hi from gwp! Try /about"))
	mux.Handle("/about", gwpHandler("About gwp.."))

	log.Println("We're up on 9090...")
	http.ListenAndServe(":9090", mux)
}
