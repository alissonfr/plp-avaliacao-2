package main

import "fmt"

type Livro struct {
	id      string
	nome    string
	autor   string
	preco   float64
	estoque int
}

func NewLivro(id, nome, autor string, preco float64, estoque int) *Livro {
	if preco < 0 {
		preco = 0
	}
	if estoque < 0 {
		estoque = 0
	}
	return &Livro{
		id:      id,
		nome:    nome,
		autor:   autor,
		preco:   preco,
		estoque: estoque,
	}
}

func (l *Livro) GetID() string {
	return l.id
}

func (l *Livro) GetNome() string {
	return l.nome
}

func (l *Livro) GetAutor() string {
	return l.autor
}

func (l *Livro) GetPreco() float64 {
	return l.preco
}

func (l *Livro) SetPreco(preco float64) {
	if preco >= 0 {
		l.preco = preco
	}
}

func (l *Livro) GetEstoque() int {
	return l.estoque
}

func (l *Livro) SetEstoque(quantidade int) {
	if quantidade >= 0 {
		l.estoque = quantidade
	}
}

func (l *Livro) CalcularDesconto(percentual float64) float64 {
	if percentual < 0 || percentual > 100 {
		return l.preco
	}
	return l.preco * (1 - percentual/100)
}

func (l *Livro) String() string {
	return fmt.Sprintf("Livro: %s por %s - R$ %.2f (Estoque: %d)", l.nome, l.autor, l.preco, l.estoque)
}

