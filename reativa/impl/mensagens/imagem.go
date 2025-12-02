package mensagens

import (
	"os"
	"plp-avaliacao-2/reativa/impl"
	"plp-avaliacao-2/reativa/mensagem"
	"plp-avaliacao-2/reativa/mensageiro"
)

type Imagem struct {
	*mensagem.BaseMensagem
	mensageiro mensageiro.Mensageiro[string, *os.File]
}

func NewImagem(mensageiro mensageiro.Mensageiro[string, *os.File], conteudo *os.File) *Imagem {
	return &Imagem{
		BaseMensagem: mensagem.NewBaseMensagem(impl.IMAGEM, conteudo),
		mensageiro:   mensageiro,
	}
}

func (i *Imagem) Enviar() mensagem.Resultado {
	return i.mensageiro.EnviarImagem(i.GetConteudo().(*os.File))
}

