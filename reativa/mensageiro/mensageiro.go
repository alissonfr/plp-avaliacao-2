package mensageiro

import "plp-avaliacao-2/reativa/mensagem"

type Mensageiro[Texto any, Arquivo any] interface {
	EnviarTexto(texto Texto) mensagem.Resultado
	EnviarAudio(audio Arquivo) mensagem.Resultado
	EnviarVideo(video Arquivo) mensagem.Resultado
	EnviarImagem(imagem Arquivo) mensagem.Resultado
}

