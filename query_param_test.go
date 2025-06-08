package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=Bobi", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
