package geraselo

import (
	"github.com/nanduzz/go-clean-architecture/escola/gamificacao/dominio/selo"
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/evento"
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"
)

type GeraSeloAlunoNovato struct {
	evento.AbstractOuvinte
	repositorio selo.RepositorioDeSelos
}

func NewGeraSeloAlunoNovato(repositorio selo.RepositorioDeSelos) *GeraSeloAlunoNovato {
	this := GeraSeloAlunoNovato{
		evento.AbstractOuvinte{},
		repositorio,
	}
	this.AbstractOuvinte.Ouvinte = &this
	return &this
}

func (l *GeraSeloAlunoNovato) ReageAo(evento evento.Evento) {
	info, ok := evento.Informacoes()["cpf"]
	if !ok {
		panic("Not ok")
	}
	if cpf, ok := info.(*sharedaluno.CPF); ok {
		selo := selo.NewSelo(cpf, "Novato")
		l.repositorio.Adicionar(selo)
	}
	panic("Evento errado")
}

func (l *GeraSeloAlunoNovato) DeveProcessar(event evento.Evento) bool {
	return event.Tipo() == evento.ALUNO_MATRICULADO
}
