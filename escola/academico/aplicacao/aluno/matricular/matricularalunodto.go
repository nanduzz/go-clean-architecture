package matricular

import (
	"github.com/nanduzz/go-clean-architecture/escola/academico/dominio/aluno"
	"github.com/nanduzz/go-clean-architecture/escola/shared/dominio/sharedaluno"
)

type MatricularAlunoDTO struct {
	nomeAluno  string
	cpfAluno   string
	emailAluno string
}

func NewMatricularAlunoDTO(nome, cpf, email string) *MatricularAlunoDTO {
	return &MatricularAlunoDTO{
		nomeAluno:  nome,
		cpfAluno:   cpf,
		emailAluno: email,
	}
}

func (m *MatricularAlunoDTO) criarAluno() (*aluno.Aluno, error) {
	cpf, err := sharedaluno.NewCPF(m.cpfAluno)
	if err != nil {
		return nil, err
	}

	email, err := aluno.NewEmail(m.emailAluno)
	if err != nil {
		return nil, err
	}

	return aluno.NewAluno(m.nomeAluno, cpf, email), nil
}
