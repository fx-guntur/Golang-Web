package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// jika ingin ngecek user sudah login atau belum dll bisa ditaru sebelum handler di eksekusi
	fmt.Println("Before execute handler")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After execute handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprint(w, "Hello middleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo executed")
		fmt.Fprint(w, "Hello foo")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Panic executed")
		panic("Ups")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
