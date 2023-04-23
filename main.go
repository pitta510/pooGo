package main

import (
	"fmt"
	"pooGo/contas"
)

func PagarBoleto(conta verificarConta, ValorDoBoleto float64) {
	conta.Sacar(ValorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDennis := contas.ContaPoupanca{}
	contaDoDennis.Depositar(100)
	PagarBoleto(&contaDoDennis, 60)

	fmt.Println(contaDoDennis.ObterSaldo())
}
