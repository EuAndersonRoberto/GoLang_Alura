package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {

	exibeIntroducao()

	for { //Golang não faz loop com -> while <- apenas o -> FOR <-
		exibeMenu()

		comando := leComando()

		switch comando { //outra forma de controlar o fluxo. Não há a necessidade do BREAK já que em GO ele executa o caso selecionado e sair.
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Log's...")
			imprimeLogs()
		case 3:
			fmt.Println("Saindo do programa.")
			os.Exit(3) //Pacote os.Exit para informar que saiu com sucesso.
		default:
			fmt.Println("Comando não reconhecido!")
			os.Exit(-1) //O -1 indica que houve algum erro no código ou algo inesperado
		}
	}

}

func exibeIntroducao() {
	nome := "Anderson Roberto"
	versao := 1.7 //hello mundo foi a versão 1.0

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
	fmt.Println("")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir LOG's")
	fmt.Println("3 - Sair do programa")
	fmt.Println("")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.google.com"}
	//fmt.Println("Slice de sites: ", sites)

	sites := lerSitesDoArquivo()
	fmt.Println("")

	// for i := 0; i < len(sites); i++ { //Pecorre o slice mas, a outra alternativa é mais enxuta
	// 	fmt.Println(sites[i])
	// }
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites { //O range é utilizado para indicar a posição dentro do slice e o conteúdo daquela posição.
			fmt.Println("Posição:", i, "do meu slice. Representa o site:", site)
			testaSites(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func exibeLogs() {
	fmt.Println("Apresentando LOG's...")
}

func testaSites(site string) {
	resp, err := http.Get(site) //Requisição de verificação WEB com a variável acima, contendo a URL
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	//fmt.Println("Resposta do monitoramento: ", resp)
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "Foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "Não foi carregado. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt") //aqui ele transforma as informações do arquivo para uma array de bites
	if err != nil {
		fmt.Println("Ocorreu um erro!", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha) //Retirando espaços e saltos de linha desnecessários.
		sites = append(sites, linha)
		//Em substituição a apresentação: fmt.Println(linha)
		if err == io.EOF { //Quando chegar ao fim do arquivo ele para o loop for.
			break
		}

	}

	arquivo.Close()
	//fmt.Println(sites)
	//fmt.Println(string(arquivo)) //Aqui é transformada a array de bites em string
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arquivo)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
