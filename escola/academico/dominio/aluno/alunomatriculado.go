package aluno

import (
	"time"

	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/evento"
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"
)

type AlunoMatriculado struct {
	cpfDoAluno *sharedaluno.CPF
	momento    time.Time
}

func NewAlunoMatriculado(cpfDoAluno *sharedaluno.CPF) *AlunoMatriculado {
	return &AlunoMatriculado{
		cpfDoAluno: cpfDoAluno,
		momento:    time.Now(),
	}
}

func (am *AlunoMatriculado) Momento() time.Time {
	return am.momento
}

func (am *AlunoMatriculado) GetCpf() *sharedaluno.CPF {
	return am.cpfDoAluno
}

func (am *AlunoMatriculado) Tipo() evento.TipoDeEvento {
	return evento.ALUNO_MATRICULADO
}

func (am *AlunoMatriculado) Informacoes() map[string]interface{} {
	info := make(map[string]interface{})
	info["cpf"] = am.cpfDoAluno

	return info
}
