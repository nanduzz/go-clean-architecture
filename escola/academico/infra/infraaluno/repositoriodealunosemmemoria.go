package infraaluno

import (
	"fmt"

	"github.com/nanduzz/go-clean-architecture/escola/academico/dominio/aluno"
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"
)

type RepositorioDeAlunosEmMemoria struct {
	store map[sharedaluno.CPF]*aluno.Aluno
}

func NewRepositorioDeAlunosEmMemoria() *RepositorioDeAlunosEmMemoria {
	return &RepositorioDeAlunosEmMemoria{
		store: make(map[sharedaluno.CPF]*aluno.Aluno),
	}
}

func (r *RepositorioDeAlunosEmMemoria) Matricular(aluno *aluno.Aluno) error {
	r.store[*aluno.GetCpf()] = aluno
	return nil
}

func (r *RepositorioDeAlunosEmMemoria) BuscaPorCPF(cpf sharedaluno.CPF) (*aluno.Aluno, error) {
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
