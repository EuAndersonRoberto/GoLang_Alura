package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()

	exibeMenu()

	// if comando == 1 { //Não há parenteses e só recebe expressões boleanas.
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Apresentando LOG's...")
	// } else if comando == 3 {
	// 	fmt.Println("Saindo do programa!")
	// } else {
	// 	fmt.Println("Comando não reconhecido!")
	// }
	comando := leComando()

	switch comando { //outra forma de controlar o fluxo. Não há a necessidade do BREAK já que em GO ele executa o caso selecionado e sair.
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Apresentando LOG's...")
	case 3:
		fmt.Println("Saindo do programa.")
		os.Exit(3) //Pacote os.Exit para informar que saiu com sucesso.
	default:
		fmt.Println("Comando não reconhecido!")
		os.Exit(-1) //O -1 indica que houve algum erro no código ou algo inesperado
	}
}

func exibeIntroducao() {
	nome := "Anderson Roberto"
	versao := 1.3 //hello mundo foi a versão 1.0

	fmt.Println("Olá Sr: ", nome)
	fmt.Println("Programa versão: ", versao)
	fmt.Println("")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) //A função Scan identifica o tipo pela definição que consta na variável (var comando int)
	//fmt.Scanf("%d", &comando) //com definição de que a variável é um inteiro (%d)
	fmt.Println("")
	fmt.Println("O endereço/ponteiro do meu comando é: ", &comandoLido)
	fmt.Println("O comando escolhido foi: ", comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir LOG's")
	fmt.Println("3 - Sair do programa")
}
