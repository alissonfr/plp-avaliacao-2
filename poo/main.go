package main

import "fmt"

func main() {
	fmt.Println("=== Sistema de E-commerce ===\n")

	cliente1 := NewCliente("C001", "João Silva", "joao@email.com", "Rua A, 123")
	cliente2 := NewCliente("C002", "Maria Santos", "maria@email.com", "Rua B, 456")

	fmt.Println("1. Criação de Produtos:")
	livro1 := NewLivro("L001", "Aprendendo Go", "Autor X", 49.90, 10)
	livro2 := NewLivro("L002", "Programação Avançada", "Autor Y", 79.90, 5)
	eletronico1 := NewEletronico("E001", "Smartphone", "TechBrand", 1299.90, 8, 12)
	eletronico2 := NewEletronico("E002", "Notebook", "TechBrand", 3499.90, 3, 24)

	fmt.Printf("   %s\n", livro1)
	fmt.Printf("   %s\n", livro2)
	fmt.Printf("   %s\n", eletronico1)
	fmt.Printf("   %s\n", eletronico2)
	fmt.Println()

	fmt.Println("2. Encapsulamento (Getters e Setters):")
	fmt.Printf("   Preço original do livro: R$ %.2f\n", livro1.GetPreco())
	livro1.SetPreco(44.90)
	fmt.Printf("   Preço com desconto: R$ %.2f\n", livro1.GetPreco())
	fmt.Printf("   Estoque atual: %d\n", livro1.GetEstoque())
	fmt.Println()

	fmt.Println("3. Interfaces e Polimorfismo:")
	var produtos []Produto
	produtos = append(produtos, livro1, livro2, eletronico1, eletronico2)

	for i, produto := range produtos {
		fmt.Printf("   Produto %d: %s\n", i+1, produto)
		precoComDesconto := produto.CalcularDesconto(10.0)
		fmt.Printf("      Preço com 10%% de desconto: R$ %.2f\n", precoComDesconto)
	}
	fmt.Println()

	fmt.Println("4. Composição (Carrinho de Compras):")
	carrinho1 := NewCarrinho(cliente1)
	carrinho1.AdicionarItem(livro1, 2)
	carrinho1.AdicionarItem(eletronico1, 1)
	fmt.Printf("   %s\n", carrinho1)
	fmt.Printf("   Itens no carrinho:\n")
	for i, item := range carrinho1.GetItens() {
		fmt.Printf("      %d. %s - Qtd: %d - Subtotal: R$ %.2f\n",
			i+1, item.GetProduto().GetNome(), item.GetQuantidade(), item.GetSubtotal())
	}
	fmt.Println()

	carrinho2 := NewCarrinho(cliente2)
	carrinho2.AdicionarItem(livro2, 1)
	carrinho2.AdicionarItem(eletronico2, 1)
	fmt.Printf("   %s\n", carrinho2)
	fmt.Println()

	fmt.Println("5. Serviço de Pedidos:")
	pedidoService := NewPedidoService()
	
	pedido1 := pedidoService.CriarPedido(carrinho1)
	if pedido1 != nil {
		fmt.Printf("   %s\n", pedido1)
		fmt.Printf("   Itens do pedido:\n")
		for i, item := range pedido1.GetItens() {
			fmt.Printf("      %d. %s - Qtd: %d\n", i+1, item.GetProduto().GetNome(), item.GetQuantidade())
		}
	}
	fmt.Println()

	pedido2 := pedidoService.CriarPedido(carrinho2)
	if pedido2 != nil {
		fmt.Printf("   %s\n", pedido2)
	}
	fmt.Println()

	fmt.Println("6. Atualização de Status:")
	pedidoService.AtualizarStatus(pedido1.GetID(), StatusEnviado)
	pedidoService.AtualizarStatus(pedido2.GetID(), StatusConfirmado)
	pedidoService.ListarPedidos()
	fmt.Printf("   Total de vendas: R$ %.2f\n\n", pedidoService.GetTotalVendas())

	fmt.Println("7. Factory Pattern:")
	factory := NewProdutoFactory()
	novoLivro := factory.CriarProduto("livro", "L003", "Novo Livro", 59.90, 15, "Autor Z")
	novoEletronico := factory.CriarProduto("eletronico", "E003", "Tablet", 899.90, 6, "TechBrand", 12)
	
	fmt.Printf("   Produto criado pela factory: %s\n", novoLivro)
	fmt.Printf("   Produto criado pela factory: %s\n", novoEletronico)
	fmt.Println()

	fmt.Println("8. Consulta de Pedidos por Cliente:")
	pedidosCliente1 := pedidoService.GetPedidosPorCliente(cliente1.GetID())
	fmt.Printf("   Pedidos de %s:\n", cliente1.GetNome())
	for _, pedido := range pedidosCliente1 {
		fmt.Printf("      %s\n", pedido)
	}
}
