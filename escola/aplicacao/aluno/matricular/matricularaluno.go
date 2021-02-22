package matricular

import (
	"github.com/nanduzz/go-clean-architecture/escola/dominio"
	"github.com/nanduzz/go-clean-architecture/escola/dominio/aluno"
)

type MatricularAluno struct {
	repositorio aluno.RepositorioDeAlunos
	publicador  *dominio.PublicadorEventos
}

func NewMatricularAluno(repositorio aluno.RepositorioDeAlunos, publicador *dominio.PublicadorEventos) *MatricularAluno {
	return &MatricularAluno{
		repositorio: repositorio,
		publicador:  publicador,
	}
}
func (m *MatricularAluno) Executa(dados *MatricularAlunoDTO) error {
	alunoE, err := dados.criarAluno()
	if err != nil {
		return err
	}
	m.repositorio.Matricular(alunoE)

	alunoMatriculadoEvento := aluno.NewAlunoMatriculado(alunoE.GetCpf())
	m.publicador.Publicar(alunoMatriculadoEvento)
	return nil
}
