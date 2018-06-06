package main

func main() {

	// 1 : utiliser la méthode ListenAndServe du package http

	// 2 :
	// - définir des HandleFunc pour définir le path et le handler à créer
	// - définir une fonction avec le handler évoqué préalablement.
	// Cette fonction prend comme paramètre un ResponseWriter et une Request du package http

	// 3 :
	// - Modifier via un Set le Header de la ResponseWriter pour y ajouter un content-type
	// - Ecrire dans la Request pour renvoyer une valeur json

}
