package main

import (
	"log"
	"net/http"
	"strconv"
)

// TODO : déclarer un channel

func main() {
	max := 35

	for i := 0; i < max; i++ {
		callApiRest(i)
	}
}
func callApiRest(value int) {
	url := "http://localhost:9999/" + strconv.Itoa(value)
	log.Printf("Call : %s", url)
	http.Get(url)
	log.Printf("Call finished : %s", url)

	// TODO :
	// - stocker la réponse de l’appel à la méthode http.Get dans une variable
	// - stocker le body via ioutil.ReadAll, comme une gestion de fichier
	// - stocker dans le channel du body de la réponse

}
