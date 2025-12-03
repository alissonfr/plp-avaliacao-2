package main

import (
	"fmt"
	"os"
	"plp-avaliacao-2/reativa/impl"
	"plp-avaliacao-2/reativa/impl/mensageiros"
	"plp-avaliacao-2/reativa/impl/mensagens"
	"plp-avaliacao-2/reativa/mensageiro"
	"plp-avaliacao-2/reativa/mensagem"
	"time"
)

func main() {
	var mensageiro mensageiro.Mensageiro[string, *os.File] = mensageiros.NewWhatsApp()
	erros := make([]*mensagens.Erro, 0)

	mensagensChan := make(chan mensagem.Mensagem, 10)
	processadasChan := make(chan string, 10)

	go func() {
		var imgFile *os.File
		var audioFile *os.File

		mensagensList := []mensagem.Mensagem{
			mensagens.NewTexto(mensageiro, "olá, como vai?"),
			mensagens.NewTexto(mensageiro, "iaí, amigão?"),
			mensagens.NewImagem(mensageiro, imgFile),
			mensagens.NewErro("falha de comunicação"),
			mensagens.NewAudio(mensageiro, audioFile),
		}

		for _, msg := range mensagensList {
			time.Sleep(10 * time.Millisecond)
			mensagensChan <- msg
		}
		close(mensagensChan)
	}()

	go func() {
		for msg := range mensagensChan {
			msg.SetDataHora(time.Now())

			if msg.GetTipo() == impl.ERRO {
				erro := msg.(*mensagens.Erro)
				erros = append(erros, erro)
				fmt.Printf("erro: %s\n", erro.String())
				continue
			}

			if msg.Enviar() == mensagem.ERRO {
				erro := mensagens.NewErro("ocorreu um erro de envio do(a) " + msg.GetTipo().(impl.TipoMensagem).String())
				erros = append(erros, erro)
				fmt.Printf("erro: %s\n", erro.String())
				continue
			}

			processadasChan <- fmt.Sprintf("processada mensagem do tipo %s", msg.GetTipo().(impl.TipoMensagem).String())
		}
		close(processadasChan)
	}()

	for processada := range processadasChan {
		fmt.Println(processada)
	}

	fmt.Println("processamento concluído")
}
