package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(":8001", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	r.Header.Write(w)
}
