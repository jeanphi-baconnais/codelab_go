package main

import (
	"log"
	"net/http"
	"strconv"
	"io/ioutil"
	"time"
)

var infos = make (chan string)

func main() {
	// Goroutine / channel
	max := 5

	log.Printf("Before call ");
	for i := 0; i < max; i++ {
		go callApiRest(i)
	}
	log.Printf("After call ");

	time.Sleep( 5 * time.Second)

	for i := 0; i < max; i++ {
		log.Printf("Channel" + <- infos);
	}
}


func callApiRest(value int) {
	url := "http://localhost:9999/generate/" +strconv.Itoa(value)

	log.Printf("Call : %s", url)

	response, _ := http.Get(url)

	log.Printf("Call finished : %s", url)

	body, _ := ioutil.ReadAll(response.Body)

	infos <- string(body)

}
