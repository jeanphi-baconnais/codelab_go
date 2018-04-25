package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"math/rand"
	"strconv"
)

const cpt = 0

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/generate/{value}", Generation)
	log.Fatal(http.ListenAndServe(":9999", router))
}

func Generation(w http.ResponseWriter, r *http.Request) {
	sleep := rand.Intn(1000)
	value := mux.Vars(r)["value"]
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message" : "=> ` + value + ` | ` + strconv.Itoa(sleep) + `"}`))

	// Mise en pause du serveur pendant un nombre de millisecondes aléatoire
	time.Sleep( time.Duration(sleep) * time.Millisecond)
}
