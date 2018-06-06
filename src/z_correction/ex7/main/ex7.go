package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"strconv"
)

const (
	portAPI              = "9999"
	mongodb_serveur      = "localhost"
	mongodb_serveur_port = "27017"
	mongodb_database     = "codelabgo"
	mongodb_collection   = "planets"
)

var session *mgo.Session

func main() {
	fmt.Println("\nExercice 7 : ")

	fmt.Println("\tConnexion à la base mongo : ", mongodb_serveur, ":", mongodb_serveur_port, "/", mongodb_database)
	Connect();

	fmt.Println("\tServeur démarré sur le port ", portAPI)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/planets/all", GetAllPlanets).Methods("GET")
	router.HandleFunc("/planets", GetPlanetByName).Methods("GET")
	router.HandleFunc("/planets", CreatePlanet).Methods("POST")
	router.HandleFunc("/planets", RemovePlanet).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", router))

}

func GetAllPlanets(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("content-type", "application/json")

	planets, _ := getAllPlanets()

	w.Header().Set("content-type", "application/json")
	planetsJson, _ := json.MarshalIndent(planets, "", "  ")

	w.Write(planetsJson)
}

func GetPlanetByName(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	planets, _ := getPlanetByName(r.FormValue("name"))

	w.Header().Set("content-type", "application/json")
	if len(planets) > 0 {
		planetsJson, _ := json.MarshalIndent(planets, "", "  ")
		w.Write(planetsJson)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func RemovePlanet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := removePlanets(r.FormValue("name"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func CreatePlanet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	diameter, _ := strconv.Atoi(r.URL.Query().Get("diameter"))

	err := createPlanet(Planet{Name : r.URL.Query().Get("name"), Diameter: diameter})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

type Planet struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Diameter int
}

func Connect() {
	var err error
	session, err = mgo.Dial(mongodb_serveur)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllPlanets() ([]Planet, error) {
	sessionCopy := session.Copy() // duplicate session for concurrent processing.
	defer sessionCopy.Close()

	var planets []Planet
	err := sessionCopy.DB(mongodb_database).C(mongodb_collection).Find(bson.M{}).All(&planets)
	return planets, err
}

func getPlanetByName(value string) ([]Planet, error) {
	sessionCopy := session.Copy() // duplicate session for concurrent processing.
	defer sessionCopy.Close()

	var planets []Planet
	err := sessionCopy.DB(mongodb_database).C(mongodb_collection).Find(bson.M{"name": value}).All(&planets)
	return planets, err
}

func removePlanets(planetsToRemove string) error {
	sessionCopy := session.Copy() // duplicate session for concurrent processing.
	defer sessionCopy.Close()

	_, err := sessionCopy.DB(mongodb_database).C(mongodb_collection).RemoveAll(bson.M{"name": planetsToRemove});

	return err
}

func createPlanet(newPlanet Planet) error {
	sessionCopy := session.Copy() // duplicate session for concurrent processing.
	defer sessionCopy.Close()
	err := sessionCopy.DB(mongodb_database).C(mongodb_collection).Insert(newPlanet);
	return err
}
