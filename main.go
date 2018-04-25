package main

import "errors"

func main() {

}

type person struct {
	name string
	age  int
}

func Sum (x int, y int) int {
	return x + y
}

/*
	Fonctions à tester.
 */

// Valeurs de retour multiples
func ReturnValeurAndValuesMutlipliees(x int, y int) (int, int) {
	return x, x*y
}

// Création d'une personne
func CreatePersonne(name string, age int) person {
	return person {name, age}
}

// Tests des erreurs
func Function42throwError(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Error, not 42 !")
	}
	return arg + 3, nil
}