package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func(writter http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writter, "Hello world")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	mux.HandleFunc("/about-us", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "about us")
	})

	/*
		kalo ditambahkan slash dibelakang maka ketika kita akses /images/bobi tidak akan notfound namun di proses
		di handler tersebut
	*/
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "images")
	})

	mux.HandleFunc("/images/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
