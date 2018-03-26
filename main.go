package main

import (
	"fmt"
	"os"
)

func main() {
	var message string    	// manual type déclaration

	// var args []string
	args := os.Args 		// type inference

	if len(args) > 1 {
		message = args[1]
	} else {
		message = "Texte par défaut"
	}

	fmt.Printf(message)

}
