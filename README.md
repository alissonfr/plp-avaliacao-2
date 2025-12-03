# PLP - Avaliação 2

## Descrição

Este projeto demonstra a implementação de três paradigmas de programação em **Golang (Go)**, desenvolvido como parte da avaliação da disciplina de Paradigmas de Linguagens de Programação (PLP). Cada módulo explora um paradigma diferente, mostrando como Go pode ser utilizado para implementar conceitos fundamentais de programação.

## Estrutura do Projeto

O projeto está organizado em três módulos principais:

### 1. 📦 [Generica](./generica/)
**Paradigma: Programação Genérica**

Demonstra o uso de type parameters (generics) em Go, introduzidos na versão 1.18. Implementa uma calculadora genérica com operações que mantêm type safety em tempo de compilação.

**Conceitos demonstrados:**
- Type parameters e constraints
- Interfaces genéricas
- Type safety em tempo de compilação
- Reutilização de código sem perder segurança de tipos

### 2. 🏗️ [POO](./poo/)
**Paradigma: Programação Orientada a Objetos**

Implementa um sistema de e-commerce completo demonstrando os pilares da POO em Go.

**Conceitos demonstrados:**
- Encapsulamento (campos privados e métodos públicos)
- Interfaces e contratos
- Polimorfismo (múltiplas implementações da interface `Produto`)
- Composição (Go não possui herança clássica)
- Factory Pattern

### 3. ⚡ [Reativa](./reativa/)
**Paradigma: Programação Reativa**

Aproximação ao paradigma reativo usando goroutines e channels, seguindo princípios de programação reativa.

**Conceitos demonstrados:**
- Streams de dados (channels)
- Processamento assíncrono (goroutines)
- Push-based communication
- Tratamento de erros sem interrupção do fluxo
- Elasticidade e responsividade

## Requisitos

- **Go 1.18+** (necessário para o módulo de generics)
- Terminal/Command Prompt

## Como Executar

Cada módulo pode ser executado independentemente:

### Módulo Generica
```bash
cd generica
go run main.go
```

### Módulo POO
```bash
cd poo
go run .
```

### Módulo Reativa
```bash
cd reativa
go run main.go
```

## Autores

- **Alisson Rodrigues**
- **Jamille Monteiro**

## Estrutura de Diretórios

```
plp-avaliacao-2/
├── generica/          # Módulo de programação genérica
│   ├── operacoes/     # Implementações de operações genéricas
│   ├── main.go        # Exemplo de uso
│   └── README.md      # Documentação específica
├── poo/               # Módulo de programação orientada a objetos
│   ├── *.go           # Arquivos do sistema de e-commerce
│   └── README.md      # Documentação específica
├── reativa/           # Módulo de programação reativa
│   ├── mensageiro/    # Interface de mensageiro
│   ├── mensagem/      # Interface e tipos de mensagem
│   ├── impl/          # Implementações concretas
│   ├── main.go        # Fluxo reativo principal
│   └── README.md      # Documentação específica
└── README.md          # Este arquivo
```

## Objetivos de Aprendizado

Este projeto visa demonstrar:

1. **Versatilidade do Go**: Como uma linguagem pode suportar múltiplos paradigmas
2. **Generics em Go**: Uso moderno de type parameters para type safety
3. **POO sem classes**: Como Go implementa conceitos OOP através de structs, interfaces e composição
4. **Concorrência e Reatividade**: Uso de goroutines e channels para programação assíncrona e reativa

## Observações

- Cada módulo possui seu próprio README com detalhes específicos
- Os exemplos são didáticos e focados em demonstrar os conceitos dos paradigmas
- O código segue convenções Go e boas práticas da linguagem

## Licença

Este projeto foi desenvolvido para fins educacionais como parte da avaliação da disciplina PLP.

