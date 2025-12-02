package main

import (
	"fmt"
	"plp-avaliacao-2/generica/operacoes"
)

func main() {
	multiplicacao := operacoes.NewMultiplicacao()
	divisao := operacoes.NewDivisao()

	fmt.Println(multiplicacao.Calcular(3, 2))
	fmt.Println(divisao.Calcular(3, 2))
}
