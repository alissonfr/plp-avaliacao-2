package impl

type TipoMensagem int

const (
	TEXTO TipoMensagem = iota
	IMAGEM
	VIDEO
	AUDIO
	ERRO
)

func (t TipoMensagem) String() string {
	switch t {
	case TEXTO:
		return "TEXTO"
	case IMAGEM:
		return "IMAGEM"
	case VIDEO:
		return "VIDEO"
	case AUDIO:
		return "AUDIO"
	case ERRO:
		return "ERRO"
	default:
		return "UNKNOWN"
	}
}

