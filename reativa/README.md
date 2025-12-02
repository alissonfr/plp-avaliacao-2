# Programação Reativa em Go

## Propósito

Este programa demonstra uma aproximação ao paradigma de programação reativa em Go usando goroutines e channels, seguindo a estrutura do exemplo Java com Reactor. O sistema implementa:

- **Mensageiro**: Interface genérica para envio de mensagens (WhatsApp como implementação)
- **Mensagem**: Classes de mensagens (Texto, Audio, Imagem, Video, Erro) que estendem BaseMensagem
- **Stream reativo**: Uso de channels e goroutines para processar mensagens de forma assíncrona
- **Tratamento de erros**: Captura e tratamento de erros sem interromper o fluxo
- **Processamento assíncrono**: Mensagens são processadas assim que chegam

O programa demonstra os princípios reativos: orientado a eventos, elasticidade, responsividade e resiliência.

## Linguagem

Golang (Go)

## Paradigma demonstrado

**Reativo** - Modelo baseado em fluxos assíncronos de dados, onde tudo é um stream e você reage a mudanças usando observables. Foca em não bloqueio, backpressure e alta escalabilidade.

## Autor(es)

Alisson Rodrigues e Jamille Monteiro

## Como executar

```bash
cd reativa
go run main.go
```

## Estrutura do código

- **mensageiro/mensageiro.go**: Interface genérica `Mensageiro[Texto, Arquivo]`
- **mensagem/mensagem.go**: Interface `Mensagem` e struct base `BaseMensagem`
- **mensagem/resultado.go**: Enum `Resultado` (SUCESSO, ERRO)
- **impl/tipomensagem.go**: Enum `TipoMensagem` (TEXTO, IMAGEM, VIDEO, AUDIO, ERRO)
- **impl/mensageiros/whatsapp.go**: Implementação de `Mensageiro`
- **impl/mensagens/**: Classes de mensagens (Texto, Audio, Imagem, Video, Erro)
- **main.go**: Fluxo reativo usando channels e goroutines

## Como Go se aproxima do paradigma reativo

### 1. **Streams de Dados**
Channels em Go funcionam como streams: dados fluem através deles de forma contínua, similar ao Flux do Reactor.

### 2. **Push-based Communication**
Quando um produtor envia dados para um channel (`channel <- data`), os consumidores recebem automaticamente quando leem do channel (`data := <-channel`), simulando push de eventos.

### 3. **Processamento Assíncrono**
Goroutines permitem processamento assíncrono de mensagens, similar ao `subscribe` do Reactor.

### 4. **Tratamento de Erros**
Erros são capturados e tratados sem interromper o fluxo, similar ao `onErrorContinue` do Reactor.

### 5. **Elasticidade e Responsividade**
O sistema pode processar múltiplas mensagens simultaneamente através de goroutines, demonstrando elasticidade e responsividade.

