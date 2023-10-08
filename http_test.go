package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test Http
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:1212/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
