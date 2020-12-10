package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Home"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is About"))
}
