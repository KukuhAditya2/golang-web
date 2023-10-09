package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// Template data menggunakan Map

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	// Menggunakan Map
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Map Tamplate",
		"Name":  "Kukuh",
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:1111", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

// Template data menggunakan Struct

type Page struct {
	Title string
	Name  string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	// Menggunakan Map
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Hello",
		Name:  "KUKUH",
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:1111", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
