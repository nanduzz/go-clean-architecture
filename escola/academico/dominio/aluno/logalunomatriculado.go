package aluno

import (
	"log"

	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/evento"
)

type LogDeAlunoMatriculado struct {
	evento.AbstractOuvinte
}

func NewLogDeAlunoMatriculado() *LogDeAlunoMatriculado {
	logMatriculado := LogDeAlunoMatriculado{evento.AbstractOuvinte{}}
	logMatriculado.AbstractOuvinte.Ouvinte = &logMatriculado
	return &logMatriculado
}

func (l *LogDeAlunoMatriculado) ReageAo(evento evento.Evento) {
	alunoMatriculado, _ := evento.(*AlunoMatriculado)
	log.Printf("aluno com CPF %s matriculado em %s",
		alunoMatriculado.cpfDoAluno.GetNumero(),
		alunoMatriculado.momento)
}

func (l *LogDeAlunoMatriculado) DeveProcessar(event evento.Evento) bool {
	return event.Tipo() == evento.ALUNO_MATRICULADO
}
