package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	)

func TestSumBis(t *testing.T) {
	total := Sum(5, 5)
	assert.Equal(t, total, 10, "not equals")
}

func TestSumNilBis(t *testing.T) {
	total := Sum(-1, 5)
	assert.Equal(t, total, 4, "not equals")
}

func TestReturnValeurAndValuesMutliplieesZeroBis(t *testing.T) {
	x, y := ReturnValeurAndValuesMutlipliees(2, 0)

	assert.Equal(t, x, 2, "not equals")
	assert.Equal(t, y, 0, "not equals")
}


func TestReturnValeurAndValuesMutliplieesBis(t *testing.T) {
	x, y := ReturnValeurAndValuesMutlipliees(2, 5)

	assert.Equal(t, x, 2, "not equals")
	assert.Equal(t, y, 10, "not equals")
}

func TestCreatePersonneBis(t *testing.T) {
	personne := CreatePersonne("test", 20)
	assert.Equal(t, personne.name, "test", "not equals")

	personneNoAge := CreatePersonne("test", 0)
	assert.Equal(t, personneNoAge.name, "test", "not equals")
	assert.Equal(t, personneNoAge.age, 0)
}
