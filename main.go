package main

import (
	"fmt"
	"log"
	"net/http"
//	"time"
)

/* understand the type Server struct */
/*

type Server struct {
	Addr	string
	Handler Handler
	ReadTimeout	time.Duration
	WriteTimeout	time.Duration
	MaxHeaderBytes	int
	TLSConfig	*tls.Config
	TLSNextProto	map[string]func(*Server, *tls.Conn, Handler)
	ConnState	func(net.Conn, ConnState)
	ErrorLog	*log.Logger
}

*/

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
