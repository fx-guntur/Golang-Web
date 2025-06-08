package belajargolangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	/*
		StripPrefix returns a handler that serves HTTP requests by removing the given prefix from the request
		URL's Path (and RawPath if set) and invoking the handler h. StripPrefix handles a request for a path that
		doesn't begin with prefix by replying with an HTTP 404 not found error.
		The prefix must match exactly: if the prefix in the request contains escaped characters the reply is also
		an HTTP 404 not found error.
	*/

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
	kalau pake cara yang atas ketika kita upload file distrivution golang ke server \
	maka kita perlu upload folder resourcesnya juga, maka lebih baik pake golang embed
*/

//go:embed resources
var resources embed.FS

func TestFileServerGoEmbed(t *testing.T) {
	/*
		menggunakan fs sub karena ketika menggunakan embed yang tadinya jika pakai cara biasa urlnya berupa :
		resources/static/nama_file menjadi static/resources/nama_file
		maka dari itu menggunakan fs sub ketika diakses pada fileServer yang diakses langsung file nya tidak
		dari direktori resources sehingga urlnya akan jadi static/nama_file
	*/
	directory, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
