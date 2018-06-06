package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

const (
	portAPI              = "9999"
	// TODO : ajouter les infos de la BDD mongo.
)

// TODo : déclarer une variable Session

func main() {
	fmt.Println("\nExercice 7 : ")

	// TODO : connexion à la BDD mongo

	fmt.Println("\tServeur démarré sur le port ", portAPI)
	router := mux.NewRouter().StrictSlash(true)

	// TODO : ajouter les handlers

	log.Fatal(http.ListenAndServe(":9999", router))

}

// TODO : créer la structure de planet : name (string), diameter (int)

// TODO : créer les méthodes de requetage à la base de données.