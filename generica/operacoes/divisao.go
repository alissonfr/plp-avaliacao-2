package operacoes

type Divisao struct{}

func (Divisao) Calcular(a int, b int) float64 {
	return float64(a) / float64(b)
}
