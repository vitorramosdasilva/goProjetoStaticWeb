package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", serveHome)
	http.ListenAndServe(":8080", nil)

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
