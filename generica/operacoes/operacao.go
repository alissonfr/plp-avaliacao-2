package operacoes

type Operacao[TipoResultado any] interface {
	Calcular(operando1 int, operando2 int) TipoResultado
}

