package mensagens

import (
	"os"
	"plp-avaliacao-2/reativa/impl"
	"plp-avaliacao-2/reativa/mensagem"
	"plp-avaliacao-2/reativa/mensageiro"
)

type Video struct {
	*mensagem.BaseMensagem
	mensageiro mensageiro.Mensageiro[string, *os.File]
}

func NewVideo(mensageiro mensageiro.Mensageiro[string, *os.File], conteudo *os.File) *Video {
	return &Video{
		BaseMensagem: mensagem.NewBaseMensagem(impl.VIDEO, conteudo),
		mensageiro:   mensageiro,
	}
}

func (v *Video) Enviar() mensagem.Resultado {
	return v.mensageiro.EnviarVideo(v.GetConteudo().(*os.File))
}

