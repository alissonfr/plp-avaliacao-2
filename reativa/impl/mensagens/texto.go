package mensagens

import (
	"os"
	"plp-avaliacao-2/reativa/impl"
	"plp-avaliacao-2/reativa/mensagem"
	"plp-avaliacao-2/reativa/mensageiro"
)

type Texto struct {
	*mensagem.BaseMensagem
	mensageiro mensageiro.Mensageiro[string, *os.File]
}

func NewTexto(mensageiro mensageiro.Mensageiro[string, *os.File], conteudo string) *Texto {
	return &Texto{
		BaseMensagem: mensagem.NewBaseMensagem(impl.TEXTO, conteudo),
		mensageiro:   mensageiro,
	}
}

func (t *Texto) Enviar() mensagem.Resultado {
	return t.mensageiro.EnviarTexto(t.GetConteudo().(string))
}

