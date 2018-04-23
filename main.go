package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/test", Test)
	log.Fatal(http.ListenAndServe(":8888", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New route , %q", html.EscapeString(r.URL.Path))
}