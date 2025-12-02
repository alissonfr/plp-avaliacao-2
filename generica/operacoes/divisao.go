package operacoes

type Divisao struct{}

func NewDivisao() *Divisao {
	return &Divisao{}
}

func (d *Divisao) Calcular(operando1 int, operando2 int) float64 {
	return float64(operando1) / float64(operando2)
}

