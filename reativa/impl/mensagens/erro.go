package mensagens

import (
	"fmt"
	"plp-avaliacao-2/reativa/impl"
	"plp-avaliacao-2/reativa/mensagem"
)

type Erro struct {
	*mensagem.BaseMensagem
}

func NewErro(conteudo string) *Erro {
	return &Erro{
		BaseMensagem: mensagem.NewBaseMensagem(impl.ERRO, conteudo),
	}
}

func (e *Erro) Enviar() mensagem.Resultado {
	return mensagem.ERRO
}

func (e *Erro) String() string {
	return fmt.Sprintf("%v: %s", e.GetDataHora(), e.GetConteudo().(string))
}

