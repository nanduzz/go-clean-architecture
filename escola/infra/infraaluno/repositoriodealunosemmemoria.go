package infraaluno

import (
	"fmt"

	"github.com/nanduzz/go-clean-architecture/escola/dominio/aluno"
)

type RepositorioDeAlunosEmMemoria struct {
	store map[aluno.CPF]*aluno.Aluno
}

func NewRepositorioDeAlunosEmMemoria() *RepositorioDeAlunosEmMemoria {
	return &RepositorioDeAlunosEmMemoria{
		store: make(map[aluno.CPF]*aluno.Aluno),
	}
}

func (r *RepositorioDeAlunosEmMemoria) Matricular(aluno *aluno.Aluno) error {
	r.store[*aluno.GetCpf()] = aluno
	return nil
}

func (r *RepositorioDeAlunosEmMemoria) BuscaPorCPF(cpf aluno.CPF) (*aluno.Aluno, error) {
	aluno, ok := r.store[cpf]
	if !ok {
		return nil, fmt.Errorf("Aluno não encontrado! cpf :", cpf.GetNumero())
	}

	return aluno, nil
}

func (r *RepositorioDeAlunosEmMemoria) ListarTodosAlunosMatriculados() ([]*aluno.Aluno, error) {
	alunos := make([]*aluno.Aluno, len(r.store))
	for _, aluno := range r.store {
		alunos = append(alunos, aluno)
	}
	return alunos, nil
}
