# Programação Orientada a Objetos em Go

## Propósito

Este programa demonstra os conceitos de Programação Orientada a Objetos (POO) em Go através de um sistema de e-commerce, incluindo:

- **Encapsulamento**: Uso de campos privados (iniciando com letra minúscula) e métodos públicos (getters/setters)
- **Interfaces**: Definição de contratos através da interface `Produto`
- **Polimorfismo**: Diferentes implementações da interface `Produto` (Livro, Eletronico) sendo tratadas uniformemente
- **Composição**: Go não possui herança clássica, mas permite composição (ex: Carrinho compõe Cliente e ItemProduto, Pedido compõe Cliente e ItemProduto)
- **Factory Pattern**: Demonstração de padrão de projeto usando métodos de fábrica para criação de produtos

O programa cria um sistema de e-commerce onde diferentes tipos de produtos (livros, eletrônicos) implementam a mesma interface, permitindo tratamento polimórfico em carrinhos de compras e pedidos.

## Linguagem

Golang (Go)

## Paradigma demonstrado

**Orientado a Objetos** - Organiza o software em objetos com estado e comportamento, usando encapsulamento, herança, polimorfismo e abstração para criar sistemas modulares e extensíveis.

## Autor(es)

Alisson Rodrigues e Jamille Monteiro

## Como executar

```bash
cd poo
go run .
```

## Estrutura do código

- **Produto**: Interface que define o contrato para produtos do e-commerce
- **Livro**: Struct que implementa Produto, representando um livro
- **Eletronico**: Struct que implementa Produto, representando um eletrônico
- **Cliente**: Struct que representa um cliente do e-commerce
- **ItemProduto**: Struct que representa um item no carrinho (composição de Produto)
- **Carrinho**: Struct que compõe Cliente e ItemProduto, demonstrando composição
- **Pedido**: Struct que representa um pedido processado, compondo Cliente e ItemProduto
- **PedidoService**: Serviço que gerencia múltiplos pedidos, demonstrando composição e polimorfismo
- **ProdutoFactory**: Factory para criação de produtos, demonstrando padrão de projeto

## Conceitos demonstrados

1. **Encapsulamento**: Campos privados (`id`, `nome`, `preco`, `estoque`) acessados via métodos públicos (getters/setters)
2. **Interfaces**: Contrato comum (`Produto`) implementado por múltiplos tipos (Livro, Eletronico)
3. **Polimorfismo**: Métodos `CalcularDesconto()` e `String()` chamados polimorficamente para diferentes tipos de produtos
4. **Composição**: `Carrinho` e `Pedido` reutilizam funcionalidade através de composição com `Cliente` e `ItemProduto`
5. **Factory Pattern**: Criação centralizada de objetos através de uma fábrica de produtos
