package main

import "fmt"

type Cliente struct {
	id      string
	nome    string
	email   string
	endereco string
}

func NewCliente(id, nome, email, endereco string) *Cliente {
	return &Cliente{
		id:       id,
		nome:     nome,
		email:    email,
		endereco: endereco,
	}
}

func (c *Cliente) GetID() string {
	return c.id
}

func (c *Cliente) GetNome() string {
	return c.nome
}

func (c *Cliente) SetNome(nome string) {
	c.nome = nome
}

func (c *Cliente) GetEmail() string {
	return c.email
}

func (c *Cliente) SetEmail(email string) {
	c.email = email
}

func (c *Cliente) GetEndereco() string {
	return c.endereco
}

func (c *Cliente) SetEndereco(endereco string) {
	c.endereco = endereco
}

func (c *Cliente) String() string {
	return fmt.Sprintf("Cliente: %s (%s) - %s", c.nome, c.email, c.endereco)
}

