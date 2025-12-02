# Programação Genérica em Go

## Propósito

Este programa demonstra o uso de programação genérica em Go através de type parameters (generics), introduzidos na versão 1.18. O código apresenta uma calculadora genérica similar ao exemplo Java, com:

- **Interface genérica**: `Operacao[TipoResultado]` que define operações genéricas
- **Implementações genéricas**: `Multiplicacao` retorna `int` e `Divisao` retorna `float64`
- **Type safety**: Cada operação mantém seu tipo de retorno específico através de generics

O programa demonstra como Go permite escrever código reutilizável e type-safe sem perder a segurança de tipos em tempo de compilação.

## Linguagem

Golang (Go 1.18+)

## Paradigma demonstrado

**Genérico** - Utilizando type parameters para criar código reutilizável que mantém a segurança de tipos.

## Autor(es)

Alisson Rodrigues e Jamille Monteiro

## Como executar

```bash
cd generica
go run main.go
```

## Estrutura do código

- **operacoes/operacao.go**: Interface genérica `Operacao[TipoResultado]`
- **operacoes/multiplicacao.go**: Implementação que retorna `int`
- **operacoes/divisao.go**: Implementação que retorna `float64`
- **main.go**: Demonstração do uso das operações genéricas

