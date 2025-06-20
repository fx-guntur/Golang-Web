package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello redirect")
}

/*
kalo misal di url yang sama beda endpoint tinggal tulis endpointnya aja, kalo beda url tulis full seperti ini :
http.Redirect(w, r, "https://www.google.com", http.StatusTemporaryRedirect)
*/

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	// logic

	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(w http.ResponseWriter, r *http.Request) {
	// logic
	http.Redirect(w, r, "https://www.google.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
