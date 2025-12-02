package main

import (
	"fmt"
	"plp-avaliacao-2/generica/operacoes"
)

func main() {
	var multiplicacao operacoes.Operacao[int] = operacoes.Multiplicacao{}
	var divisao operacoes.Operacao[float64] = operacoes.Divisao{}

	fmt.Println(multiplicacao.Calcular(3, 2))
	fmt.Println(divisao.Calcular(3, 2))
}
