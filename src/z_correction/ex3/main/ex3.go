package main

import (
	"net/http"
	"fmt"
)
var port string = "8888"

func handler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message" : "Allo la Terre !"}`))
}

func handler_mars (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message" : "On a trouvé de l'eau sur Mars ?"}`))
}

func main() {
	fmt.Println("\nExercice 3 : ")
	fmt.Println("Serveur démarré sur le port ", port)

	http.HandleFunc("/", handler)
	http.HandleFunc("/mars", handler_mars)
	http.ListenAndServe(":" + port, nil)
}
