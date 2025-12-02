package main

type ProdutoFactory struct{}

func NewProdutoFactory() *ProdutoFactory {
	return &ProdutoFactory{}
}

func (pf *ProdutoFactory) CriarProduto(tipo, id, nome string, preco float64, estoque int, extras ...interface{}) Produto {
	switch tipo {
	case "livro":
		autor := ""
		if len(extras) > 0 {
			if a, ok := extras[0].(string); ok {
				autor = a
			}
		}
		return NewLivro(id, nome, autor, preco, estoque)
	case "eletronico":
		marca := ""
		garantia := 0
		if len(extras) > 0 {
			if m, ok := extras[0].(string); ok {
				marca = m
			}
		}
		if len(extras) > 1 {
			if g, ok := extras[1].(int); ok {
				garantia = g
			}
		}
		return NewEletronico(id, nome, marca, preco, estoque, garantia)
	default:
		return nil
	}
}

