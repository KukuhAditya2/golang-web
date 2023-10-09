package belajargolangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// http.FileServer untuk membaca file code jenis lain
func TestFileServer(t *testing.T) {
	deroctory := http.Dir("./resource")
	fileServer := http.FileServer(deroctory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:3434",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// http static embed

// go:embed resource
var resource embed.FS

func TestFileServerEmbed(t *testing.T) {
	deroctory, _ := fs.Sub(resource, "resoruce")
	fileServer := http.FileServer(http.FS(deroctory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:3434",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
