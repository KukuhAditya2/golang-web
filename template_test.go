package belajargolangweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// Template String
func SimpleTamplate(w http.ResponseWriter, r *http.Request) {
	tampateText := "<html><body>{{.}}</body></html>"
	tamplate := template.Must(template.New("SIMPLE").Parse(tampateText))

	tamplate.ExecuteTemplate(w, "SIMPLE", "Hello Kukuh Aditya Template")

}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost1212", nil)
	recoreder := httptest.NewRecorder()

	SimpleTamplate(recoreder, request)

	body, _ := io.ReadAll(recoreder.Result().Body)
	fmt.Println(string(body))

}

// Template File
func TemplateFile(w http.ResponseWriter, r *http.Request) {
	// ini hanya include data dengan nama file yang kita tentukan
	// t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	// ini include semua data di dalam folder directory
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	// jagan lupa untuk memasukan nama filenya di sini
	t.ExecuteTemplate(w, "simple.gohtml", "hello my name is")

}

// Testing Template File
func TestTemplateFile(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:1111", nil)
	recorder := httptest.NewRecorder()

	TemplateFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// Ingat golang embed tidak adak tulisan derectory saat ini langsung sebutkan directorinya

//go:embed templates/*.gohtml
var templates embed.FS

// Template Go-Lang Embed
func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "hello ADITYA")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:2324", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
