package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoAnderson := ContaCorrente{titular: "Anderson", numeroAgencia: 589, numeroConta: 123456, saldo: 250.5}

	contaDaAna := ContaCorrente{"Ana", 222, 111222, 200.1}

	fmt.Println("", contaDoAnderson, "\n", contaDaAna)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.numeroAgencia = 555
	contaDaCris.numeroConta = 333444
	contaDaCris.saldo = 500

	fmt.Println(contaDaCris)
}