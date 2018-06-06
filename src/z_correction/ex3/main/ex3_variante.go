package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

var port string = "8888"

func main() {
	fmt.Println("\nExercice 3 : ")
	fmt.Println("Serveur démarré sur le port ", port)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Handler)
	router.HandleFunc("/mars", Handler_mars)
	log.Fatal(http.ListenAndServe(":" + port, router))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message" : "Allo la Terre !"}`))
}

func Handler_mars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message" : "On a trouvé de l'eau sur Mars ?"}`))
}