package mensagem

import "time"

type Mensagem interface {
	GetTipo() interface{}
	GetConteudo() interface{}
	GetDataHora() time.Time
	SetDataHora(dataHora time.Time)
	Enviar() Resultado
}

type BaseMensagem struct {
	tipo     interface{}
	conteudo interface{}
	dataHora time.Time
}

func NewBaseMensagem(tipo interface{}, conteudo interface{}) *BaseMensagem {
	return &BaseMensagem{
		tipo:     tipo,
		conteudo: conteudo,
	}
}

func (m *BaseMensagem) GetTipo() interface{} {
	return m.tipo
}

func (m *BaseMensagem) GetConteudo() interface{} {
	return m.conteudo
}

func (m *BaseMensagem) GetDataHora() time.Time {
	return m.dataHora
}

func (m *BaseMensagem) SetDataHora(dataHora time.Time) {
	m.dataHora = dataHora
}

