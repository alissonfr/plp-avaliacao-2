package mensagem

type Resultado int

const (
	SUCESSO Resultado = iota
	ERRO
)

func (r Resultado) String() string {
	switch r {
	case SUCESSO:
		return "SUCESSO"
	case ERRO:
		return "ERRO"
	default:
		return "UNKNOWN"
	}
}

