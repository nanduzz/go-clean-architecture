package matricular

import "github.com/nanduzz/go-clean-architecture/escola/dominio/aluno"

type MatricularAluno struct {
	repositorio aluno.RepositorioDeAlunos
}

func NewMatricularAluno(repositorio aluno.RepositorioDeAlunos) *MatricularAluno {
	return &MatricularAluno{
		repositorio: repositorio,
	}
}
func (m *MatricularAluno) Executa(dados *MatricularAlunoDTO) error {
	aluno, err := dados.criarAluno()
	if err != nil {
		return err
	}
	m.repositorio.Matricular(aluno)
	return nil
}
