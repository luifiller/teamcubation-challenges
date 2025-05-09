package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numNuvens := lerQtdNuvens()
	estadosNuvens := lerEstadosDasNuvens(numNuvens)

	saltos := saltarEmNuvens(estadosNuvens)
	fmt.Printf("Número de saltos: %d \n", saltos)
}

func lerQtdNuvens() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("+----------------------------------------------------------+")
		fmt.Println("|                    VIAJANTE DAS NUVENS                   |")
		fmt.Println("+----------------------------------------------------------+")
		fmt.Print("Insira a quantidade de nuvens (entre 2 e 100): ")

		qtdNuvens, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao criar quantidade de nuvens.")
			continue
		}

		qtdNuvens = strings.TrimSpace(qtdNuvens)
		numNuvens, err := strconv.Atoi(qtdNuvens)
		if err != nil {
			fmt.Println("Erro ao converter string em integer.")
			continue
		}

		if numNuvens < 2 || numNuvens > 100 {
			fmt.Println("A quantidade de nuvens deve ser entre 2 e 100.")
			continue
		}

		return numNuvens
	}
}

func lerEstadosDasNuvens(qtdNuvens int) []int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Insira os %d estados das nuvens (0 = segura, 1 = perigosa), separados por espaço: ", qtdNuvens)
		estadosNuvens, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler os estados das nuvens.")
			continue
		}

		estadosNuvens = strings.TrimSpace(estadosNuvens)
		estadosString := strings.Split(estadosNuvens, " ")
		if len(estadosString) != qtdNuvens {
			fmt.Println("Quantidade de estados diferente da quantidade de nuvens informada previamente.")
			continue
		}

		valid := true

		estadosNuvensInt := make([]int, qtdNuvens)
		for i, estado := range estadosString {
			estadoInt, err := strconv.Atoi(estado)

			if err != nil || (estadoInt != 0 && estadoInt != 1) {
				fmt.Println("Os estados das nuvens devem ser 0 ou 1.")
				valid = false
				break
			}

			estadosNuvensInt[i] = estadoInt
		}

		if !valid {
			continue
		}

		if estadosNuvensInt[0] != 0 || estadosNuvensInt[qtdNuvens-1] != 0 {
			fmt.Println("A primeira e a última nuvem devem ser seguras (0).")
			continue
		}

		return estadosNuvensInt
	}
}

func saltarEmNuvens(nuvens []int) int {
	inicio := 0
	var qtdSaltos int

	for inicio < len(nuvens)-1 {
		if inicio+2 < len(nuvens) && nuvens[inicio+2] == 0 {
			inicio += 2
		} else {
			inicio++
		}
		qtdSaltos++
	}

	return qtdSaltos
}
