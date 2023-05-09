package main

import "fmt"

func main() {
	nome := "Halysson"
	idade := 34
	versao := 1.1
	comando := 0

	fmt.Println("Olá, ", nome, " sua idade é: ", idade)
	fmt.Println("Este programa está na versão: ", versao)
	fmt.Println("")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")

	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi : ", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Não conheço este comando")
	}

}
