package main

import (
	"log"
	"net/http"
	"strconv"
	"io/ioutil"
)

var c = make (chan string)

func main() {
	// Goroutine / channel
	max := 35
	for i := 0; i < max; i++ {
		go callApiRest(i)
	}

	for i := 0; i < max; i++ {
		log.Printf("Channel : %s", <- c)
	}


}
func callApiRest(value int) {
	url := "http://localhost:9999/generate/" +strconv.Itoa(value)
	log.Printf("Call : %s", url)
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	c <- string(body)
	log.Printf("Call finished : %s", url)
}
