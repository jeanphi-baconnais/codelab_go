package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Somme incorrecte, actuelle : %d, souhaitée : %d.", total, 10)
	}
}

func TestSumNil(t *testing.T) {
	total := Sum(-1, 5)
	if total != 4 {
		t.Errorf("Somme incorrecte, actuelle : %d, souhaitée : %d.", total, 10)
	}
}

func TestReturnValeurAndValuesMutliplieesZero(t *testing.T) {
	x, y := ReturnValeurAndValuesMutlipliees(2, 0)
	if x != 2 && y != 0 {
		t.Errorf("Résultat incorrect, actuel : %d %d, souhaité : %d %d.", x , y , 2, 0)
	}
}


func TestReturnValeurAndValuesMutlipliees(t *testing.T) {
	x, y := ReturnValeurAndValuesMutlipliees(2, 5)
	if x != 2 && y != 10 {
		t.Errorf("Résultat incorrect, actuel : %d %d, souhaité : %d %d.", x , y , 2, 10)
	}
}

func TestCreatePersonne(t *testing.T) {
	personne := CreatePersonne("test", 20)

	if personne.name != "test" {
		t.Errorf("Nom incorrect, actuel : %d, attendu : %d", personne.name, "test")
	}

}