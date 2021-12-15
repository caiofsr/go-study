package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Caio")
	esperado := "Olá, Caio"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
