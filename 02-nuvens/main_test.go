package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSaltarEmNuvens(t *testing.T) {
	// Arrange
	stubs := []struct {
		nuvens            []int
		resultadoEsperado int
	}{
		{[]int{0, 1, 0, 0, 0, 1, 0}, 3},
		{[]int{0, 0}, 1},
		{[]int{0, 1, 0}, 1},
		{[]int{0, 0, 0, 0, 0}, 2},
		{[]int{0, 0, 0, 1, 0, 0}, 3},
	}

	for _, stub := range stubs {
		// Act
		resultado := saltarEmNuvens(stub.nuvens)

		// Assert
		if resultado != stub.resultadoEsperado {
			t.Errorf(
				"[Erro saltarEmNuvens()] - Para a entrada %v, era esperado o valor %d de saltos, mas obteve %d",
				stub.nuvens, stub.resultadoEsperado, resultado,
			)
		}
	}
}

func TestLerQtdNuvens(t *testing.T) {
	// Arrange
	stubs := []struct {
		nome              string
		entrada           string
		resultadoEsperado int
	}{
		{
			nome:              "Entrada válida direta",
			entrada:           "10\n",
			resultadoEsperado: 10,
		},
		{
			nome:              "Primeira entrada inválida, depois válida",
			entrada:           "abc\n101\n5\n",
			resultadoEsperado: 5,
		},
		{
			nome:              "Entrada mínima válida",
			entrada:           "2\n",
			resultadoEsperado: 2,
		},
		{
			nome:              "Entrada máxima válida",
			entrada:           "100\n",
			resultadoEsperado: 100,
		},
	}

	// Act
	for _, tt := range stubs {
		t.Run(tt.nome, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.entrada))
			result := lerQtdNuvens(reader)

			// Assert
			if result != tt.resultadoEsperado {
				t.Errorf("lerQtdNuvens() = %d, esperado %d", result, tt.resultadoEsperado)
			}
		})
	}
}
