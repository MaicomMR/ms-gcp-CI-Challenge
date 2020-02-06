package main

import "testing"

func TestSomaDois(t *testing.T){

	resultado := somaDois(5, 5)

	if resultado != 10 {
		t.Error("╔════════════════════════════════════════════╗")
		t.Error("║         Falha ao somar dois valores        ║")
		t.Error("╚════════════════════════════════════════════╝")

	}

}