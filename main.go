package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Goroutine / channel
	max := 35

	for i := 0; i < max; i++ {
		callApiRest(i)
	}
}
func callApiRest(value int) {
	url := "http://localhost:9999/generate/" +strconv.Itoa(value)
	log.Printf("Call : %s", url)
	http.Get(url)
	log.Printf("Call finished : %s", url)
}
