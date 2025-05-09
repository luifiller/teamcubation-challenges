package main

import (
	"fmt"
	"strings"
)

func inverterOrdem(frase string) string {
	fmt.Println("Invertendo ordem... ")
	palavras := strings.Split(frase, " ")

	var resultado strings.Builder
	resultado.Grow(len(palavras))

	for i := len(palavras) - 1; i >= 0; i-- {
		resultado.WriteString(palavras[i])

		if i != 0 {
			resultado.WriteString(" ")
		}
	}

	return resultado.String()
}

func inverterCaracteres(frase string) string {
	fmt.Println("Invertendo caracteres...")
	palavras := strings.Split(frase, " ")
	var resultado strings.Builder

	for i, palavra := range palavras {
		runas := []rune(palavra)

		var palavraInvertida strings.Builder
		palavraInvertida.Grow(len(runas))
		for j := len(runas) - 1; j >= 0; j-- {
			palavraInvertida.WriteRune(runas[j])
		}

		if i != 0 {
			resultado.WriteString(" ")
		}

		resultado.WriteString(palavraInvertida.String())
	}

	return resultado.String()
}

func inverterOrdemECaracteres(frase string) string {
	fmt.Println("Invertendo ordem e caracteres...")
	palavras := strings.Split(frase, " ")

	var resultado strings.Builder
	resultado.Grow(len(palavras))
	for i := len(palavras) - 1; i >= 0; i-- {
		runas := []rune(palavras[i])

		var palavraInvertida strings.Builder
		for j := len(runas) - 1; j >= 0; j-- {
			palavraInvertida.WriteRune(runas[j])
		}

		resultado.WriteString(palavraInvertida.String())
		if i != 0 {
			resultado.WriteString(" ")
		}
	}

	return resultado.String()
}

func inverterAcumulando(frase string) string {
	fmt.Println("Invertendo com método acumulativo...")

	var invertida string
	for _, char := range frase {
		invertida = string(char) + invertida
	}

	return invertida
}

func concluirInversoes(resultado string) {
	fmt.Println("Inversão concluída...")
	fmt.Println("Frase formatada: " + resultado)
}
