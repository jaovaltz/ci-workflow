package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Errorf("Soma incorreta, obtido: %d, esperado: %d.", total, 30)
	}
}