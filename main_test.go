package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	assert.Equal(t, total, 10, "not equals")
}

func TestSumNil(t *testing.T) {
	total := Sum(-1, 5)
	assert.Equal(t, total, 4, "not equals")
}

func TestReturnValeurAndValuesMutliplieesZero(t *testing.T) {
	x, y := ReturnValeurAndValuesMutlipliees(2, 0)

	assert.Equal(t, x, 2, "not equals")
	assert.Equal(t, y, 0, "not equals")
}


func TestReturnValeurAndValuesMutlipliees(t *testing.T) {
	x, y := ReturnValeurAndValuesMutlipliees(2, 5)

	assert.Equal(t, x, 2, "not equals")
	assert.Equal(t, y, 10, "not equals")
}

func TestCreatePersonne(t *testing.T) {
	personne := CreatePersonne("test", 20)
	assert.Equal(t, personne.name, "test", "not equals")


	personneNoAge := CreatePersonne("test", nil)
	assert.Equal(t, personneNoAge.age, "test", "not equals")
}
