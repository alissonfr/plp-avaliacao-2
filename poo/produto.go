package main

type Produto interface {
	GetID() string
	GetNome() string
	GetPreco() float64
	GetEstoque() int
	SetEstoque(quantidade int)
	CalcularDesconto(percentual float64) float64
	String() string
}

type ItemProduto struct {
	produto   Produto
	quantidade int
}

func NewItemProduto(produto Produto, quantidade int) *ItemProduto {
	if quantidade < 0 {
		quantidade = 0
	}
	return &ItemProduto{
		produto:   produto,
		quantidade: quantidade,
	}
}

func (i *ItemProduto) GetProduto() Produto {
	return i.produto
}

func (i *ItemProduto) GetQuantidade() int {
	return i.quantidade
}

func (i *ItemProduto) SetQuantidade(quantidade int) {
	if quantidade >= 0 {
		i.quantidade = quantidade
	}
}

func (i *ItemProduto) GetSubtotal() float64 {
	return i.produto.GetPreco() * float64(i.quantidade)
}

