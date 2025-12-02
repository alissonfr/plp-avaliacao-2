package operacoes

type Multiplicacao struct{}

func NewMultiplicacao() *Multiplicacao {
	return &Multiplicacao{}
}

func (m *Multiplicacao) Calcular(operando1 int, operando2 int) int {
	return operando1 * operando2
}

