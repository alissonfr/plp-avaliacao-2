package mensageiros

import (
	"fmt"
	"os"
	"plp-avaliacao-2/reativa/mensagem"
)

type WhatsApp struct{}

func NewWhatsApp() *WhatsApp {
	return &WhatsApp{}
}

func (w *WhatsApp) EnviarTexto(texto string) mensagem.Resultado {
	fmt.Println("texto enviado")
	return mensagem.SUCESSO
}

func (w *WhatsApp) EnviarAudio(audio *os.File) mensagem.Resultado {
	fmt.Println("áudio enviado")
	return mensagem.SUCESSO
}

func (w *WhatsApp) EnviarVideo(video *os.File) mensagem.Resultado {
	fmt.Println("vídeo enviado")
	return mensagem.SUCESSO
}

func (w *WhatsApp) EnviarImagem(imagem *os.File) mensagem.Resultado {
	fmt.Println("imagem enviada")
	return mensagem.SUCESSO
}

