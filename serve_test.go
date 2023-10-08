package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// Handler function
func TestServer(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	}
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Request test jika ingin melihat URL atau method pada server
func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	}
	server := http.Server{
		Addr:    "localhost:9898",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
