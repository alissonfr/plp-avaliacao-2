package main

import "fmt"

type PedidoService struct {
	pedidos []*Pedido
}

func NewPedidoService() *PedidoService {
	return &PedidoService{
		pedidos: make([]*Pedido, 0),
	}
}

func (ps *PedidoService) CriarPedido(carrinho *Carrinho) *Pedido {
	if len(carrinho.GetItens()) == 0 {
		return nil
	}
	
	itens := make([]*ItemProduto, len(carrinho.GetItens()))
	for i, item := range carrinho.GetItens() {
		itens[i] = NewItemProduto(item.GetProduto(), item.GetQuantidade())
	}
	
	id := fmt.Sprintf("PED-%d", len(ps.pedidos)+1)
	pedido := NewPedido(id, carrinho.GetCliente(), itens)
	
	if pedido.Processar() {
		ps.pedidos = append(ps.pedidos, pedido)
		carrinho.Limpar()
		return pedido
	}
	
	return nil
}

func (ps *PedidoService) GetPedido(id string) *Pedido {
	for _, pedido := range ps.pedidos {
		if pedido.GetID() == id {
			return pedido
		}
	}
	return nil
}

func (ps *PedidoService) GetPedidos() []*Pedido {
	return ps.pedidos
}

func (ps *PedidoService) GetPedidosPorCliente(clienteID string) []*Pedido {
	var pedidosCliente []*Pedido
	for _, pedido := range ps.pedidos {
		if pedido.GetCliente().GetID() == clienteID {
			pedidosCliente = append(pedidosCliente, pedido)
		}
	}
	return pedidosCliente
}

func (ps *PedidoService) AtualizarStatus(id string, status StatusPedido) bool {
	pedido := ps.GetPedido(id)
	if pedido == nil {
		return false
	}
	pedido.SetStatus(status)
	return true
}

func (ps *PedidoService) GetTotalVendas() float64 {
	total := 0.0
	for _, pedido := range ps.pedidos {
		if pedido.GetStatus() != StatusCancelado {
			total += pedido.GetValorTotal()
		}
	}
	return total
}

func (ps *PedidoService) ListarPedidos() {
	fmt.Println("Pedidos registrados:")
	for i, pedido := range ps.pedidos {
		fmt.Printf("  %d. %s\n", i+1, pedido)
	}
}

