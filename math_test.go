package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(20, 20)

	if total != 40 {
		t.Errorf("Resultado inválido: Resulta %d. Esperado %d", total, 20)
	}
}

func TestSum0(t *testing.T) {
	total := Sum0(20, 20)

	if total != 40 {
		t.Errorf("Resultado inválido: Resulta %d. Esperado %d", total, 20)
	}
}

func TestSum1(t *testing.T) {
	total := Sum1(20, 20)

	if total != 40 {
		t.Errorf("Resultado inválido: Resulta %d. Esperado %d", total, 20)
	}
}

func TestSum2(t *testing.T) {
	total := Sum2(20, 20)

	if total != 40 {
		t.Errorf("Resultado inválido: Resulta %d. Esperado %d", total, 20)
	}
}

func TestSum3(t *testing.T) {
	total := Sum3(20, 20)

	if total != 40 {
		t.Errorf("Resultado inválido: Resulta %d. Esperado %d", total, 20)
	}
}
