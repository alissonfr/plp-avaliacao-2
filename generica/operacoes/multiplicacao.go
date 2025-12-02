package operacoes

type Multiplicacao struct{}

func (Multiplicacao) Calcular(a int, b int) int {
	return a * b
}
