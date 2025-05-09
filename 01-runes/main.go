package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* # DOC

unicode -> padrão universal que define código único para reperesentar caracteres em distintos idiomas, símbolos, emojis, sinais.

rune -> alias de int32 e representa/armazena um único caracter unicode
	|-> Útil para trabalhar com texto internacionalizado, emojis

Lidar com runes
- Converter string p/ slice de runes com o `[]rune(nomeString)`
- Há como iterar diretamente sobre a string usando `for i, r := range nomeString`
	|-> `r` é uma rune (ponto de código)
	|-> `i` é a posição em bytes
*/

/* # CHALLENGE
Implemente funções em Go que realizem as seguintes operações:

- Inverter palavras de uma frase: Dada uma string, inverta a ordem das palavras mantendo seu conteúdo original.
- Inverter caracteres de uma string (método com runas): Retorna uma string com seus caracteres em ordem inversa, utilizando um slice de runas.
- Inverter caracteres de uma string (método acumulativo): Retorna uma string com seus caracteres em ordem inversa, acumulando-os em uma nova string.
*/

func main() {
	novaFrase := true
	reader := bufio.NewReader(os.Stdin)

	for novaFrase {
		fmt.Println("+----------------------------------------------------------+")
		fmt.Println("|                 INVERSOR 3000 TEAMCUBATION               |")
		fmt.Println("+----------------------------------------------------------+")
		fmt.Print("Digite uma frase: ")

		frase, error := reader.ReadString('\n')
		if error != nil || len(frase) == 0 {
			fmt.Println("Erro ao ler a frase. Tente novamente.")
			continue
		}

		frase = strings.TrimSpace(frase)
		if len(frase) == 0 {
			fmt.Println("A frase não contém nenhum caracter. Tente novamente.")
			continue
		}
		novaFrase = false

		for !novaFrase {
			exibirMenu()

			opcao, error := reader.ReadString('\n')
			if error != nil {
				fmt.Println("Erro ao processar opção. Tente novamente.")
			}
			opcao = strings.TrimSpace(opcao)

			switch opcao {
			case "0":
				fmt.Println("Saindo... Até logo!")
				return
			case "1":
				inverterOrdem(frase)
			case "2":
				inverterCaracteres(frase)
			case "3":
				inverterOrdemECaracteres(frase)
			case "4":
				inverterAcumulando(frase)
			case "5":
				novaFrase = true
			default:
				fmt.Println("Opção não encontrada. Tente novamente.")
			}
		}
	}

}

func exibirMenu() {
	fmt.Println("+----------------------------------------------------------+")
	fmt.Println("|                      MENU DE INVERSÃO                    |")
	fmt.Println("+----------------------------------------------------------+")
	fmt.Println("| [1] Inverter ordem das palavras                          |")
	fmt.Println("| [2] Inverter caracteres das palavras (ordem mantida)     |")
	fmt.Println("| [3] Inverter ordem e caracteres das palavras             |")
	fmt.Println("| [4] Inverter todos os caracteres (método acumulativo)    |")
	fmt.Println("| [5] Escrever nova frase                                  |")
	fmt.Println("| [0] Sair                                                 |")
	fmt.Println("+----------------------------------------------------------+")
	fmt.Print("Escolha uma opção: ")
}
