package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	var message string    	// manual type déclaration

	// var args []string
	args := os.Args[1:]		// type inference
	nbParam := len(args)

	fmt.Println("\nExercice 2 : ")

	if nbParam == 0 {
		message = "Aucun paramètre !"
	} else if nbParam == 1 {
		message = args[0]
	} else {
		err := errors.New("Paramètres invalides \n")
		fmt.Println("\t",err)
		panic(fmt.Sprintf("Mayday !"))
	}

	fmt.Println("\t", message)
	fmt.Println("--------------")

}
