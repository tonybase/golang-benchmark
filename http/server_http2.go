package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr: ":8002", // Normally ":443"
	}
	http2.ConfigureServer(srv, &http2.Server{})
	http.HandleFunc("/", http2Handler)
	log.Fatal(srv.ListenAndServeTLS("pem/localhost.cert", "pem/localhost.key"))
}

func http2Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	r.Header.Write(w)
}
