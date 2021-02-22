package aluno

import (
	"log"

	"github.com/nanduzz/go-clean-architecture/escola/dominio"
)

type LogDeAlunoMatriculado struct {
	dominio.AbstractOuvinte
}

func NewLogDeAlunoMatriculado() *LogDeAlunoMatriculado {
	logMatriculado := LogDeAlunoMatriculado{dominio.AbstractOuvinte{}}
	logMatriculado.AbstractOuvinte.Ouvinte = &logMatriculado
	return &logMatriculado
}

// func reageAo(alunoMatriculado AlunoMatriculado) {
func (l *LogDeAlunoMatriculado) ReageAo(evento dominio.Evento) {
	alunoMatriculado, _ := evento.(*AlunoMatriculado)
	log.Printf("aluno com CPF %s matriculado em %s",
		alunoMatriculado.cpfDoAluno.numero,
		alunoMatriculado.momento)
}

func (l *LogDeAlunoMatriculado) DeveProcessar(event dominio.Evento) bool {
	_, ok := event.(*AlunoMatriculado)
	return ok
}
