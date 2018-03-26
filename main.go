package main

import (
	"net/http"
)

func handler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message" : "hello world!"}`))
}

func handler_test (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message" : "test : hello world!"}`))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", handler_test)
	http.ListenAndServe(":8888", nil)
}
