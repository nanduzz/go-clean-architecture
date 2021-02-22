package aluno

import "time"

type AlunoMatriculado struct {
	cpfDoAluno *CPF
	momento    time.Time
}

func NewAlunoMatriculado(cpfDoAluno *CPF) *AlunoMatriculado {
	return &AlunoMatriculado{
		cpfDoAluno: cpfDoAluno,
		momento:    time.Now(),
	}
}

func (am *AlunoMatriculado) Momento() time.Time {
	return am.momento
}

func (am *AlunoMatriculado) GetCpf() *CPF {
	return am.cpfDoAluno
}
