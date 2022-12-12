package main

import "fmt"

func main() {
	var nome = "Anderson Roberto"
	var versao = 1.2 //hello mundo foi a versão 1.0

	fmt.Println("Olá Sr: ", nome)
	fmt.Println("Programa versão: ", versao)
	fmt.Println("")

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir LOG's")
	fmt.Println("3 - Sair do programa")

	var comando int
	fmt.Scan(&comando) //A função Scan identifica o tipo pela definição que consta na variável (var comando int)
	//fmt.Scanf("%d", &comando)//com definição de que a variável é um inteiro (%d)
	fmt.Println("")
	fmt.Println("O endereço/ponteiro do meu comando é: ", &comando)
	fmt.Println("O comando escolhido foi: ", comando)
}
