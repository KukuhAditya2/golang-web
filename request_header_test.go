package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Request Header
func RequestHeaders(writter http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writter, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5656", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeaders(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
