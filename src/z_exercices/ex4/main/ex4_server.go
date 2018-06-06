package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"math/rand"
	"strconv"
	"fmt"
)

const port string = "9999"

func main() {
	fmt.Println("\nExercice 4 : ")
	fmt.Println("Serveur démarré sur le port ", port)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{value}", Generation)
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
