package mensagens

import (
	"os"
	"plp-avaliacao-2/reativa/impl"
	"plp-avaliacao-2/reativa/mensagem"
	"plp-avaliacao-2/reativa/mensageiro"
)

type Audio struct {
	*mensagem.BaseMensagem
	mensageiro mensageiro.Mensageiro[string, *os.File]
}

func NewAudio(mensageiro mensageiro.Mensageiro[string, *os.File], conteudo *os.File) *Audio {
	return &Audio{
		BaseMensagem: mensagem.NewBaseMensagem(impl.AUDIO, conteudo),
		mensageiro:   mensageiro,
	}
}

func (a *Audio) Enviar() mensagem.Resultado {
	return a.mensageiro.EnviarAudio(a.GetConteudo().(*os.File))
}

