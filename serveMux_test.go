package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

// Handler path atau router
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hii")
	})

	server := http.Server{
		Addr:    "localhost:7070",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
