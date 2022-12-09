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

func main() {
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	fmt.Println(contaDaSilvia)
}
