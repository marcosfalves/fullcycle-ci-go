package main

import "testing"

func TestSoma(t *testing.T) {

	result := sum(15, 15)

	if result != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", result, 30)
	}
}
