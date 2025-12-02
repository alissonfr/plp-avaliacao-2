package main

import "fmt"

type Eletronico struct {
	id        string
	nome      string
	marca     string
	preco     float64
	estoque   int
	garantia  int
}

func NewEletronico(id, nome, marca string, preco float64, estoque, garantia int) *Eletronico {
	if preco < 0 {
		preco = 0
	}
	if estoque < 0 {
		estoque = 0
	}
	if garantia < 0 {
		garantia = 0
	}
	return &Eletronico{
		id:       id,
		nome:     nome,
		marca:    marca,
		preco:    preco,
		estoque:  estoque,
		garantia: garantia,
	}
}

func (e *Eletronico) GetID() string {
	return e.id
}

func (e *Eletronico) GetNome() string {
	return e.nome
}

func (e *Eletronico) GetMarca() string {
	return e.marca
}

func (e *Eletronico) GetPreco() float64 {
	return e.preco
}

func (e *Eletronico) SetPreco(preco float64) {
	if preco >= 0 {
		e.preco = preco
	}
}

func (e *Eletronico) GetEstoque() int {
	return e.estoque
}

func (e *Eletronico) SetEstoque(quantidade int) {
	if quantidade >= 0 {
		e.estoque = quantidade
	}
}

func (e *Eletronico) GetGarantia() int {
	return e.garantia
}

func (e *Eletronico) CalcularDesconto(percentual float64) float64 {
	if percentual < 0 || percentual > 100 {
		return e.preco
	}
	return e.preco * (1 - percentual/100)
}

func (e *Eletronico) String() string {
	return fmt.Sprintf("Eletrônico: %s %s - R$ %.2f (Estoque: %d, Garantia: %d meses)", e.marca, e.nome, e.preco, e.estoque, e.garantia)
}

