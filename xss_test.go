package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template auto escape",
		"Body":  "<p>Ini adalah body</p>",
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateAutoDisabled(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template auto escape",
		"Body":  template.HTML("<p>Ini adalah body</p>"),
	})
}

func TestTemplateAutoDisabled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoDisabled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateAutoXSS(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Template auto escape",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	})
}

func TestTemplateAutoXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<script>alert('hacked')</script>", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
