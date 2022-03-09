package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(1, 2, 4)
	resultado := 6
	if teste != 6 {
		t.Error("Expected: ", resultado, "Got: ", teste)
	}
}
