package main

import (
	"fmt"
	"time"
)

type StatusPedido string

const (
	StatusPendente   StatusPedido = "Pendente"
	StatusConfirmado StatusPedido = "Confirmado"
	StatusEnviado    StatusPedido = "Enviado"
	StatusEntregue   StatusPedido = "Entregue"
	StatusCancelado  StatusPedido = "Cancelado"
)

type Pedido struct {
	id        string
	cliente   *Cliente
	itens     []*ItemProduto
	status    StatusPedido
	data      time.Time
	valorTotal float64
}

func NewPedido(id string, cliente *Cliente, itens []*ItemProduto) *Pedido {
	valorTotal := 0.0
	for _, item := range itens {
		valorTotal += item.GetSubtotal()
	}
	
	return &Pedido{
		id:        id,
		cliente:   cliente,
		itens:     itens,
		status:    StatusPendente,
		data:      time.Now(),
		valorTotal: valorTotal,
	}
}

func (p *Pedido) GetID() string {
	return p.id
}

func (p *Pedido) GetCliente() *Cliente {
	return p.cliente
}

func (p *Pedido) GetItens() []*ItemProduto {
	return p.itens
}

func (p *Pedido) GetStatus() StatusPedido {
	return p.status
}

func (p *Pedido) SetStatus(status StatusPedido) {
	p.status = status
}

func (p *Pedido) GetData() time.Time {
	return p.data
}

func (p *Pedido) GetValorTotal() float64 {
	return p.valorTotal
}

func (p *Pedido) Processar() bool {
	if p.status != StatusPendente {
		return false
	}
	
	for _, item := range p.itens {
		produto := item.GetProduto()
		if produto.GetEstoque() < item.GetQuantidade() {
			return false
		}
		produto.SetEstoque(produto.GetEstoque() - item.GetQuantidade())
	}
	
	p.status = StatusConfirmado
	return true
}

func (p *Pedido) String() string {
	return fmt.Sprintf("Pedido #%s - %s - Status: %s - Total: R$ %.2f", p.id, p.cliente.GetNome(), p.status, p.valorTotal)
}

