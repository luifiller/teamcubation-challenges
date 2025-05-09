package main

import (
	"fmt"
	"strings"
)

func inverterOrdem(frase string) string {
	fmt.Println("Invertendo ordem... ")
	palavras := strings.Split(frase, " ")

	var fraseInvertida string
	for i := len(palavras) - 1; i >= 0; i-- {
		fraseInvertida += palavras[i]

		if i != 0 {
			fraseInvertida += " "
		}
	}

	return fraseInvertida
}

func inverterCaracteres(frase string) string {
	fmt.Println("Invertendo caracteres...")
	palavras := strings.Split(frase, " ")

	var fraseFormatada string
	for index, palavra := range palavras {
		runas := []rune(palavra)

		var palavraInvertida string
		for i := len(runas) - 1; i >= 0; i-- {
			palavraInvertida += string(runas[i])

		}

		if index != 0 {
			fraseFormatada += " "
		}
		fraseFormatada += palavraInvertida
	}

	return fraseFormatada
}

func inverterOrdemECaracteres(frase string) string {
	fmt.Println("Invertendo ordem e caracteres...")
	palavras := strings.Split(frase, " ")

	var fraseFormatada string
	for i := len(palavras) - 1; i >= 0; i-- {
		runas := []rune(palavras[i])

		var palavraInvertida string
		for j := len(runas) - 1; j >= 0; j-- {
			palavraInvertida += string(runas[j])
		}

		fraseFormatada += palavraInvertida
		if i != 0 {
			fraseFormatada += " "
		}
	}

	return fraseFormatada
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
