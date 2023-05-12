package main

import "testing"

func TestSoma(t *testing.T) {
	// arrange
	a := 1
	b := 2
	expected := 3
	// act
	actual := Soma(a, b)

	// assert
	if actual != expected {
		t.Errorf("Resultado da soma é inválido: Actual: %v, Expected: %v", actual, expected)
	}

}
