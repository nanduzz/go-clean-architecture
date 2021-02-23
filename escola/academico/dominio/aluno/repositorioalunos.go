package aluno

import "github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"

type RepositorioDeAlunos interface {
	Matricular(aluno *Aluno) error

	BuscaPorCPF(cpf sharedaluno.CPF) (*Aluno, error)

	ListarTodosAlunosMatriculados() ([]*Aluno, error)
}
