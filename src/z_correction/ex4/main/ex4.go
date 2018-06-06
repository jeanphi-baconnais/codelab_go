package main

import (
	"log"
	"net/http"
	"strconv"
	"io/ioutil"
	"fmt"
)

var channel = make(chan string)

func main() {
	max := 35

	for i := 0; i < max; i++ {
		go callApiRest(i)
	}

	<- channel // wait all goroutines

	for i := 0; i < max; i++ {
		fmt.Println("Channel %s", <- channel)
	}
}

func callApiRest(value int) {

	url := "http://localhost:9999/" + strconv.Itoa(value)
	log.Printf("Call : %s", url)

	// call api
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)

	// stockage dans le channel
	channel <- string(body)
	log.Printf("Call finished : %s", url)

}
