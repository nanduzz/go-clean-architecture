package matricular

import (
	"github.com/nanduzz/go-clean-architecture/escola/academico/dominio/aluno"
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/evento"
)

type MatricularAluno struct {
	repositorio aluno.RepositorioDeAlunos
	publicador  *evento.PublicadorEventos
}

func NewMatricularAluno(repositorio aluno.RepositorioDeAlunos, publicador *evento.PublicadorEventos) *MatricularAluno {
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
