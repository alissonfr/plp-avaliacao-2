package main

import "fmt"

type Carrinho struct {
	cliente *Cliente
	itens   []*ItemProduto
}

func NewCarrinho(cliente *Cliente) *Carrinho {
	return &Carrinho{
		cliente: cliente,
		itens:   make([]*ItemProduto, 0),
	}
}

func (c *Carrinho) GetCliente() *Cliente {
	return c.cliente
}

func (c *Carrinho) AdicionarItem(produto Produto, quantidade int) bool {
	if quantidade <= 0 || produto.GetEstoque() < quantidade {
		return false
	}
	
	for _, item := range c.itens {
		if item.GetProduto().GetID() == produto.GetID() {
			novaQuantidade := item.GetQuantidade() + quantidade
			if produto.GetEstoque() >= novaQuantidade {
				item.SetQuantidade(novaQuantidade)
				return true
			}
			return false
		}
	}
	
	c.itens = append(c.itens, NewItemProduto(produto, quantidade))
	return true
}

func (c *Carrinho) RemoverItem(produtoID string) bool {
	for i, item := range c.itens {
		if item.GetProduto().GetID() == produtoID {
			c.itens = append(c.itens[:i], c.itens[i+1:]...)
			return true
		}
	}
	return false
}

func (c *Carrinho) GetItens() []*ItemProduto {
	return c.itens
}

func (c *Carrinho) GetTotal() float64 {
	total := 0.0
	for _, item := range c.itens {
		total += item.GetSubtotal()
	}
	return total
}

func (c *Carrinho) Limpar() {
	c.itens = make([]*ItemProduto, 0)
}

func (c *Carrinho) String() string {
	return fmt.Sprintf("Carrinho de %s - %d item(ns) - Total: R$ %.2f", c.cliente.GetNome(), len(c.itens), c.GetTotal())
}

