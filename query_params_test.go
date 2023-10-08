package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(writter http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	lastname := request.URL.Query().Get("lastname")

	fmt.Fprintf(writter, "Hello my name is %s %s", name, lastname)

}

// http query params
func TestQueryParams(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1313/hello?name=kukuh", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}

func TestMultipleQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1313/hello?name=kukuh&lastname=aditya", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
