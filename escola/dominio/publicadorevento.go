package dominio

type PublicadorEventos struct {
	ouvintes []AbstractOuvinte
}

func NewPublicadorEventos() *PublicadorEventos {
	return &PublicadorEventos{}
}

func (p *PublicadorEventos) Adicionar(ouvinte AbstractOuvinte) {
	p.ouvintes = append(p.ouvintes, ouvinte)
}

func (p *PublicadorEventos) Publicar(evento Evento) {
	for _, ouvinte := range p.ouvintes {
		ouvinte.Processa(evento)
	}
}
