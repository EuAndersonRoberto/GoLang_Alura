package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso! Saldo atual:", c.saldo
	} else {
		return "Saldo insuficiente para depósito! Saldo atual:", c.saldo
	}
}

func (c *ContaCorrente)transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool{ 
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0{
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}else{
		return false
	}
}

func main() {
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contadoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}
	
	status := contaDaSilvia.transferir(200, &contadoGustavo)
	
	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contadoGustavo)
}