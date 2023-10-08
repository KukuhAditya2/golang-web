package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Response http code
func ResponseCode(writter http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writter.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writter, "name is empty")
	} else {
		fmt.Fprintf(writter, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:1919?name=kukuh", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
